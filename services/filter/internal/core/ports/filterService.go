package ports

import "github.com/JuanGQCadavid/now-project/services/filter/internal/core/models"

type FilterService interface {
	FilterByProximity(centralPointLat float32, centralPointLng float32, radious float32) models.Locations
}
