package lambda

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, body interface{}) (string, error) {
	log.Println("Hello!")
	log.Printf("%+v", body)
	log.Printf(body.(string))
	log.Printf("%+v", ctx)
	return body.(string), nil
}

func main() {
	lambda.Start(HandleRequest)
}
