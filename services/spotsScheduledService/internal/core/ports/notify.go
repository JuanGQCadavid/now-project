package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/domain"
)

const (
	SchedulePatternAppended  NotifyOperator = "schedulePatternAppended"
	SchedulePatternConcluded NotifyOperator = "schedulePatternConcluded"
	SchedulePatternResumed   NotifyOperator = "schedulePatternResumed"
	SchedulePatternFreezed   NotifyOperator = "schedulePatternFreezed"
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
