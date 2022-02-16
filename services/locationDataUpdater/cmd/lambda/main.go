package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, body *events.SQSEvent) (string, error) {
	log.Println("Hello!")
	log.Printf("%+v", body)
	log.Printf("%+v", ctx)

	log.Printf("%+v", body.Records)

	records := body.Records

	for _, record := range records {
		log.Printf("%+v", record)
		log.Printf("%+v", record.Body)
		log.Printf("%+v", record.MessageAttributes)
	}

	return "Base", nil
}

func main() {
	lambda.Start(HandleRequest)
}
