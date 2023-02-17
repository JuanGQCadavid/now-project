package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"
)

var (
	ErrSpotUpdatedFail = errors.New("The spot was update but it is not the same as the one that we send to update")
)

type SpotRepository interface {
	Get(id string, format OutputFormat) (domain.Spot, error)
	GetSpotByUserId(personId string) (domain.Spot, error)
	GetSpots(spotIds []string, format OutputFormat) (domain.MultipleSpots, error)

	CreateSpot(spot domain.Spot) error
	CreateSpotTags(spotId string, principalTag domain.Optional, secondaryTags []string) error

	UpdateSpotEvent(spotEvent domain.Event, spotId string) error
	UpdateSpotPlace(spotEvent domain.Place, spotId string) error
	UpdateSpotTopic(spotEvent domain.Topic, spotId string) error

	EndSpot(spotId string) error
}
