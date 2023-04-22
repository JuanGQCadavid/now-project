package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/core/domain"
)

const (
	OnlineStart    NotifyOperator = "onlineStart"
	OnlineStop     NotifyOperator = "onlineStop"
	OnlineResume   NotifyOperator = "onlineResume"
	OnlineFinalize NotifyOperator = "onlineFinalize"
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
