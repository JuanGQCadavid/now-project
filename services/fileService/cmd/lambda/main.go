package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func init() {

}

func handleRequest(ctx context.Context, event json.RawMessage) error {
	log.Printf("Tere!")
	return nil
}

func main() {
	lambda.Start(handleRequest)
}
