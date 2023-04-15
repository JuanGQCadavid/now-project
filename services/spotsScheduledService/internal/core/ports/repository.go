package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/domain"
)

var (
	ErrNotRecordsToProcess = errors.New("There is not records to process from repository")
	ErrSpotStatusUndefined = errors.New("The spot status is unknown")
)

type Repository interface {
	GetScheduleSpot(spotId string, flags domain.ScheduleStateFlags) (*domain.ScheduledSpot, error)
	AssociateSpotWithSchedulePatterns(spotId string, hostId string, schedulesPattern []domain.SchedulePattern) error
	UpdateScheculeStatus(spotId string, scheduleId string, status domain.State) error
}
