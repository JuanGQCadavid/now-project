package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/JuanGQCadavid/now-project/infra/services/awsServiceChangesNotifier/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/infra/services/awsServiceChangesNotifier/internal/core/service"
	"github.com/JuanGQCadavid/now-project/infra/services/awsServiceChangesNotifier/internal/repositories/email"
	"github.com/JuanGQCadavid/now-project/infra/services/awsServiceChangesNotifier/internal/repositories/whatsapp"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(body *events.SQSEvent) error {
	log.Println(fmt.Sprintf("body: %+v", body))
	emailProvider := email.NewSesEmailProvider()
	whatsappProvider := whatsapp.NewWhatsAppProvider()
	srv := service.NewNotificationService(emailProvider, whatsappProvider)

	records := body.Records

	for _, record := range records {

		bodyMesssge := domain.LambdaRequest{}
		err := json.Unmarshal([]byte(record.Body), &bodyMesssge)

		if err != nil {
			log.Fatalln("Error while doing unmarshal: error ->", err.Error())
		}

		log.Println(fmt.Sprintf("bodyMesssge: %+v", bodyMesssge))

		srv.SendNotification(domain.NotificationRequest{
			Title:        bodyMesssge.Title,
			Body:         bodyMesssge.ContentBody,
			SendWhatsApp: true,
			SendEmail:    true,
		})

	}

	log.Println(fmt.Sprintf("Request : \n %+v", body))

	return nil
}

func main() {
	lambda.Start(Handler)
}
