package spotactivityservices

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type AWSSpotActivityTopic struct {
	snsService *sns.SNS
	targetArn  string
}

func NewAWSSpotActivityTopic() *AWSSpotActivityTopic {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	snsArn, isPresent := os.LookupEnv("snsArn")

	if !isPresent {
		log.Fatal("snsArn is not configured in the env.")
	}

	svc := sns.New(sess)
	return &AWSSpotActivityTopic{
		snsService: svc,
		targetArn:  snsArn,
	}
}

func (r AWSSpotActivityTopic) NotifySpotCreated(spot domain.Spot) error {
	log.Println("NotifySpotCreated: ", "\n\t", " spot: ", fmt.Sprintf("%+v", spot))

	body, err := json.Marshal(&spot)
	log.Println("body -> ", fmt.Sprintf("%+v", body))
	if err != nil {
		return err
	}

	return r.sendMessageToTopic(string(body), "spot_created")
}
func (r AWSSpotActivityTopic) NotifySpotStopped(spotId string) error {
	log.Println("NotifySpotStopped: ", "\n\t", " spotId: ", spotId)

	return r.sendMessageToTopic(spotId, "spot_removed")
}

func (r AWSSpotActivityTopic) sendMessageToTopic(messageBody string, operation string) error {
	log.Println("sendMessageToTopic: ", "\n\t", " Operation: ", operation, "\n\t", " messageBody:", messageBody)

	log.Println("Target arn: ", r.targetArn)
	operationResult, err := r.snsService.Publish(&sns.PublishInput{
		Message: &messageBody,
		Subject: aws.String(operation),
		MessageAttributes: map[string]*sns.MessageAttributeValue{
			"Operation": {
				DataType:    aws.String("String"),
				StringValue: aws.String(operation),
			},
		},
		TargetArn: aws.String(r.targetArn),
	})

	if err != nil {
		log.Print(err)
		log.Print(err.Error())
	} else {
		log.Printf("%+v", operationResult)
	}

	return err
}
