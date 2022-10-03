package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request interface{}) error {
	log.Println(fmt.Sprintf("Request: \n %+v", request))

	return nil
}

func main() {
	lambda.Start(Handler)
}
