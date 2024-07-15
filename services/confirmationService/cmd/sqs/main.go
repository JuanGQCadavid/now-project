package main

import (
	"context"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, body *events.SQSEvent) (string, error) {
	logs.Info.Println("Greetings form the sqs handler my boy!")
	logs.Info.Printf("%+v", body.Records)
	return "Okey", nil
}
