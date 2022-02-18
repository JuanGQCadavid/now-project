package ports

import (
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"
)

type FilterService interface {
	FilterByProximity(centralPointLat float64, centralPointLng float64, radious float64) domain.Locations
}
