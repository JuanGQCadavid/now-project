package ports

import "github.com/JuanGQCadavid/now-project/services/spots/internal/core/domain"

type SpotService interface {
	Get(spotId string) (domain.Spot, error)
	GoOnline(spot domain.Spot) (domain.Spot, error)
	EndSpot(spotId string) error
	//scheduledSpot()
}
