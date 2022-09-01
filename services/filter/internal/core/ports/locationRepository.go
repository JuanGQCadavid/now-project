package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"
)

var (
	ErrDBEnvCredentialsMissing = errors.New("Missing db env data")
	ErrUnableToCreateDBSession = errors.New("Unable to create db session")
	ErrQueringData             = errors.New("A error occoure while quering the data")
)

type LocationRepository interface {
	FetchSpotsIdsByArea(pointA domain.LatLng, pointB domain.LatLng) (domain.Locations, error)
	FetchSpotsIdsByAreaExcludingSpots(pointA domain.LatLng, pointB domain.LatLng, spotsIdsToExclude []string) (domain.Locations, error)
}
