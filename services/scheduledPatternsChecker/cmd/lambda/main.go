package main

import (
	"context"
	"log"
	"runtime"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, body *events.SQSEvent) (string, error) {

	for _, record := range body.Records {
		log.Printf("%+v \n", record)
		log.Println("------------")
		log.Printf("%+v \n", record.Body)

		operation := record.MessageAttributes["Operation"]
		log.Println(" ********** ")
		log.Println(operation.StringValue)
	}

	log.Println("Number of CPU", runtime.NumCPU())
	return "Done", nil
}

func main() {
	lambda.Start(Handler)
}
