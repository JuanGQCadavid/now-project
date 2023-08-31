package main

import (
	"fmt"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	snstopic "github.com/JuanGQCadavid/now-project/services/pkgs/common/snstopic"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/sqsqueue"
)

func main() {
	testSQS()
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

func testSQS() {
	actions, err := sqsqueue.NewSQSTopicActionsFromArn("sendConfirmationSQS")

	if err != nil {
		logs.Error.Fatalln(err.Error())
	}

	limit := 2

	bodies := make([]interface{}, limit)

	for i := 0; i < limit; i++ {
		body := map[string]string{
			"spotId":  "Dude!",
			"counter": fmt.Sprintf("%d", i),
		}

		bodies[i] = body
	}

	errs := actions.SendBulkMessages(bodies)

	if errs != nil && len(errs) > 0 {
		logs.Error.Println("There is erros dude!")
		for key, err := range errs {
			logs.Error.Printf("%+v - %s \n", key, err.Error())
		}
	}
}

func testSingleSQS() {
	actions, err := sqsqueue.NewSQSTopicActionsFromArn("sendConfirmationSQS")

	if err != nil {
		logs.Error.Fatalln(err.Error())
	}

	body := map[string]string{
		"spotId": "Dude!",
	}

	actions.SendMessage(body)
}
