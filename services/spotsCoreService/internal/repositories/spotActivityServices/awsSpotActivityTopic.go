package spotactivityservices

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
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
		logs.Error.Fatal("snsArn is not configured in the env.")
	}

	svc := sns.New(sess)
	return &AWSSpotActivityTopic{
		snsService: svc,
		targetArn:  snsArn,
	}
}

func (r AWSSpotActivityTopic) NotifySpotCreated(spot domain.Spot) error {
	logs.Info.Println("NotifySpotCreated: ", "\n\t", " spot: ", fmt.Sprintf("%+v", spot))

	body, err := json.Marshal(&spot)
	logs.Info.Println("body -> ", fmt.Sprintf("%+v", body))
	if err != nil {
		return err
	}

	return r.sendMessageToTopic(string(body), "spot_created")
}
func (r AWSSpotActivityTopic) NotifySpotStopped(spotId string) error {
	logs.Info.Println("NotifySpotStopped: ", "\n\t", " spotId: ", spotId)

	return r.sendMessageToTopic(spotId, "spot_removed")
}

func (r AWSSpotActivityTopic) sendMessageToTopic(messageBody string, operation string) error {
	logs.Info.Println("sendMessageToTopic: ", "\n\t", " Operation: ", operation, "\n\t", " messageBody:", messageBody)

	logs.Info.Println("Target arn: ", r.targetArn)
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
		logs.Error.Println(err)
		logs.Error.Println(err.Error())
	} else {
		logs.Info.Printf("%+v", operationResult)
	}

	return err
}
