package main

import (
	"context"
	"log"
	"runtime"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, body *events.SQSEvent) (string, error) {
	log.Println("HiiiI!")

	log.Println("Number of CPU", runtime.NumCPU())
	return "Done", nil
}

func main() {
	lambda.Start(Handler)
}
