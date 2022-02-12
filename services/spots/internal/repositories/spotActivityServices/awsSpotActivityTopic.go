package spotactivityservices

import (
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

	svc := sns.New(sess)
	return &AWSSpotActivityTopic{
		sqsService: svc,
		targetArn:  "arn:aws:sns:us-east-2:732596568988:spotActivityTopic",
	}
}
func (r AWSSpotActivityTopic) AppendSpot(spotId string) error {

	var message = "dude"

	_, err := r.sqsService.Publish(&sns.PublishInput{
		Message:   &message,
		TargetArn: &r.targetArn,
	})

	if err != nil {
		println(err)
	}

	return err
}
func (r AWSSpotActivityTopic) RemoveSpot(spotId string) error {
	return nil
}
