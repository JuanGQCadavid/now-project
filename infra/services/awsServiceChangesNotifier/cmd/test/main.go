package main

import (
	"github.com/JuanGQCadavid/now-project/infra/services/awsServiceChangesNotifier/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/infra/services/awsServiceChangesNotifier/internal/repositories/infrachannels"
)

func main() {
	sns, err := infrachannels.NewInfraSNSChangeTopic()

	if err != nil {
		panic(err)
	}

	sns.Publish(domain.InfraTopicBody{
		Title:       "The db instance as being STOPED",
		ContentBody: "Yeah dude! Lets rock",
	})

}
