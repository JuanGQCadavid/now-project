package snstopic

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

var (
	ErrSNSArnEnvNotFound   = errors.New("The sns arn env is empty or does not exist")
	ErrOnSnsPublishMessage = errors.New("We found a problem while sending the message")
)

type SNSTopicActions struct {
	snsService *sns.SNS
	targetArn  string
}

func NewSNSTopicActionsFromEnv(envName string) (*SNSTopicActions, error) {
	snsArn, isPresent := os.LookupEnv(envName)
	if !isPresent {
		logs.Error.Println("snsArn is not configured in the env.")
		return nil, ErrSNSArnEnvNotFound
	}
	return NewSNSTopicActionsFromArn(snsArn)
}

func NewSNSTopicActionsFromArn(snsArn string) (*SNSTopicActions, error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)
	return &SNSTopicActions{
		snsService: svc,
		targetArn:  snsArn,
	}, nil

}

func (r SNSTopicActions) NotifyWithBody(action string, body interface{}) error {
	logs.Info.Printf("NotifyWithBody: action ->  %s", action)

	bodyMarshaled, err := json.Marshal(&body)
	logs.Info.Println("body -> ", fmt.Sprintf("%+v", bodyMarshaled))

	if err != nil {
		logs.Error.Println(err.Error())
		return err
	}

	return r.sendMessageToTopic(string(bodyMarshaled), action)
}

// func (r SNSTopicActions) Notify(action string,,spotId string) error {
// 	log.Println("NotifySpotStopped: ", "\n\t", " spotId: ", spotId)

// 	return r.sendMessageToTopic(spotId, "spot_removed")
// }

func (r SNSTopicActions) sendMessageToTopic(messageBody string, operation string) error {
	logs.Info.Printf("sendMessageToTopic: \n\toperation: %s, \n\tmessagebody: %s", operation, messageBody)
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
		logs.Error.Println("We found an error while sending the message to the topic: \n", err.Error())
		return ErrOnSnsPublishMessage
	}

	logs.Info.Printf("%+v\n", operationResult)

	return err
}
