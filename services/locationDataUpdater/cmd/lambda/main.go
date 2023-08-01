package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/service"
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
)

func HandleRequest(ctx context.Context, body *events.SQSEvent) (string, error) {
	records := body.Records

	var service ports.Service = service.NewLocationService()

	for _, record := range records {

		snsMessage := events.SNSEntity{}
		json.Unmarshal([]byte(record.Body), &snsMessage)

		methodsMap(snsMessage.Subject, snsMessage.Message, service)

	}

	return "Base", nil
}

func methodsMap(method string, body string, service ports.Service) error {
	log.Printf("methodsMap: \t\nmethod -> %s, \t\nBody -> %s", method, body)

	switch method {

	case onlineStart:
		spot := domain.Spot{}
		json.Unmarshal([]byte(body), &spot)

		// TODO -> perform some checks
		if err := service.OnSpotCreation(spot); err != nil {
			log.Println("And error! : ", err)
			return err
		}

	case onlineFinalize:
		// TODO -> perform some checks
		if err := service.OnSpotDeletion(body); err != nil {
			log.Println("And error! : ", err)
			return err
		}
	case onlineResume:
	case onlineStop:
	case dateConfirmed:
	case dateUnconfirmed:

	}

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
