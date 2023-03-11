package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/core/domain"
)

var (
	ErrOnRepository        = errors.New("we found a problem while requesting to the repository")
	ErrUserIsNotTheOwner   = errors.New("the user is not the owner of the spot")
	ErrSpotNotFound        = errors.New("the spot does not exist")
	ErrSpotIsAlreadyOnline = errors.New("the spot is already online")
)

type SpotOnlineService interface {
	Start(spotId string, requestUserId string, durationApproximated int64, maximunCapacity int) (domain.OnlineSpot, error)
	Stop(spotId string, requestUserId string)
	Resume(spotId string, requestUserId string)
	Finalize(spotId string, requestUserId string)
}
