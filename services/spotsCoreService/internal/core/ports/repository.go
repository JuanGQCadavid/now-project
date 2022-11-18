package ports

import "github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"

type SpotRepository interface {
	Get(id string, format OutputFormat) (domain.Spot, error)
	CreateOnline(spot domain.Spot) error
	GetSpotByUserId(personId string) (domain.Spot, error)
	EndSpot(spotId string) error
	GetSpots(spotIds []string, format OutputFormat) (domain.MultipleSpots, error)
	CreateSpotTags(spotId string, principalTag domain.Optional, secondaryTags []string) error
}
