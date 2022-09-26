package infrachangestopic

import (
	"errors"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

const (
	TOPIC_ARN_ENV_NAME string = "infraChangesTopicArn"
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

func (infra *InfraSNSChangeTopic) Publish() {

	input := &sns.PublishInput{
		TopicArn: aws.String(infra.TopicArn),
		Message:  aws.String("Missing"),
	}

	infra.Service.Publish(input)
}
