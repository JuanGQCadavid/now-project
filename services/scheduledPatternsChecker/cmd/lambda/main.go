package main

import (
	"context"
	"encoding/json"
	"log"
	"runtime"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/cmd/ssm"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/cmd/lambda/utils"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/confirmation/localconfirmation"

	// "github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/confirmation/queue"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/service"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/notifiers/topics"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/repository/neo4jrepo"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Operations string

const (
	Operation                        string     = "Operation"
	SchedulePatternAppended          Operations = "schedulePatternAppended"
	SchedulePatternConcluded         Operations = "schedulePatternConcluded"
	SchedulePatternResumed           Operations = "schedulePatternResumed"
	SchedulePatternFreezed           Operations = "schedulePatternFreezed"
	GenrateDatesFromSchedulePatterns Operations = "generateDatesFromSchedulePatterns"
	Other                            Operations = "other"
	DefaultTimeWindow                int64      = 604800
	TopicArnEnvName                  string     = "snsArn"
	SqsConfirmationArn               string     = "sqsConfirmationArn"
	MaxBatchSize                     int        = 10
	AcknowledgeKeyworkd              string     = "Done"
)

var (
	srv ports.Service
)

func Handler(ctx context.Context, body *events.SQSEvent) (string, error) {
	toDelete := make([]string, 0, MaxBatchSize)
	toCreate := make([]domain.Spot, 0, MaxBatchSize)

	for _, record := range body.Records {

		log.Printf("%+v \n", record)
		log.Println("------------")
		log.Printf("%+v \n", record.Body)

		log.Println(" ********** ")

		operation := GetOperationNameFromAttributes(record)

		if operation == Other {
			logs.Warning.Println("Operation not founded, looking on the body")
			operation = GetOperationNameFromBody(record)
		}

		log.Println(operation)

		if operation == GenrateDatesFromSchedulePatterns {
			logs.Info.Println("Operation: ", GenrateDatesFromSchedulePatterns)

			var body utils.BatchRequest
			err := json.Unmarshal([]byte(record.Body), &body)

			if err != nil {
				logs.Error.Println("Oh shit, aborting this record due to:", err.Error())
				continue
			}

			timeWindow := body.TimeWindow

			if timeWindow == 0 {
				logs.Warning.Println("Time window is 0, using defautl time window of ", DefaultTimeWindow)
				timeWindow = DefaultTimeWindow
			}

			result, err := srv.GenerateDatesFromRepository(timeWindow)

			if err != nil {
				logs.Error.Println("Service faile on error :", err.Error())
			}

			if result != nil {
				logs.Info.Println("Result total length:", len(result))
			} else {
				logs.Info.Println("Empty result")
			}

			continue
		}

		var body utils.Body
		err := json.Unmarshal([]byte(record.Body), &body)

		if err != nil {
			logs.Error.Println("Oh shit, aborting this record due to:", err.Error())
			continue
		}

		if operation == SchedulePatternAppended || operation == SchedulePatternResumed {

			if (len(body.SpotRequest.SpotInfo.SpotId) + len(body.SpotRequest.SpotPatterns)) == 0 {
				logs.Error.Println("SpotRequest is empty")
				continue
			}

			toCreate = append(toCreate, utils.FromSpotRequestToSpot(body.SpotRequest))

		} else if operation == SchedulePatternConcluded || operation == SchedulePatternFreezed {

			if len(body.ScheduleId) == 0 {
				logs.Error.Println("ScheduleId is empty")
				continue
			}

			toDelete = append(toDelete, body.ScheduleId)
		} else {
			logs.Warning.Println("Operation not recognized, aborting message")

		}
	}

	if len(toDelete) > 0 {
		logs.Info.Println("Dates to delete from schedules Id:")
		for _, spot := range toDelete {
			logs.Info.Printf("Scheudle Id %s \n", spot)
		}

		err := srv.DeleteScheduleDatesFromSchedulePattern(toDelete)

		if err != nil {
			logs.Error.Println("We found the next error ", err.Error())
		}

	}

	if len(toCreate) > 0 {
		logs.Info.Println("Dates to create from schedules Id:")
		for _, spot := range toCreate {
			logs.Info.Printf("Scheudle Id %+v \n", spot)
		}

		if _, errs := srv.CreateScheduledDatesFromSchedulePattern(toCreate, DefaultTimeWindow); errs != nil {
			for err, errSpot := range errs {
				logs.Error.Println("To create Error ", err.Error(), " on ", errSpot)
			}
		}

	}

	return AcknowledgeKeyworkd, nil
}
func GetOperationNameFromBody(record events.SQSMessage) Operations {
	var body utils.BatchRequest
	err := json.Unmarshal([]byte(record.Body), &body)

	if err != nil {
		logs.Error.Println("There where an error while unmarshaling the body ", err.Error())
		return Other
	}

	return stringToOperation(body.Operation)
}

func GetOperationNameFromAttributes(record events.SQSMessage) Operations {

	operation := record.MessageAttributes[Operation]

	if operation.StringValue != nil {
		return stringToOperation(*operation.StringValue)
	}

	logs.Warning.Println("Operation parameter is missing")
	return Other
}

func stringToOperation(value string) Operations {
	switch value {
	case string(SchedulePatternAppended):
		return SchedulePatternAppended
	case string(SchedulePatternConcluded):
		return SchedulePatternConcluded
	case string(SchedulePatternResumed):
		return SchedulePatternResumed
	case string(SchedulePatternFreezed):
		return SchedulePatternFreezed
	case string(GenrateDatesFromSchedulePatterns):
		return GenrateDatesFromSchedulePatterns
	default:
		logs.Warning.Printf("Operation %s is not recognized \n", value)
		return Other
	}
}

func init() {
	logs.Info.Println("On init")

	credsFinder := ssm.NewSSMCredentialsFinder()

	neo4jDriver, err := credsFinder.FindNeo4jCredentialsFromDefaultEnv()

	if err != nil {
		logs.Error.Println("There were an error while attempting to create drivers")
		logs.Error.Fatalln(err.Error())
	}
	repo := neo4jrepo.NewNeo4jRepoWithDriver(neo4jDriver)
	// queueConfirmation, err := queue.NewSQSConfirmationFromEnv(SqsConfirmationArn)
	queueConfirmation := localconfirmation.NewLocalConfirmation()

	if err != nil {
		logs.Error.Fatalln("error while creatin repo", err.Error())
	}

	notifier, err := topics.NewNotifierFromEnv(TopicArnEnvName)

	log.Println("Number of CPU", runtime.NumCPU())

	srv = service.NewCheckerService(repo, queueConfirmation, notifier, runtime.NumCPU())
	logs.Info.Println("Service created")
}

func main() {
	lambda.Start(Handler)
}
