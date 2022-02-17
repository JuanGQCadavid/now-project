package ports

import (
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"
)

type FilterService interface {
	FilterByProximity(centralPointLat float32, centralPointLng float32, radious float32) domain.Locations
}
