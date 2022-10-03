package domain

type NotificationRequest struct {
	Title        string
	Body         string
	SendWhatsApp bool
	SendEmail    bool
}
