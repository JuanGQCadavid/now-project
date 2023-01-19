package ports

import "github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"

type SpotRepository interface {
	Get(id string, format OutputFormat) (domain.Spot, error)
	GetSpotByUserId(personId string) (domain.Spot, error)
	GetSpots(spotIds []string, format OutputFormat) (domain.MultipleSpots, error)

	CreateSpot(spot domain.Spot) error
	CreateSpotTags(spotId string, principalTag domain.Optional, secondaryTags []string) error

	EndSpot(spotId string) error
}
