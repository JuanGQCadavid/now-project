package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/domain"
)

var (
	ErrOnRepository                              = errors.New("we found a problem while requesting to the repository")
	ErrUserIsNotTheOwner                         = errors.New("the user is not the owner of the spot")
	ErrSpotNotFound                              = errors.New("the spot does not exist")
	ErrTimeConflictWithRequestedSchedulePattenrs = errors.New("There is one or more time conflicts with the requrested schedule patterns")
	ErrTimeConflictWithActualSchedulePattenrs    = errors.New("There is one or more time conflicts with the requrested schedule patterns against the actual ones")
	ErrValidatingPatterns                        = errors.New("We face an error while validating the patterns, could be due to the format")
	ErrScheduleIsAlreadyFreezed                  = errors.New("The schedule id is already freezed")
	ErrScheduleIsConclude                        = errors.New("The schedule id is concluded")
	ErrScheduleIsAlreadyActivated                = errors.New("The schedule id is already activated")
	ErrScheduleIsDoesNotExist                    = errors.New("The schedule id does not exist or is concluded")
)

type Service interface {
	GetSchedules(spotId string, userRequestId string, flags domain.ScheduleStateFlags) (*domain.ScheduledSpot, error)
	AppendSchedule(spotId string, userRequestId string, schedulesPattern []domain.SchedulePattern) (*domain.ScheduledSpot, *[]domain.TimeConflict, error)
	ResumeSchedule(spotId string, scheduleId string, userRequestId string) error
	FreezeSchedule(spotId string, scheduleId string, userRequestId string) error
	ConcludeSchedule(spotId string, scheduleId string, userRequestId string) error
	GetDates(spotId string, userRequestId string) ([]domain.Date, error)
}
