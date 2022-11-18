package ports

import "github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"

type OutputFormat string

const (
	SMALL_FORMAT OutputFormat = "SMALL"
	FULL_FORMAT  OutputFormat = "FULL"
)

type SpotService interface {
	Get(spotId string, format OutputFormat) (domain.Spot, error)
	GoOnline(spot domain.Spot) (domain.Spot, error)
	EndSpot(spotId string) error
	GetSpots(spotIds []string, format OutputFormat) (domain.MultipleSpots, error)
	//scheduledSpot()
}
