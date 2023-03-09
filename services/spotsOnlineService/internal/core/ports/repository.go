package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/core/domain"
)

var (
	ErrFetchingData    = errors.New("we got an error while fetching the data from the repository")
	ErrAssociatingDate = errors.New("We got an error while associating the date")
)

type Repository interface {
	FetchOnlineSpot(spotId string) (domain.OnlineSpot, error)
	AssociateDateWithSpot(domain.OnlineSpot) error
}
