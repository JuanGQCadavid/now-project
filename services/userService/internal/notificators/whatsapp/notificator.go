package whatsapp

import (
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type WhatsAppNotificator struct {
	svc *sns.SNS
}

func NewSNSNotificator(session *session.Session) *WhatsAppNotificator {
	svc := sns.New(session)
	return &WhatsAppNotificator{
		svc: svc,
	}
}

func (not *WhatsAppNotificator) SendNotification(msg, dest string) error {
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

// curl -i -X POST \
//   https://graph.facebook.com/v17.0/134248176448539/messages \
//   -H 'Authorization: Bearer EAAJiuCegkRUBOwhkLZCe9PJ2KtawoVQS2j2I7nxpTmGRtOXY7wsmfKZCWAOKiirCrSS84et9qKH74gK3yj0oYGhJyjFp7yZCbcsDh4EVyAaOzebHK42OOh0gJgKkIVdBTclusZCZCtsARQ975di8esbqGkhvALSirZCESoPnVf6PxJJiTLrkqoGJXdiaL7KZBic8WB18QRcYloiEseybgSVMZAYMI1HoAUoZD' \
//   -H 'Content-Type: application/json' \
//   -d '{ "messaging_product": "whatsapp", "to": "573013475995", "text": {"body" : "This is a test"}}'

// curl -X POST \
//   'https://graph.facebook.com/v18.0/134248176448539/messages' \
//   -H 'Authorization: Bearer EAAJiuCegkRUBOwhkLZCe9PJ2KtawoVQS2j2I7nxpTmGRtOXY7wsmfKZCWAOKiirCrSS84et9qKH74gK3yj0oYGhJyjFp7yZCbcsDh4EVyAaOzebHK42OOh0gJgKkIVdBTclusZCZCtsARQ975di8esbqGkhvALSirZCESoPnVf6PxJJiTLrkqoGJXdiaL7KZBic8WB18QRcYloiEseybgSVMZAYMI1HoAUoZD' \
//   -d '{
//     "messaging_product": "whatsapp",
//     "to": "573235237844",
//     "text": {"body" : "hi"}
//    }'
