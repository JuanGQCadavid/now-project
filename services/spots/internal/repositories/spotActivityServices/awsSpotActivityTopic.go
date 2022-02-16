package spotactivityservices

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/JuanGQCadavid/now-project/services/spots/internal/core/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type AWSSpotActivityTopic struct {
	sqsService *sns.SNS
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
		sqsService: svc,
		targetArn:  snsArn,
	}
}

func (r AWSSpotActivityTopic) AppendSpot(spot domain.Spot) error {
	log.Println("AppendSpot: ", "\n\t", " spot: ", fmt.Sprintf("%+v", spot))

	body, err := json.Marshal(&spot)
	log.Println("body -> ", fmt.Sprintf("%+v", body))
	if err != nil {
		return err
	}

	return r.sendMessageToTopic(string(body), "spot_created")
}
func (r AWSSpotActivityTopic) RemoveSpot(spotId string) error {
	log.Println("RemoveSpot: ", "\n\t", " spotId: ", spotId)

	return r.sendMessageToTopic(spotId, "spot_removed")
}

func (r AWSSpotActivityTopic) sendMessageToTopic(messageBody string, operation string) error {
	log.Println("sendMessageToTopic: ", "\n\t", " Operation: ", operation, "\n\t", " messageBody:", messageBody)

	operationResult, err := r.sqsService.Publish(&sns.PublishInput{
		Message: &messageBody,
		Subject: aws.String(operation),
		MessageAttributes: map[string]*sns.MessageAttributeValue{
			"Operation": {
				DataType:    aws.String("String"),
				StringValue: aws.String(operation),
			},
		},
		TargetArn: &r.targetArn,
	})

	if err != nil {
		log.Print(err)
	} else {
		log.Printf("%+v", operationResult)
	}

	return err
}
