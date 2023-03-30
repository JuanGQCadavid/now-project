package ports

import "github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/domain"

type Repository interface {
	GetScheduleSpot(spotId string, flags domain.ScheduleStateFlags) (*domain.ScheduledSpot, error)
	AssociateSpotWithSchedulePatterns(spotId string, hostId string, schedulesPattern *[]domain.SchedulePattern) error
}
