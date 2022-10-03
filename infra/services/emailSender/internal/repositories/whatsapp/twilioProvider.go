package whatsapp

import (
	"fmt"
	"log"

	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

const (
	FROM_NUMER = "+14155238886"
)

type WhatsAppProvider struct {
	// accountId string
	// authToken string
	client *twilio.RestClient
}

func NewWhatsAppProvider() *WhatsAppProvider {

	client := twilio.NewRestClient()

	return &WhatsAppProvider{
		client: client,
	}
}

func (provider *WhatsAppProvider) SendMessage(number string, body string) error {

	params := &openapi.CreateMessageParams{}
	params.SetTo(fmt.Sprintf("whatsapp:%s", number))
	params.SetFrom(fmt.Sprintf("whatsapp:%s", FROM_NUMER))
	params.SetBody(body)

	out, err := provider.client.Api.CreateMessage(params)

	if err != nil {
		log.Println("SendMessage error on sending message: ", err.Error())
		return err
	}

	log.Println("SendMessage output: ", fmt.Sprintf("%+v", out))
	return nil

}
