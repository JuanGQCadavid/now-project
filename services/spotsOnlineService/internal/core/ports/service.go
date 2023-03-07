package ports

import "errors"

var (
	ErrOnRepository        = errors.New("we found a problem while requesting to the repository")
	ErrUserIsNotTheOwner   = errors.New("the user is not the owner of the spot")
	ErrSpotNotFound        = errors.New("the spot does not exist")
	ErrSpotIsAlreadyOnline = errors.New("the spot is already online")
)

type SpotOnlineService interface {
	Start(spotId string, requestUserId string)
	Stop(spotId string, requestUserId string)
	Resume(spotId string, requestUserId string)
	Finalize(spotId string, requestUserId string)
}
