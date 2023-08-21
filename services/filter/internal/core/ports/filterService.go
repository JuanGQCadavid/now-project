package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain/session"
)

type OutputFormat string

const (
	SMALL_FORMAT OutputFormat = "SMALL"
	FULL_FORMAT  OutputFormat = "FULL"
)

var (
	ErrSpotServiceFail = errors.New("There were an error while calling Spot service")
	ErrRepositoryFail  = errors.New("There were an error while fetching the data in the repository")
)

type FilterService interface {
	FilterByProximity(centralPointLat float64, centralPointLng float64, radious float64, session session.SearchSessionData, format OutputFormat) (domain.Locations, error)
}
