package ports

import (
	"context"
	"errors"

	authDomain "github.com/JuanGQCadavid/now-project/services/authService/core/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"
)

type OutputFormat string

const (
	SMALL_FORMAT OutputFormat = "SMALL"
	FULL_FORMAT  OutputFormat = "FULL"
)

var (
	ErrSpotUserNotOwnerWhenUpdatingSpot       = errors.New("err the user is not the owner of the spot")
	ErrSpotToUpdateIsTheSameAsTheDb           = errors.New("err the spot to save is the same as the one in the db")
	ErrSpotNotFounded                         = errors.New("err the spot is not founded in the repository")
	ErrUserIDMissing                          = errors.New("err user ID is empty")
	ErrDateDoesNotExist                 error = errors.New("err date does not exsit")
)

type SpotService interface {
	// Fetch
	Get(spotId string, format OutputFormat) (domain.Spot, error)
	GetSpotsByDatesIds(spotIds []string, format OutputFormat) (domain.MultipleSpots, error)
	CreateSpot(spot domain.Spot, userDetails authDomain.UserDetails) (domain.Spot, error)

	UpdateSpotEvent(spotId string, ownerId string, spotEvent *domain.Event) error
	UpdateSpotPlace(spotId string, ownerId string, spotEvent *domain.Place) error
	UpdateSpotTopic(spotId string, ownerId string, spotEvent *domain.Topic) error

	// Missing Specification
	DeleteSpot(spotId string, requestUserId string) error

	// Access
	GetAccess(ctx context.Context, userId, eventId, dateId string) (*domain.Access, error)
}
