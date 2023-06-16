package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/domain"
)

const (
	DateIdDeleted NotifyOperator = "DateIdDeleted"
)

var (
	//Errors
	ErrNotifyNotCreated         = errors.New("We could not create the notifer")
	ErrNotificationFailToBeSent = errors.New("We fail to send the notification")
)

type NotifyOperator string

type Notify interface {
	SchedulePatternActivity(operation NotifyOperator, notification domain.Notification) error
}
