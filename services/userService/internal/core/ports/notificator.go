package ports

type Notificator interface {
	SendNotification(msg, dest string) error
}
