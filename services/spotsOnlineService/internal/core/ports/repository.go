package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/core/domain"
)

var (
	ErrFetchingData     = errors.New("we got an error while fetching the data from the repository")
	ErrAssociatingDate  = errors.New("We got an error while associating the date")
	ErrUpdatingAtStatus = errors.New("We got an error while changing the status of the date")
)

type Repository interface {
	FetchOnlineSpot(spotId string) (domain.OnlineSpot, error)
	FetchSpotWithStatus(spotId string, status domain.SpotStatus) (domain.OnlineSpot, error)
	FetchSpots(spotId string) (domain.OnlineSpot, error)
	AssociateDateWithSpot(domain.OnlineSpot) error
	StopDateOnSpot(spotId string, dateId string) error
	ResumeDateOnSpo(spotId string, dateId string) error
	FinalizeDateOnSpot(spotId string, dateId string) error
}
