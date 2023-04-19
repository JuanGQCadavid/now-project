package main

import (
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	snstopic "github.com/JuanGQCadavid/now-project/services/pkgs/common/snstopic"
)

func main() {
	testSns()
}

func testSns() {
	actions, err := snstopic.NewSNSTopicActionsFromEnv("snsArn")

	if err != nil {
		logs.Error.Fatalln(err.Error())
	}

	body := map[string]string{
		"spotId": "Dude!",
	}

	actions.NotifyWithBody("spotScheduleAdded", body)
}
