package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/domain"
)

var (
	ErrOnRepository         = errors.New("err on repository")
	ErrInvalidCors          = errors.New("The cores number is 0 or negative")
	ErrProcessingDates      = errors.New("We got an error while pocessing dates")
	ErrSendingConfirmation  = errors.New("Confirmation service fail to send")
	ErrServiceParcialOutage = errors.New("We face a partial outage while creating dates")
)

type Service interface {
	DeleteScheduleDatesFromSchedulePattern(schedulePatternIds []string) error
	CreateScheduledDatesFromSchedulePattern(spots []domain.Spot, timeWindow int64) ([]domain.Spot, map[error][]domain.Spot)
	GenerateDatesFromRepository(timeWindow int64) ([]domain.Spot, error)
}
