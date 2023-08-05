package ports

import "github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/core/domain"

type LocationRepository interface {
	CrateLocation(domain.DatesLocation) error
	RemoveLocation(string) error
	UpdateLocationStatus(string, domain.DateState) error
}
