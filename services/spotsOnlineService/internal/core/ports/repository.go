package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/core/domain"
)

var (
	ErrFetchingData = errors.New("we got and error while fetching the data from the repository")
)

type Repository interface {
	FetchOnlineSpot(spotId string) (domain.OnlineSpot, error)
	AssociateDateWithSpot(domain.OnlineSpot) (string, error)
}
