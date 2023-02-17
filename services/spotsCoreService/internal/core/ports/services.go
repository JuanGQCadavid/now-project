package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"
)

type OutputFormat string

const (
	SMALL_FORMAT OutputFormat = "SMALL"
	FULL_FORMAT  OutputFormat = "FULL"
)

var (
	ErrSpotUserNotOwnerWhenUpdatingSpot = errors.New("err the user that request the update is not the owner of the spot")
	ErrSpotToUpdateIsTheSameAsTheDb     = errors.New("err the spot to save is the same as the one in the db")
)

type SpotService interface {
	// Fetch
	Get(spotId string, format OutputFormat) (domain.Spot, error)
	GetSpots(spotIds []string, format OutputFormat) (domain.MultipleSpots, error)
	CreateSpot(spot domain.Spot) (domain.Spot, error)

	UpdateSpotEvent(spotId string, ownerId string, spotEvent *domain.Event) error
	UpdateSpotPlace(spotId string, ownerId string, spotEvent *domain.Place) error
	UpdateSpotTopic(spotId string, ownerId string, spotEvent *domain.Topic) error

	// Missing Specification
	FinalizeSpot(spotId string) error
}
