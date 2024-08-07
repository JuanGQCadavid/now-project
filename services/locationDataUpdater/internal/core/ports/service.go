package ports

import "github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/core/domain"

type Service interface {
	OnDateCreation(domain.DatesLocation) error
	OnDateRemoved(string) error
	OnDateStateChanged(string, domain.DateState) error
	OnDateTypeChanged(string, domain.DateType) error
}
