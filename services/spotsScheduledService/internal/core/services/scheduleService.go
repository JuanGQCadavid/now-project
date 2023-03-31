package services

import (
	"log"

	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/logs"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/ports"
	"github.com/google/uuid"
)

type ScheduledService struct {
	repository ports.Repository
}

func NewScheduledService(repository ports.Repository) *ScheduledService {
	return &ScheduledService{
		repository: repository,
	}
}

func (service *ScheduledService) GetSchedules(spotId string, userRequestId string) (*domain.ScheduledSpot, error) {

	return nil, nil
}

/*
This service will append a schedule patter to a spot.
Constrains:

 1. There should not be a freeze nor active schedule pattern that share portion of the time of the new schedule pattern

Steps:

 0. Verify that the incomming schedules does not have tinme conflict
    1.Get the spot with all its schedules patterns
    2.check that the user is the owner of the spot
    3.Check time conflicts
    4.Associate new schedule
    5.Return the new schedule pattern with its UUID
*/
func (service *ScheduledService) AppendSchedule(spotId string, userRequestId string, schedulesPattern *[]domain.SchedulePattern) (*domain.ScheduledSpot, *[]domain.TimeConflict, error) {
	logs.Info.Printf("AppendSchedule: userRequestId: %s \n", userRequestId)

	//  0. Verify that the incomming schedules does not have tinme conflict
	service.createIdsForSchedulePatterns(*schedulesPattern)
	timeConflicts, err := service.verifyTimeConflics(schedulesPattern, schedulesPattern, true)

	if err != nil {
		return nil, nil, ports.ErrValidatingPatterns
	}

	if timeConflicts != nil && len(*timeConflicts) > 0 {
		return nil, timeConflicts, ports.ErrTimeConflictWithRequestedSchedulePattenrs
	}

	// 1.Get the spot with all its schedules patterns
	// 2.check that the user is the owner of the spot
	spot, err := service.getAndVerifyScheduleSpot(spotId, userRequestId, domain.ActivateFlag|domain.FreezeFlag)

	if err != nil {
		return nil, nil, err
	}
	// 3.Check time conflicts
	timeConflicts, err = service.verifyTimeConflics(&spot.Patterns, schedulesPattern, false)

	if err != nil {
		return nil, nil, ports.ErrValidatingPatterns
	}

	if timeConflicts != nil && len(*timeConflicts) > 0 {
		return nil, timeConflicts, ports.ErrTimeConflictWithActualSchedulePattenrs
	}

	// 4.Associate new schedule
	err = service.repository.AssociateSpotWithSchedulePatterns(spotId, userRequestId, schedulesPattern)

	if err != nil {
		return nil, nil, err
	}

	return spot, nil, nil
}
func (service *ScheduledService) ResumeSchedule(spotId string, userRequestId string) error {
	return nil
}
func (service *ScheduledService) FreezeSchedule(spotId string, userRequestId string) error {
	return nil
}
func (service *ScheduledService) ConcludeSchedule(spotId string, userRequestId string) error {
	return nil
}

func (service *ScheduledService) getAndVerifyScheduleSpot(spotId string, userRequestId string, flags domain.ScheduleStateFlags) (*domain.ScheduledSpot, error) {
	spot, err := service.repository.GetScheduleSpot(spotId, flags)

	if err != nil {
		log.Println("We found an error while fetching the spot owner\n\t\t", err.Error())
		return nil, err
	}

	if len(spot.SpotInfo.SpotId) == 0 {
		log.Println("The spot does not exist")
		return nil, ports.ErrSpotNotFound
	}

	if spot.SpotInfo.OwnerId != userRequestId {
		log.Println("The owner id is differente than the spot owner")
		return nil, ports.ErrUserIsNotTheOwner
	}

	return spot, nil
}

func (service *ScheduledService) verifyTimeConflics(firstPatterns *[]domain.SchedulePattern, secondPatterns *[]domain.SchedulePattern, selfVerification bool) (*[]domain.TimeConflict, error) {

	logs.Info.Printf("verifyTimeConflics: firstPattern %+v \n", firstPatterns)

	result := make([]domain.TimeConflict, len(*firstPatterns))
	var conflictFound bool

	for index, pattern := range *firstPatterns {

		for ii, subPattern := range *secondPatterns {

			if selfVerification && ii <= index {
				continue
			}
			isOverlaped, err := pattern.Overlaps(subPattern)

			if err != nil {
				logs.Error.Println("We found an error while calling Overlaps function")
				return nil, err
			}
			if isOverlaped {
				conflictFound = true
				if len(result[index].SchedulePattern.Id) == 0 {
					result[index] = domain.TimeConflict{
						SchedulePattern: pattern,
					}
				}
				result[index].ConflictWith = append(result[index].ConflictWith, subPattern)
			}
		}

	}

	if conflictFound {
		return &result, nil
	}

	return nil, nil
}

func (service *ScheduledService) createIdsForSchedulePatterns(schedulePattern []domain.SchedulePattern) {
	logs.Info.Println("createIdsForSchedulePatterns")
	for i, sp := range schedulePattern {
		sp.Id = uuid.NewString()
		schedulePattern[i] = sp
	}
	logs.Info.Println(schedulePattern)

}
