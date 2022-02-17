package ports

import (
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"
)

type LocationRepository interface {
	FetchSpotsIdsByArea(pointA domain.LatLng, pointB domain.LatLng) (domain.Locations, error)
}
