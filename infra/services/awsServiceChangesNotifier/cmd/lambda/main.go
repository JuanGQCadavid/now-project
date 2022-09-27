package main

import (
	"fmt"
	"log"

	"github.com/JuanGQCadavid/now-project/infra/services/awsServiceChangesNotifier/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/infra/services/awsServiceChangesNotifier/internal/core/service"
	"github.com/JuanGQCadavid/now-project/infra/services/awsServiceChangesNotifier/internal/repositories/infrachannels"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(event domain.EventNotification) error {
	log.Println("HandleRequest: ", fmt.Sprintf("%+v", event))

	sns, err := infrachannels.NewInfraSNSChangeTopic()

	if err != nil {
		panic(err)
	}

	svc := service.NewServicesChangeNotifier(sns)

	return svc.OnRDSEvent(event)
}

func main() {
	lambda.Start(HandleRequest)
}
