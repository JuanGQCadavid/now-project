package awssns

import (
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type SNSNotificator struct {
	svc *sns.SNS
}

func NewSNSNotificator(session *session.Session) *SNSNotificator {
	svc := sns.New(session)
	return &SNSNotificator{
		svc: svc,
	}
}

func (not *SNSNotificator) SendNotification(msg, dest string) error {
	logs.Info.Println("*******   Notification   *******")

	params := &sns.PublishInput{
		Message:     aws.String(msg),
		PhoneNumber: aws.String(dest),
	}

	resp, err := not.svc.Publish(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		logs.Error.Println("We fail to send the notificator err: ", err.Error())
		return err
	}

	logs.Info.Println(resp)
	logs.Info.Println("******* END NOTIFICATION *******")

	return nil
}
