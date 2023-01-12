package ports

import "github.com/JuanGQCadavid/now-project/services/spotsCore/internal/core/domain"

type OutputFormat string

const (
	SMALL_FORMAT OutputFormat = "SMALL"
	FULL_FORMAT  OutputFormat = "FULL"
)

type SpotService interface {
	// Fetch
	Get(spotId string, format OutputFormat) (domain.Spot, error)
	GetSpots(spotIds []string, format OutputFormat) (domain.MultipleSpots, error)
	CreateSpot(spot domain.Spot) (domain.Spot, error)

	// Missing Specification
	UpdateSpot() error
	FinalizeSpot(spotId string) error
}
