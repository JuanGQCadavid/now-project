package ports

import (
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/models"
)

type LocationRepository interface {
	FetchSpotsIdsByArea(city string, pointA models.LatLng, pointB models.LatLng) (models.Locations, error)
}
