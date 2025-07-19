package ports

import (
	"context"
	"errors"

	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"
)

var (
	ErrSpotUpdatedFail   = errors.New("err the spot was update but it is not the same as the one that we send to update")
	ErrCallingRepository = errors.New("err repository crash")
)

type SpotRepository interface {
	Get(id string, format OutputFormat) (domain.Spot, error)
	GetSpotByUserId(personId string) (domain.Spot, error)
	GetSpotsByDatesId(datesIds []string, format OutputFormat) (domain.MultipleSpots, error)

	CreateSpot(spot domain.Spot) error
	CreateSpotTags(spotId string, principalTag domain.Optional, secondaryTags []string) error

	UpdateSpotEvent(spotEvent domain.Event, spotId string) error
	UpdateSpotPlace(spotEvent domain.Place, spotId string) error
	UpdateSpotTopic(spotEvent domain.Topic, spotId string) error

	DeleteSpot(spotId string) error

	GetUserEventRole(ctx context.Context, userId, eventId string) (*domain.Access, error)
	GetDateAttendantsWithRole(ctx context.Context, eventId, dateId string) ([]*domain.Access, error)
}
