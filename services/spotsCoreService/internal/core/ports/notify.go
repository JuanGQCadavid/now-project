package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"
)

const (
	SpotCreated NotifyOperator = "spotCreated"
	SpotUpdated NotifyOperator = "spotEdited"
	SpotDeleted NotifyOperator = "spotDeleted"
)

var (
	//Errors
	ErrNotifyNotCreated         = errors.New("We could not create the notifer")
	ErrNotificationFailToBeSent = errors.New("We fail to send the notification")
)

type NotifyOperator string

type Notify interface {
	Send(operation NotifyOperator, notification domain.Notification) error
}
