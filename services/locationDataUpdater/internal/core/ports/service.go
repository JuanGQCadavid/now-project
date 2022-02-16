package ports

import "github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/core/domain"

type Service interface {
	OnSpotCreation(domain.Spot) error
	OnSpotDeletion(string) error
}
