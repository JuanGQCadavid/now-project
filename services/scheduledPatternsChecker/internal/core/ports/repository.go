package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/domain"
)

var (
	ErrFetchingData        = errors.New("We face a problem on the repository")
	ErrDeletingSpotsFormSp = errors.New("We face a problem while removinf dates from Sp")
)

type Repository interface {
	FetchActiveSchedulePatterns() ([]domain.Spot, error)
	UpdateSpotsByBatch(spots []domain.Spot, batchSize int) map[*domain.Spot]error
	ConditionalDatesCreation(spot domain.Spot) error
	DeleteScheduleDatesFromSchedulePattern(schedulePatternIds []string) ([]string, error)
}
