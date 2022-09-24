package main

import (
	"fmt"
	"log"

	"github.com/JuanGQCadavid/now-project/infra/services/rdsMoneySaver/internal/core/domain"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(event domain.ServiceEvent) error {
	log.Println("On handle Request")
	log.Println(fmt.Sprintf("%+v", event))
	log.Println("END")

	fmt.Println("Hi Dude")
	return nil //errors.New("I'm working?")
}

func main() {
	lambda.Start(HandleRequest)
}
