package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, body *events.SQSEvent) (string, error) {
	log.Println("Hello!")
	records := body.Records

	for _, record := range records {
		snsMessage := events.SNSEvent{}
		json.Unmarshal([]byte(record.Body), &snsMessage)

		snsRecords := snsMessage.Records

		for _, snsRecord := range snsRecords {
			fmt.Printf("The message %s", snsRecord.SNS.Message)
		}

	}

	return "Base", nil
}

func main() {
	lambda.Start(HandleRequest)
}
