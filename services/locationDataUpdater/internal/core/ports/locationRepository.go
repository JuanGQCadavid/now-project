package ports

import "github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/core/domain"

type LocationRepository interface {
	CrateLocation(domain.Date) error
	RemoveLocation(string) error
	UpdateLocationStatus(string, domain.DateStatus) error
}
