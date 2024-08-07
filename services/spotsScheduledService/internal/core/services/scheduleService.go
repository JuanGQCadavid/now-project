package services

import (
	"time"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/ports"
	"github.com/google/uuid"
)

type ScheduledService struct {
	repository ports.Repository
	notifier   ports.Notify
}

func NewScheduledService(repository ports.Repository, notifier ports.Notify) *ScheduledService {
	return &ScheduledService{
		repository: repository,
		notifier:   notifier,
	}
}

func (service *ScheduledService) GetSchedules(spotId string, userRequestId string, flags domain.ScheduleStateFlags) (*domain.ScheduledSpot, error) {
	logs.Info.Printf("GetSchedules: spotId: %s, userRequestId: %s, flags: %08b \n", spotId, userRequestId, flags)
	spot, err := service.repository.GetScheduleSpot(spotId, flags)

	if err != nil {
		logs.Error.Println("We found an error while fetching the spot \n\t\t", err.Error())
		return nil, ports.ErrOnRepository
	}

	if len(spot.SpotInfo.SpotId) == 0 {
		logs.Error.Println("The spot does not exist")
		return nil, ports.ErrSpotNotFound
	}

	if spot.SpotInfo.OwnerId != userRequestId {
		logs.Warning.Println("The owner id is differente than the spot owner, so changing the flags to only ACTIVE")

		filteredSchedulePatterns := []domain.SchedulePattern{}

		for _, schedulePattern := range spot.Patterns {
			if schedulePattern.State.Status == domain.ACTIVATE {
				filteredSchedulePatterns = append(filteredSchedulePatterns, schedulePattern)
			}
		}

		spot.Patterns = filteredSchedulePatterns
	}

	return spot, nil
}

func (service *ScheduledService) GetDates(spotId string, userRequestId string) ([]domain.Date, error) {
	logs.Info.Printf("GetDates: spotId: %s, userRequestId: %s\n", spotId, userRequestId)
	dates, err := service.repository.GetDatesFromSpot(spotId)

	if err != nil {
		logs.Error.Println("We found an error while fetching the spot \n\t\t", err.Error())
		return nil, ports.ErrOnRepository
	}

	return dates, nil
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
func (service *ScheduledService) AppendSchedule(spotId string, userRequestId string, schedulesPattern []domain.SchedulePattern) (*domain.ScheduledSpot, *[]domain.TimeConflict, error) {
	logs.Info.Printf("AppendSchedule: userRequestId: %s \n", userRequestId)

	//  0. Verify that the incomming schedules does not have tinme conflict
	service.createIdsForSchedulePatterns(schedulesPattern)
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
	timeConflicts, err = service.verifyTimeConflics(spot.Patterns, schedulesPattern, false)

	if err != nil {
		return nil, nil, ports.ErrValidatingPatterns
	}

	if timeConflicts != nil && len(*timeConflicts) > 0 {
		return nil, timeConflicts, ports.ErrTimeConflictWithActualSchedulePattenrs
	}

	// Update status for each schedule.

	for i, scheduleSpot := range schedulesPattern {
		scheduleSpot.State = domain.State{
			Status: domain.ACTIVATE,
			Since:  time.Now().Unix(),
		}
		schedulesPattern[i] = scheduleSpot
	}

	// 4.Associate new schedule
	err = service.repository.AssociateSpotWithSchedulePatterns(spotId, userRequestId, schedulesPattern)

	if err != nil {
		return nil, nil, err
	}

	spot.Patterns = schedulesPattern

	service.notifier.SchedulePatternActivity(ports.SchedulePatternAppended, domain.Notification{
		SpotId:           spotId,
		UserId:           userRequestId,
		Aditionalpayload: spot,
	})

	return spot, nil, nil
}

func (service *ScheduledService) ResumeSchedule(spotId string, scheduleId string, userRequestId string) error {
	logs.Info.Printf("ResumeSchedule: spotId: %s, userRequestId: %s \n", spotId, userRequestId)
	spot, err := service.getAndVerifyScheduleSpot(spotId, userRequestId, domain.ActivateFlag|domain.FreezeFlag)

	if err != nil {
		return err
	}

	index := -1
	for ii, schedulePattern := range spot.Patterns {
		if schedulePattern.Id == scheduleId {
			index = ii
			if schedulePattern.State.Status == domain.ACTIVATE {
				return ports.ErrScheduleIsAlreadyActivated
			} else if schedulePattern.State.Status == domain.CONCLUDE {
				return ports.ErrScheduleIsConclude
			}
		}
	}

	if index == -1 {
		return ports.ErrScheduleIsDoesNotExist
	}

	newStatus := domain.State{
		Status: domain.ACTIVATE,
		Since:  time.Now().Unix(),
	}

	err = service.repository.UpdateScheculeStatus(spotId, scheduleId, newStatus)

	if err != nil {
		logs.Error.Println("We found an error whule Updating the schedule status, error: ", err.Error())
	}

	spot.Patterns = []domain.SchedulePattern{
		spot.Patterns[index],
	}

	service.notifier.SchedulePatternActivity(ports.SchedulePatternResumed, domain.Notification{
		SpotId:           spotId,
		ScheduleId:       scheduleId,
		UserId:           userRequestId,
		Aditionalpayload: spot,
	})

	return err
}

func (service *ScheduledService) FreezeSchedule(spotId string, scheduleId string, userRequestId string) error {
	logs.Info.Printf("FreezeSchedule: spotId: %s, userRequestId: %s \n", spotId, userRequestId)
	spot, err := service.getAndVerifyScheduleSpot(spotId, userRequestId, domain.ActivateFlag|domain.FreezeFlag)

	if err != nil {
		return err
	}

	index := -1
	for ii, schedulePattern := range spot.Patterns {
		if schedulePattern.Id == scheduleId {
			index = ii
			if schedulePattern.State.Status == domain.FREEZE {
				return ports.ErrScheduleIsAlreadyFreezed
			} else if schedulePattern.State.Status == domain.CONCLUDE {
				return ports.ErrScheduleIsConclude
			}
		}
	}

	if index == -1 {
		return ports.ErrScheduleIsDoesNotExist
	}

	newStatus := domain.State{
		Status: domain.FREEZE,
		Since:  time.Now().Unix(),
	}

	err = service.repository.UpdateScheculeStatus(spotId, scheduleId, newStatus)

	if err != nil {
		logs.Error.Println("We found an error whule Updating the schedule status, error: ", err.Error())
	}

	spot.Patterns = []domain.SchedulePattern{
		spot.Patterns[index],
	}

	service.notifier.SchedulePatternActivity(ports.SchedulePatternFreezed, domain.Notification{
		SpotId:           spotId,
		ScheduleId:       scheduleId,
		UserId:           userRequestId,
		Aditionalpayload: spot,
	})

	return err
}

func (service *ScheduledService) ConcludeSchedule(spotId string, scheduleId string, userRequestId string) error {
	logs.Info.Printf("ConcludeSchedule: spotId: %s, userRequestId: %s \n", spotId, userRequestId)
	spot, err := service.getAndVerifyScheduleSpot(spotId, userRequestId, domain.ActivateFlag|domain.FreezeFlag)

	if err != nil {
		return err
	}

	index := -1
	for ii, schedulePattern := range spot.Patterns {
		if schedulePattern.Id == scheduleId {
			index = ii
		}
	}

	if index == -1 {
		return ports.ErrScheduleIsDoesNotExist
	}

	newStatus := domain.State{
		Status: domain.CONCLUDE,
		Since:  time.Now().Unix(),
	}

	err = service.repository.UpdateScheculeStatus(spotId, scheduleId, newStatus)

	if err != nil {
		logs.Error.Println("We found an error whule Updating the schedule status, error: ", err.Error())
	}

	service.notifier.SchedulePatternActivity(ports.SchedulePatternConcluded, domain.Notification{
		SpotId:           spotId,
		ScheduleId:       scheduleId,
		UserId:           userRequestId,
		Aditionalpayload: newStatus,
	})

	return err
}

func (service *ScheduledService) getAndVerifyScheduleSpot(spotId string, userRequestId string, flags domain.ScheduleStateFlags) (*domain.ScheduledSpot, error) {
	spot, err := service.repository.GetScheduleSpot(spotId, flags)

	if err != nil {
		logs.Error.Println("We found an error while fetching the spot \n\t\t", err.Error())
		return nil, ports.ErrOnRepository
	}

	if len(spot.SpotInfo.SpotId) == 0 {
		logs.Error.Println("The spot does not exist")
		return nil, ports.ErrSpotNotFound
	}

	if spot.SpotInfo.OwnerId != userRequestId {
		logs.Error.Println("The owner id is differente than the spot owner")
		return nil, ports.ErrUserIsNotTheOwner
	}

	return spot, nil
}

func (service *ScheduledService) verifyTimeConflics(firstPatterns []domain.SchedulePattern, secondPatterns []domain.SchedulePattern, selfVerification bool) (*[]domain.TimeConflict, error) {

	logs.Info.Printf("verifyTimeConflics: firstPattern %+v \n", firstPatterns)

	result := make([]domain.TimeConflict, len(firstPatterns))
	var conflictFound bool

	for index, pattern := range firstPatterns {

		for ii, subPattern := range secondPatterns {

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
