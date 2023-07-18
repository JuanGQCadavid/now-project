package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/confirmationService/internal/core/domain"
)

const (
	DateConfirmed   NotifyOperator = "dateConfirmed"
	DateUnconfirmed NotifyOperator = "dateUnconfirmed"
)

var (
	//Errors
	ErrNotifyNotCreated         = errors.New("We could not create the notifer")
	ErrNotificationFailToBeSent = errors.New("We fail to send the notification")
)

type NotifyOperator string

type Notify interface {
	ConfirmationActivity(operation NotifyOperator, notification domain.Notification) error
}
