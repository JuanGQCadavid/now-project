package ports

import "github.com/JuanGQCadavid/now-project/services/spots/internal/core/domain"

type SpotRepository interface {
	Get(id string) (domain.Spot, error)
	CreateOnline(spot domain.Spot) error
	GetSpotByUserId(personId string) (domain.Spot, error)
	EndSpot(spotId string) error
	GetSpots(spotIds []string) (domain.MultipleSpots, error)
}
