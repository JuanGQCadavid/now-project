package infrachannels

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/JuanGQCadavid/now-project/infra/services/awsServiceChangesNotifier/internal/core/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

const (
	TOPIC_ARN_ENV_NAME string = "infraChangesTopicArn"
	GROUP_ID           string = "InfraEvents"
)

var (
	ErrEnvsMissing = errors.New("Missing env variables")
)

type InfraSNSChangeTopic struct {
	Service  *sns.SNS
	TopicArn string
}

func NewInfraSNSChangeTopic() (*InfraSNSChangeTopic, error) {
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	service := sns.New(session)

	topicArn, isTopicArnEnv := os.LookupEnv(TOPIC_ARN_ENV_NAME)

	if !isTopicArnEnv {
		log.Println("ERROR: Missing Topic arn env")
		return &InfraSNSChangeTopic{}, ErrEnvsMissing
	}

	return &InfraSNSChangeTopic{
		Service:  service,
		TopicArn: topicArn,
	}, nil
}

func (infra *InfraSNSChangeTopic) Publish(body domain.InfraTopicBody) error {

	bodyContent, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return err
	}

	fmt.Println(string(bodyContent))

	input := &sns.PublishInput{
		TopicArn:       aws.String(infra.TopicArn),
		Message:        aws.String(string(bodyContent)),
		MessageGroupId: aws.String(GROUP_ID),
	}
	_, errOnPublish := infra.Service.Publish(input)

	if errOnPublish != nil {
		log.Println("ERROR: ", errOnPublish.Error())
		return errOnPublish
	}
	return nil

}
