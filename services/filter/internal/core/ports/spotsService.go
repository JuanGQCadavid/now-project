package ports

import "github.com/JuanGQCadavid/now-project/services/filter/internal/core/models"

type SpotService interface {
	GetSpotsCardsInfo(spots []string) ([]models.Spot, error)
}
