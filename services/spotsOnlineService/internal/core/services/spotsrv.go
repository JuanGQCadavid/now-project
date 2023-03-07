package services

import (
	"log"
	"time"

	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/core/ports"
	"github.com/google/uuid"
)

type Service struct {
	repository ports.Repository
}

func NewService(repository ports.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Stop(spotId string, requestUserId string)     {}
func (s *Service) Resume(spotId string, requestUserId string)   {}
func (s *Service) Finalize(spotId string, requestUserId string) {}

/*
	The user should not be hosting more than two events at the same time

	Procedure:
		1. Verify that the user is the owner of the spot. --- OK
			If no -> Return an error

		2. Check that the spot does not have a current online event.
			If so -> Return and error

		3. Create Date - Spot association with the RequestUserId as the IS_HOSTING relation

*/

func (s *Service) Start(spotId string, requestUserId string, durationApproximated int64, maximunCapacity int) (domain.OnlineSpot, error) {
	log.Printf("Service Start: spotId %s, requestUserId: %s \n", spotId, requestUserId)

	// 1. Verify that the user is the owner of the spot
	onlineSpot, err := s.fetchAndVerifySpot(spotId, requestUserId)

	if err != nil {
		log.Println("We found an error while verifying the spot\n\t\t", err.Error())
		return domain.OnlineSpot{}, err
	}

	// 2. Check that the spot does not have a current online event.
	if len(onlineSpot.DatesInfo) > 0 {
		log.Printf("DateInfo is not empty, so the event is already online: spotDate %+v \n", onlineSpot.DatesInfo)
		return domain.OnlineSpot{}, ports.ErrSpotIsAlreadyOnline
	}

	creationTime := time.Now().String()

	dateInfo := domain.SpotDate{
		DateId:                        s.generateUUID(),
		DurationApproximatedInSeconds: durationApproximated,
		StartTime:                     creationTime,
		Confirmed:                     true,
		MaximunCapacty:                maximunCapacity,
		Date:                          creationTime,
		HostInfo: domain.HostInfo{
			HostId: requestUserId,
		},
	}

	onlineSpot.DatesInfo = []domain.SpotDate{
		dateInfo,
	}

	// 3. Create Date - Spot association with the RequestUserId as the IS_HOSTING relation
	dateId, err := s.repository.AssociateDateWithSpot(onlineSpot)

	if err != nil {
		log.Println("An error happen while associating the date with the spot, error: ", err.Error())
		// log.Printf("OnlineSpot: %+v \n")
		return domain.OnlineSpot{}, err
	}

	dateInfo.DateId = dateId

	return onlineSpot, nil

}

func (s *Service) fetchAndVerifySpot(spotId string, requestUserId string) (domain.OnlineSpot, error) {
	log.Printf("Service fetchAndVerifySpot: spotId %s, requestUserId: %s \n", spotId, requestUserId)

	onlineSpot, err := s.repository.FetchOnlineSpot(spotId)

	if err != nil {
		log.Println("We found an error while fetching the spot owner\n\t\t", err.Error())
		return domain.OnlineSpot{}, err
	}

	if len(onlineSpot.SpotInfo.SpotId) == 0 {
		log.Println("The spot does not exist")
		return domain.OnlineSpot{}, ports.ErrSpotNotFound
	}

	if onlineSpot.SpotInfo.OwnerId != requestUserId {
		log.Println("The owner id is differente than the spot owner")
		return domain.OnlineSpot{}, ports.ErrUserIsNotTheOwner
	}

	return onlineSpot, nil
}

func (s *Service) generateUUID() string {
	return uuid.New().String()
}

// func (s *Service) CreateSpot(spot domain.Spot) (domain.Spot, error) {
// 	// TODO -> Missing refactor from go online to createSpot

// 	log.Println("GoOnline -> ", fmt.Sprintf("%+v", spot))
// 	//TODO -> Missing body validation

// 	// TODO -> I think this is not working as I spect, we need to check this.
// 	if returnedSpot, returnedError := s.GetSpotByUserId(spot.HostInfo.Id); returnedError == nil {
// 		if err := s.EndSpot(returnedSpot.EventInfo.UUID); err != nil {
// 			return domain.Spot{}, err
// 		}
// 	}

// 	uuid := s.uuidGen.New()
// 	spot.EventInfo.UUID = uuid

// 	if returnedError := s.createEvent(spot); returnedError != nil {
// 		return domain.Spot{}, returnedError
// 	}

// 	// Check if method contains tags for the event

// 	if spot.TopicsInfo.PrincipalTopic != "" || spot.TopicsInfo.SecondaryTopics != nil {
// 		if returnedError := s.createSpotTags(spot); returnedError != nil {
// 			return domain.Spot{}, returnedError
// 		}
// 	}

// 	return spot, nil
// }
