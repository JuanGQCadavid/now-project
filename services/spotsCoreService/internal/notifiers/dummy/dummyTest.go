package dummy

import (
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/ports"
)

type Notifier struct {
}

func (notifier *Notifier) Send(operation ports.NotifyOperator, notification domain.Notification) error {
	logs.Info.Printf("Operation %s, Notification: %+v \n", string(operation), notification)
	return nil
}
