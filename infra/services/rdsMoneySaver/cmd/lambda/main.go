package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/JuanGQCadavid/now-project/infra/services/rdsMoneySaver/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/infra/services/rdsMoneySaver/internal/core/service"
	"github.com/JuanGQCadavid/now-project/infra/services/rdsMoneySaver/internal/repositories/rds"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(event domain.ServiceEvent) error {
	log.Println(fmt.Sprintf("HandleRequest - Request: %+v", event))

	rdsRepo := rds.NewRDSRepository()
	service := service.NewMoneySaver(*rdsRepo)

	switch event.Action {
	case domain.START:
		return service.StartDatabases(event.DBInstance)
	case domain.STOP:
		return service.StopDatabases(event.DBInstance)
	}

	return errors.New("Event action not found. EventAction: " + string(event.Action))
}

func main() {
	lambda.Start(HandleRequest)
}
