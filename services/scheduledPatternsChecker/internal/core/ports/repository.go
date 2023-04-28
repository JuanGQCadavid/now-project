package ports

import "github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/domain"

type Repository interface {
	FetchActiveSchedulePatterns() ([]domain.Spot, error)
}
