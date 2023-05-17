package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/domain"
)

var (
	ErrFetchingData = errors.New("We face a problem on the repository")
)

type Repository interface {
	FetchActiveSchedulePatterns() ([]domain.Spot, error)
	UpdateSpotsByBatch(spots []domain.Spot, batchSize int) map[*domain.Spot]error
}
