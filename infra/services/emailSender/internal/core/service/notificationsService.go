package service

import (
	"fmt"
	"log"
	"os"

	"github.com/JuanGQCadavid/now-project/infra/services/awsServiceChangesNotifier/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/infra/services/awsServiceChangesNotifier/internal/repositories/email"
	"github.com/JuanGQCadavid/now-project/infra/services/awsServiceChangesNotifier/internal/repositories/whatsapp"
)

const (
	PHONE_NUMBER  string = "+573013475995"
	FROM_ENV             = "from"
	TO_ENV               = "to"
	EMAIL_SUBJECT        = "Infra notifications"
)

type NotificationService struct {
	emailProvider    *email.SesEmailProvider
	whatsAppProvider *whatsapp.WhatsAppProvider
	emailFrom        string
	emailTo          string
}

func NewNotificationService(emailProvider *email.SesEmailProvider, whatsAppProvider *whatsapp.WhatsAppProvider) *NotificationService {

	emailFrom, isEmailFrom := os.LookupEnv(FROM_ENV)
	emailTo, isEmailTo := os.LookupEnv(TO_ENV)

	if !isEmailFrom || !isEmailTo {
		log.Println("Missing envs, emailFrom:", emailFrom, " - emailTo:", emailTo)
	}

	return &NotificationService{
		emailProvider:    emailProvider,
		whatsAppProvider: whatsAppProvider,
		emailFrom:        emailFrom,
		emailTo:          emailTo,
	}
}

func (svc *NotificationService) SendNotification(request domain.NotificationRequest) {
	log.Println(fmt.Sprintf("SendNotification: body -> %+v ", request))

	if request.SendEmail {
		log.Println("On request.SendEmail")
		htmlBody, txtBody := svc.GenerateBodies(request.Title, request.Body)
		log.Println(fmt.Sprintf("htmlBody: %s, txtBody: %s", htmlBody, txtBody))
		svc.emailProvider.SendEmail(svc.emailFrom, svc.emailTo, htmlBody, txtBody, EMAIL_SUBJECT)
	}

	if request.SendWhatsApp {
		log.Println("On request.SendWhatsApp")
		body := svc.GenerateMessageBody(request.Title, request.Body, EMAIL_SUBJECT)
		log.Println(fmt.Sprintf("body: %s", body))
		svc.whatsAppProvider.SendMessage(PHONE_NUMBER, body)
	}
}

func (svc *NotificationService) GenerateMessageBody(title string, body string, subject string) string {

	textBody := fmt.Sprintf("*%s*,\n\nGood morning!\n%s\n\n%s", subject, title, body)

	return textBody
}

func (svc *NotificationService) GenerateBodies(title string, body string) (string, string) {

	htmlBody := fmt.Sprintf(`
		<h3> Good morning! </h3>
		<h2> %s </h2> 

		<p> %s </p>

	`, title, body)

	textBody := fmt.Sprintf(`
		Good morning!
		%s 
		------
		%s
	`, title, body)

	return htmlBody, textBody
}

// map[
// 	Records:[
// 		map[
// 			attributes:
// 				map[
// 					ApproximateFirstReceiveTimestamp:1664763637959
// 					ApproximateReceiveCount:1
// 					MessageDeduplicationId:23e43791d3d68bdb3a22ad9ff0fa299e4358f1e8db1718d86e22660081ce2610
// 					MessageGroupId:InfraEvents
// 					SenderId:AIDARIMVCWVEZ3VL54FDA
// 					SentTimestamp:1664763637959
// 					SequenceNumber:18872923565027055616
// 				]
// 			awsRegion:us-east-2
// 			body:{
// 				"title": "The db instance as being STARTED",
// 				"contentBody": "Yeah dude! Lets rock"
// 			}
// 		eventSource:aws:sqs
// 		eventSourceARN:arn:aws:sqs:us-east-2:732596568988:emailNotifierSQS.fifo
// 		md5OfBody:ef8d8deb82e25c8ae39b571ca7529a5e
// 		messageAttributes:map[]
// 		messageId:c24aae46-98a6-42b5-bc7e-d44f8e4329b5
// 		receiptHandle:AQEBGebqvYUasX2jsp97ot3EUL9qRMrv0PTxIak3renUNH7QiqkHHWyUpN8MJVq8Xg67TQmlXuBaHEtNcwf17QQdn2iKErfNx3+dTLuaatbHKANiUlbf0/7vCosBE/vrrxUEOx+1EnTPtHu85QIpuGpyj+zjtETt/B1SEdXcAccScRdUGKtAm/TvqkrW2+2wmXBpOqBX5L+syNgcwmjyhbtYY/EKEWtQ2vpHZ0GlThfdHPv+hR3liraqAk1lB4g/ultThEAwk40Iwk80ndcuXD6l7+Gnd8lbdFcYX5faghLUWLc=
// 		]
// 	]
// ]
