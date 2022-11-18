package ports

import "github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"

type SpotActivityTopic interface {
	AppendSpot(spot domain.Spot) error
	RemoveSpot(spotId string) error
}
