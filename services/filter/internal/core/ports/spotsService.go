package ports

import (
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"
)

type SpotService interface {
	GetSpotsCardsInfo(spots []string, format OutputFormat) ([]domain.Spot, error)
}
