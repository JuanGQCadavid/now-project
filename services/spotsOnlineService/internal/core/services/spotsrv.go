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
	notifer    ports.Notify
}

func NewService(repository ports.Repository, notifer ports.Notify) *Service {
	return &Service{
		repository: repository,
		notifer:    notifer,
	}
}

func (s *Service) filterDates(spots domain.OnlineSpot, flags domain.SpotStateFlags) []domain.SpotDate {
	datesToReturn := make([]domain.SpotDate, 0)

	for _, date := range spots.DatesInfo {
		to_append := false
		if (date.State.Status == domain.ONLINE_SPOT) && (flags&domain.FlagOnline == domain.FlagOnline) {
			to_append = true
		}
		if (date.State.Status == domain.FINALIZED_SPOT) && (flags&domain.FlagFinalized == domain.FlagFinalized) {
			to_append = true
		}
		if (date.State.Status == domain.PAUSED_SPOT) && (flags&domain.FlagPaused == domain.FlagPaused) {
			to_append = true
		}

		if to_append {
			datesToReturn = append(datesToReturn, date)
		}
	}
	return datesToReturn

}
func (s *Service) fetchSpotsByStatus(spotId string, requestUserId string, flags domain.SpotStateFlags) (domain.OnlineSpot, error) {

	spots, err := s.repository.FetchSpots(spotId)

	if err != nil {
		log.Println("We found an error while fetching the spot owner\n\t\t", err.Error())
		return domain.OnlineSpot{}, err
	}

	if len(spots.SpotInfo.SpotId) == 0 {
		log.Println("The spot does not exist")
		return domain.OnlineSpot{}, ports.ErrSpotNotFound
	}

	if spots.SpotInfo.OwnerId != requestUserId {
		log.Println("The owner id is differente than the spot owner")
		return domain.OnlineSpot{}, ports.ErrUserIsNotTheOwner
	}

	spots.DatesInfo = s.filterDates(spots, flags)
	return spots, nil
}

func (s *Service) GetDates(spotId string, flags domain.SpotStateFlags) (domain.OnlineSpot, error) {
	log.Printf("Service GetDates: spotId %s, flags: %+v \n", spotId, flags)
	spots, err := s.repository.FetchSpots(spotId)

	if err != nil {
		log.Println("We found an error while fetching the spot owner\n\t\t", err.Error())
		return domain.OnlineSpot{}, err
	}

	if len(spots.SpotInfo.SpotId) == 0 {
		log.Println("The spot does not exist")
		return domain.OnlineSpot{}, ports.ErrSpotNotFound
	}

	spots.DatesInfo = s.filterDates(spots, flags)
	return spots, err
}

func (s *Service) Finalize(spotId string, requestUserId string) error {
	log.Printf("Service Finalize: spotId %s, requestUserId: %s \n", spotId, requestUserId)

	// 1. Verify that the user is the owner of the spot
	spots, err := s.fetchSpotsByStatus(spotId, requestUserId, domain.FlagOnline|domain.FlagPaused)

	if err != nil {
		log.Println("We found an error while verifying the spot\n\t\t", err.Error())
		return err
	}

	var dateIdToStop string
	for _, date := range spots.DatesInfo {
		if date.HostInfo.HostId == requestUserId {
			dateIdToStop = date.DateId
		}
	}

	if len(dateIdToStop) == 0 {
		return ports.ErrUserIsNotHostingAnDate
	}

	if err = s.repository.FinalizeDateOnSpot(spotId, dateIdToStop); err != nil {
		log.Printf("We foinf an error while finalizing the date %s on spot %s \n", spotId, dateIdToStop)
		return err
	}

	s.notifer.SchedulePatternActivity(ports.OnlineFinalize, domain.Notification{
		DateId: dateIdToStop,
		SpotId: spotId,
		UserId: requestUserId,
	})

	return nil

}
func (s *Service) Resume(spotId string, requestUserId string) error {
	log.Printf("Service Resume: spotId %s, requestUserId: %s \n", spotId, requestUserId)

	// 1. Verify that the user is the owner of the spot
	pausedSpot, err := s.fetchAndVerifySpot(spotId, requestUserId, domain.PAUSED_SPOT)

	if err != nil {
		log.Println("We found an error while verifying the spot\n\t\t", err.Error())
		return err
	}

	var dateIdToStop string
	for _, date := range pausedSpot.DatesInfo {
		if date.HostInfo.HostId == requestUserId {
			dateIdToStop = date.DateId
		}
	}

	if len(dateIdToStop) == 0 {
		return ports.ErrUserDoesNotHaveStoppedDate
	}

	if err = s.repository.ResumeDateOnSpo(spotId, dateIdToStop); err != nil {
		log.Printf("We foinf an error while resuming the date %s on spot %s \n", spotId, dateIdToStop)
		return err
	}

	s.notifer.SchedulePatternActivity(ports.OnlineResume, domain.Notification{
		SpotId: spotId,
		DateId: dateIdToStop,
		UserId: requestUserId,
	})

	return nil
}

/*
	The user is going to stop the event he is hosting at the spot.

	Procedure:

	1. Verify that he is the owner of the spot

	2. Check the dates associated, for now just stop the one
		associated with is id.
*/

func (s *Service) Stop(spotId string, requestUserId string) error {
	log.Printf("Service Stop: spotId %s, requestUserId: %s \n", spotId, requestUserId)

	// 1. Verify that the user is the owner of the spot
	onlineSpot, err := s.fetchAndVerifySpot(spotId, requestUserId, domain.ONLINE_SPOT)

	if err != nil {
		log.Println("We found an error while verifying the spot\n\t\t", err.Error())
		return err
	}

	var dateIdToStop string
	for _, date := range onlineSpot.DatesInfo {
		if date.HostInfo.HostId == requestUserId {
			dateIdToStop = date.DateId
		}
	}

	if len(dateIdToStop) == 0 {
		return ports.ErrUserIsNotHostingAnDate
	}

	if err = s.repository.StopDateOnSpot(spotId, dateIdToStop); err != nil {
		log.Printf("We foinf an error while stoping the date %s on spot %s \n", spotId, dateIdToStop)
		return err
	}

	s.notifer.SchedulePatternActivity(ports.OnlineStop, domain.Notification{
		SpotId: spotId,
		DateId: dateIdToStop,
		UserId: requestUserId,
	})

	return nil
}

/*
	The user should not be hosting more than two events at the same time

	Procedure:
		1. Verify that the user is the owner of the spot. --- OK
			If no -> Return an error

		2. Check that the spot does not have a current online event.
			If so -> Return and error

		3. Create Date - Spot association with the RequestUserId as the IS_HOSTING relation

*/

func (s *Service) Start(spotId string, requestUserId string, durationApproximated int64, maximunCapacity int64) (domain.OnlineSpot, error) {
	log.Printf("Service Start: spotId %s, requestUserId: %s \n", spotId, requestUserId)

	// 1. Verify that the user is the owner of the spot
	onlineSpot, err := s.fetchSpotsByStatus(spotId, requestUserId, domain.FlagOnline|domain.FlagPaused)

	if err != nil {
		log.Println("We found an error while verifying the spot\n\t\t", err.Error())
		return domain.OnlineSpot{}, err
	}

	// 2. Check that the spot does not have a current online event.
	if len(onlineSpot.DatesInfo) > 0 {
		log.Printf("DateInfo is not empty, so the event is already online: spotDate %+v \n", onlineSpot.DatesInfo)
		return domain.OnlineSpot{}, ports.ErrSpotIsAlreadyOnline
	}

	creationTime := time.Now().UTC()

	dateInfo := domain.SpotDate{
		DateId:                        s.generateUUID(),
		DurationApproximatedInSeconds: durationApproximated,
		StartTime:                     creationTime.Format(time.TimeOnly),
		State: domain.SpotState{
			Confirmed: true,
		},
		MaximunCapacty: maximunCapacity,
		Date:           creationTime.Format(time.DateOnly),
		HostInfo: domain.HostInfo{
			HostId: requestUserId,
		},
	}

	onlineSpot.DatesInfo = []domain.SpotDate{
		dateInfo,
	}

	// 3. Create Date - Spot association with the RequestUserId as the IS_HOSTING relation
	err = s.repository.AssociateDateWithSpot(onlineSpot)

	if err != nil {
		log.Println("An error happen while associating the date with the spot, error: ", err.Error())
		// log.Printf("OnlineSpot: %+v \n")
		return domain.OnlineSpot{}, err
	}

	s.notifer.SchedulePatternActivity(ports.OnlineStart, domain.Notification{
		DateId:           dateInfo.DateId,
		SpotId:           spotId,
		UserId:           requestUserId,
		Aditionalpayload: onlineSpot,
	})

	return onlineSpot, nil

}

func (s *Service) fetchAndVerifySpot(spotId string, requestUserId string, spotStatus domain.SpotStatus) (domain.OnlineSpot, error) {
	log.Printf("Service fetchAndVerifySpot: spotId %s, requestUserId: %s \n", spotId, requestUserId)

	onlineSpot, err := s.repository.FetchSpotWithStatus(spotId, spotStatus)

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
