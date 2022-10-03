package main

// import (
// 	"github.com/JuanGQCadavid/now-project/infra/services/awsServiceChangesNotifier/internal/repositories/email"
// )

// import (
// 	"github.com/JuanGQCadavid/now-project/infra/services/awsServiceChangesNotifier/internal/repositories/whatsapp"
// )

import (
	"github.com/JuanGQCadavid/now-project/infra/services/awsServiceChangesNotifier/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/infra/services/awsServiceChangesNotifier/internal/core/service"
	"github.com/JuanGQCadavid/now-project/infra/services/awsServiceChangesNotifier/internal/repositories/email"
	"github.com/JuanGQCadavid/now-project/infra/services/awsServiceChangesNotifier/internal/repositories/whatsapp"
)

func main() {
	emailProvider := email.NewSesEmailProvider()

	// emailProvider.SendEmail("now.inc.project@gmail.com", "jquirozcadavid@gmail.com", "<hi>Test 2</h1>", "HI DUDE!", "Testing SES two")

	whatsappProvider := whatsapp.NewWhatsAppProvider()
	// whatsappProvider.SendMessage("+573013475995", "Ey Guapo! Como estas?")

	srv := service.NewNotificationService(emailProvider, whatsappProvider)

	srv.SendNotification(domain.NotificationRequest{
		Title:     "This is the title",
		Body:      "This is the body",
		SendEmail: true,
	})

}
