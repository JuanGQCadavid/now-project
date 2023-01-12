package ports

import "github.com/JuanGQCadavid/now-project/services/spotsCore/internal/core/domain"

type SpotActivityTopic interface {
	NotifySpotCreated(spot domain.Spot) error
	NotifySpotStopped(spotId string) error
}
