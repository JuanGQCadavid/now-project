package service

import (
	"fmt"
	"log"

	"github.com/JuanGQCadavid/now-project/infra/services/awsServiceChangesNotifier/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/infra/services/awsServiceChangesNotifier/internal/repositories/infrachannels"
)

type ServicesChangeNotifier struct {
	infraTopic *infrachannels.InfraSNSChangeTopic
}

func NewServicesChangeNotifier(infraTopic *infrachannels.InfraSNSChangeTopic) *ServicesChangeNotifier {
	return &ServicesChangeNotifier{
		infraTopic: infraTopic,
	}
}

func (svc *ServicesChangeNotifier) OnRDSEvent(event domain.EventNotification) error {
	log.Println(fmt.Sprintf("OnRDSEvent: %+v", event))

	switch event.Detail.EventID {
	case string(domain.DB_INSTANCE_RESTARTED):
		log.Println("event.Detail.EventID: ", string(domain.DB_INSTANCE_RESTARTED))
		err := svc.infraTopic.Publish(domain.InfraTopicBody{
			Title:       "The db instance as being RESTARTED",
			ContentBody: "Yeah dude! Lets rock",
		})

		if err != nil {
			log.Println("ERROR: ", err.Error())
			return nil
		}
		break
	case string(domain.DB_INSTANCE_STARTED):
		log.Println("event.Detail.EventID: ", string(domain.DB_INSTANCE_STARTED))
		err := svc.infraTopic.Publish(domain.InfraTopicBody{
			Title:       "The db instance as being STARTED",
			ContentBody: "Yeah dude! Lets rock",
		})
		if err != nil {
			log.Println("ERROR: ", err.Error())
			return nil
		}
		break
	case string(domain.DB_INSTANCE_STOPED):
		log.Println("event.Detail.EventID: ", string(domain.DB_INSTANCE_STOPED))
		err := svc.infraTopic.Publish(domain.InfraTopicBody{
			Title:       "The db instance as being STOPED",
			ContentBody: "Yeah dude! Lets rock",
		})
		if err != nil {
			log.Println("ERROR: ", err.Error())
			return nil
		}
		break
	default:
		log.Println("DEFAULT - NO MAPPED: event.Detail.EventID: ", event.Detail.EventID)
	}
	return nil
}
