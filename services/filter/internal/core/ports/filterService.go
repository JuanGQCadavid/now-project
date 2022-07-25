package ports

import (
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain/session"
)

type FilterService interface {
	FilterByProximity(centralPointLat float64, centralPointLng float64, radious float64, session session.SearchSessionData) domain.Locations
}
