package ports

import (
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/models"
)

type LocationRepository interface {
	FetchSpotsIdsByArea(pointA models.LatLng, pointB models.LatLng) (models.Locations, error)
}
