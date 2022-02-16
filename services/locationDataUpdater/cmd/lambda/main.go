package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, body *events.SNSEvent) (string, error) {
	log.Println("Hello!")
	log.Printf("%+v", body)
	log.Printf("%+v", ctx)

	log.Printf("%+v", body.)

	return "Base", nil
}

func main() {
	lambda.Start(HandleRequest)
}
