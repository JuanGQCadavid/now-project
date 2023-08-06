package main

import (
	"context"
	"encoding/json"
	"log"

	"errors"

	"github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/core/service"
	"github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/repositories/rds"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"

	"github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/cmd/ssm"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const (
	onlineStart     = "onlineStart"
	onlineResume    = "onlineResume"
	onlineStop      = "onlineStop"
	onlineFinalize  = "onlineFinalize"
	dateConfirmed   = "dateConfirmed"
	dateUnconfirmed = "dateUnconfirmed"
	operation       = "Operation"
)

func HandleRequest(ctx context.Context, body *events.SQSEvent) (*string, error) {
	logs.Info.Println("Gin cold start")
	records := body.Records

	credsFinder := ssm.NewSSMCredentialsFinder()
	credentials, err := credsFinder.GetDBCredentialsFromDefaultEnv()

	if err != nil {
		logs.Error.Println("we fail to Fetch the envs")
		return nil, err
	}

	connector, err := rds.NewConector(credentials.User, credentials.Password, credentials.Name, credentials.Url)

	if err != nil {
		logs.Error.Println("we fail to create the connector")
		return nil, err
	}

	location, err := rds.NewRDSRepo(connector)

	if err != nil {
		logs.Error.Println("we fail to create the repository")
		return nil, err
	}

	var service ports.Service = service.NewLocationService(location)

	var onError bool = false

	for _, record := range records {
		logs.Info.Printf("The  SQS: %+v\n", record)

		// It is not a SNS message, the body is now the Notification itself
		notification := domain.Notification{}
		err := json.Unmarshal([]byte(record.Body), &notification)

		if err != nil {
			logs.Error.Println("We found an error while marshalling the notification: ", err.Error())
			onError = true
			continue
		}

		err = methodsMap(*record.MessageAttributes[operation].StringValue, notification, service)

		if err != nil {
			logs.Error.Println("We found an error while running the methods map:", err.Error())
			onError = true
			continue
		}

	}

	if onError {
		return nil, errors.New("A error happen during handling the records")
	}

	return nil, nil
}

func methodsMap(subject string, notification domain.Notification, service ports.Service) error {
	log.Printf("methodsMap: \t\nsubject -> %s, \t\nnotification -> %s", subject, notification)

	switch subject {

	case onlineStart:
		date, err := castNotificationToDatesLocation(notification, domain.OnlineDateStatus, domain.Online)

		if err != nil {
			return err
		}

		return service.OnDateCreation(*date)
	case dateConfirmed:
		date, err := castNotificationToDatesLocation(notification, domain.OnlineDateStatus, domain.Scheduled)

		if err != nil {
			return err
		}

		return service.OnDateCreation(*date)
	case onlineFinalize:
		return service.OnDateRemoved(notification.DateId)
	case onlineResume:
		return service.OnDateStateChanged(notification.DateId, domain.OnlineDateStatus)
	case onlineStop:
		return service.OnDateStateChanged(notification.DateId, domain.StoppedDateStatus)
	case dateUnconfirmed:
		return service.OnDateRemoved(notification.DateId)

	}

	return nil
}

func castNotificationToDatesLocation(
	notification domain.Notification,
	dateState domain.DateState,
	dateType domain.DateType,
) (*domain.DatesLocation, error) {

	var placeMap map[string]interface{}

	if notification.Aditionalpayload["place"] != nil {
		placeMap = notification.Aditionalpayload["place"].(map[string]interface{})
	} else if notification.Aditionalpayload["placeInfo"] != nil {
		placeMap = notification.Aditionalpayload["placeInfo"].(map[string]interface{})
	} else {
		return nil, errors.New("Missing place in body")
	}

	lat := placeMap["lat"].(float64)
	lon := placeMap["lon"].(float64)

	return &domain.DatesLocation{
		StateID: notification.DateId,
		Lat:     lat,
		Lon:     lon,
		State: domain.States{
			StateID: dateState,
		},
		Type: domain.Types{
			TypeID: dateType,
		},
	}, nil

}
func init() {

}

func main() {
	lambda.Start(HandleRequest)
}
