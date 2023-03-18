package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/core/domain"
)

var (
	ErrOnRepository               = errors.New("we found a problem while requesting to the repository")
	ErrUserIsNotTheOwner          = errors.New("the user is not the owner of the spot")
	ErrSpotNotFound               = errors.New("the spot does not exist")
	ErrSpotIsAlreadyOnline        = errors.New("the spot is already online or there is a paused spot")
	ErrUserIsNotHostingAnDate     = errors.New("The user is not hosting an event in the spot")
	ErrUserDoesNotHaveStoppedDate = errors.New("The user does not have a stopped date")
)

type SpotOnlineService interface {
	Start(spotId string, requestUserId string, durationApproximated int64, maximunCapacity int64) (domain.OnlineSpot, error)
	Stop(spotId string, requestUserId string) error
	Resume(spotId string, requestUserId string) error
	Finalize(spotId string, requestUserId string) error
}
