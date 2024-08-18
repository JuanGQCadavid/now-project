package main

import (
	"context"
	"encoding/json"
	"log"
	"runtime"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/cmd/ssm"
	credFinderCore "github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/core/core/domain"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/cmd/lambda/utils"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"

	// "github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/confirmation/queue"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/confirmation/localconfirmation"
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
	DetectPendingDatesToClose        Operations = "detectPendingDatesToClose"
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

		// SPECIALIZED ROUTINES
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
		} else if operation == DetectPendingDatesToClose {
			// TODO: Here, what to do when pending dates to close is invoked
			continue
		}

		var body utils.Body
		err := json.Unmarshal([]byte(record.Body), &body)

		if err != nil {
			logs.Error.Println("Oh shit, aborting this record due to:", err.Error())
			continue
		}

		// SCHEDULE PATTERNS SNS

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
	// TODO: We should simplify this switch stuff with a map maybe
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
	case string(DetectPendingDatesToClose):
		return DetectPendingDatesToClose
	default:
		logs.Warning.Printf("Operation %s is not recognized \n", value)
		return Other
	}
}

//////////////////
//
// New proposal
//
//////////////////

type Record struct {
	body      string
	operation Operations
}

type Result struct {
	err error
}

// type PipelineEventConfig struct {
// 	operations []Operations
// 	pipeline   func(<-chan interface{}, <-chan Record) chan Result
// }

func emitRecords(done <-chan interface{}, source *events.SQSEvent) <-chan Record {
	valueStream := make(chan Record)

	go func() {
		defer close(valueStream)

		for _, record := range source.Records {
			var (
				operation Operations
			)

			// logs.Info.Printf("Sending record %d with data %s \n", index, record.Body)
			if operation = GetOperationNameFromAttributes(record); operation == Other {
				logs.Warning.Println("Operation not founded, looking on the body")
				operation = GetOperationNameFromBody(record)
			}

			// logs.Info.Printf("Operation: %s\n", string(operation))

			select {
			case <-done:
				return
			case valueStream <- Record{
				body:      record.Body,
				operation: operation,
			}:
			}

		}
	}()

	return valueStream
}

func creationOfDates(done <-chan interface{}, events <-chan domain.Spot) <-chan Result {
	out := make(chan Result)
	go func() {

		// toCreate = append(toCreate, utils.FromSpotRequestToSpot(body.SpotRequest))

		toCreate := make([]domain.Spot, 0, len(events))
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
	}()

	return out
}

func deletionOfDates(done <-chan interface{}, events <-chan string) <-chan Result {
	out := make(chan Result)

	go func() {
		defer close(out)

		toDelete := make([]string, 0, len(events))

		for event := range events {
			toDelete = append(toDelete, event)
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

			select {
			case <-done:
				return
			case out <- Result{
				err: err,
			}:

			}
		}
	}()
	return out
}

func scheduleEventsIdentifer(done <-chan interface{}, events <-chan Record) (<-chan domain.Spot, <-chan string) {
	createStream, deleteStream := make(chan domain.Spot), make(chan string)

	go func() {
		defer close(createStream)
		defer close(deleteStream)

		for record := range events {

			var body utils.Body
			err := json.Unmarshal([]byte(record.body), &body)

			if err != nil {
				logs.Error.Println("Oh shit, aborting this record due to:", err.Error())
				continue
			}

			if record.operation == SchedulePatternAppended || record.operation == SchedulePatternResumed {
				if (len(body.SpotRequest.SpotInfo.SpotId) + len(body.SpotRequest.SpotPatterns)) == 0 {
					logs.Error.Println("SpotRequest is empty")
					continue
				}

				select {
				case <-done:
					return
				case createStream <- utils.FromSpotRequestToSpot(body.SpotRequest):
				}

			} else if record.operation == SchedulePatternConcluded || record.operation == SchedulePatternFreezed {
				if len(body.ScheduleId) == 0 {
					logs.Error.Println("ScheduleId is empty")
					continue
				}

				select {
				case <-done:
					return
				case deleteStream <- body.ScheduleId:
				}

			} else {
				logs.Warning.Println("Operation not recognized, aborting message")
			}

		}

	}()

	return createStream, deleteStream
}

// func filterRecords(done <-chan interface{}, events <-chan Record, configMap map[string]PipelineEventConfig) map[string]chan Result {

// 	return nil
// }

func HandlerV2(ctx context.Context, body *events.SQSEvent) (string, error) {

	// New code
	var (
		done chan interface{} = make(chan interface{})
	)

	emitRecords(done, body)

	// filter(done, chan of results, configmap) - a map of chans, one per config in the configmap
	for _, record := range body.Records {

		operation := GetOperationNameFromAttributes(record)

		if operation == Other {
			logs.Warning.Println("Operation not founded, looking on the body")
			operation = GetOperationNameFromBody(record)
		}

		log.Println(operation)

		// SPECIALIZED ROUTINES
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
		} else if operation == DetectPendingDatesToClose {
			// TODO: Here, what to do when pending dates to close is invoked
			continue
		}
	}

	return AcknowledgeKeyworkd, nil
}

//////////////////
//
// end New propolsal
//
//////////////////

func initWithDefaults() {
	var (
		credsFinder credFinderCore.CredentialsFinder
		neo4jDriver neo4j.Driver
		err         error
	)

	logs.Info.Println("On init")

	credsFinder = ssm.NewSSMCredentialsFinder()
	neo4jDriver, err = credsFinder.FindNeo4jCredentialsFromDefaultEnv()

	if err != nil {
		logs.Error.Println("There were an error while attempting to create drivers")
		logs.Error.Fatalln(err.Error())
	}
	logs.Info.Println("Well we are now here")

	repo := neo4jrepo.NewNeo4jRepoWithDriver(neo4jDriver)
	// queueConfirmation, err := queue.NewSQSConfirmationFromEnv(SqsConfirmationArn)
	queueConfirmation := localconfirmation.NewLocalConfirmation()

	if err != nil {
		logs.Error.Fatalln("error while creatin repo", err.Error())
	}

	notifier, err := topics.NewNotifierFromEnv(TopicArnEnvName)

	if err != nil {
		logs.Error.Fatalln("error while creatin notifer", err.Error())
	}

	log.Println("Number of CPU", runtime.NumCPU())

	srv = service.NewCheckerService(repo, queueConfirmation, notifier, runtime.NumCPU())
	logs.Info.Println("Service created")
}

func main() {
	initWithDefaults()
	lambda.Start(Handler)
}
