package localnotificator

import "github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"

type LocalNotificator struct {
}

func (LocalNotificator) SendNotification(msg, dest string) error {
	logs.Info.Println("*******   Notification   *******")
	logs.Info.Println(dest, msg)
	logs.Info.Println("******* END NOTIFICATION *******")
	return nil
}
