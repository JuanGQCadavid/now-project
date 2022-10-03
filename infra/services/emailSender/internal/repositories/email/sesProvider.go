package email

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	CharSet = "UTF-8"
)

type SesEmailProvider struct {
	service *ses.SES
}

func NewSesEmailProvider() *SesEmailProvider {

	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	service := ses.New(session)

	return &SesEmailProvider{
		service: service,
	}
}

func (email *SesEmailProvider) SendEmail(fromAddress string, toAddresses string, htmlBoddy string, textBoddy string, subject string) error {

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(toAddresses),
			},
		},
		Message: &ses.Message{
			Subject: &ses.Content{
				Data:    aws.String(subject),
				Charset: aws.String(CharSet),
			},
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(htmlBoddy),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(textBoddy),
				},
			},
		},
		Source: aws.String(fromAddress),
	}

	out, err := email.service.SendEmail(input)

	if err != nil {
		log.Println("ERROR: There were an error while trying to send the email. \nError: ", err.Error())
		return err
	}

	log.Println("SendEmail output: ", fmt.Sprintf("%+v", out))
	return nil
}
