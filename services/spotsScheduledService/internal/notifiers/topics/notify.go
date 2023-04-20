package topics

import (
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/snstopic"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/ports"
)

type Notifier struct {
	actions *snstopic.SNSTopicActions
}

func NewNotifierFromEnv(envName string) (*Notifier, error) {

	actions, err := snstopic.NewSNSTopicActionsFromEnv(envName)

	if err != nil {
		logs.Error.Println("We got and erro while attempting to create the topic, error: ", err.Error())
		return nil, err
	}

	return &Notifier{
		actions: actions,
	}, nil
}

func (notifier *Notifier) SchedulePatternActivity(operation ports.NotifyOperator, notification domain.Notification) error {
	logs.Info.Printf("Operation %s, Notification: %+v \n", string(operation), notification)

	err := notifier.actions.NotifyWithBody(string(operation), notification)

	if err != nil {
		logs.Error.Println(ports.ErrNotificationFailToBeSent, err.Error())
		return ports.ErrNotificationFailToBeSent
	}
	return nil
}
