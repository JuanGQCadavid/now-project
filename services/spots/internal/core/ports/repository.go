package ports

import "github.com/JuanGQCadavid/now-project/services/spots/internal/core/domain"

type SpotRepository interface {
	Get(id string) (domain.Spot, error)
	CreateOnline(spot domain.Spot) error
}
