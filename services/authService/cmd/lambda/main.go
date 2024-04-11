package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func Handler(ctx context.Context, event MyEvent) (string, error) {
	fmt.Println(event)
	return "Hi", nil
}

func main() {
	lambda.Start(Handler)
}
