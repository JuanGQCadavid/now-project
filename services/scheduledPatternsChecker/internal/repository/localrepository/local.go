package localrepository

import (
	"math/rand"

	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/domain"
)

type LocalRepository struct {
	upTo    int16
	maxDeep int16
}

func NewLocalRepository(upTo int16, maxDeep int16) *LocalRepository {
	return &LocalRepository{
		upTo:    upTo,
		maxDeep: maxDeep,
	}
}

func (r *LocalRepository) FetchActiveSchedulePatterns() ([]domain.Spot, error) {

	return r.generateSpots(r.upTo, r.maxDeep), nil
}

func (r *LocalRepository) generateSpots(upTo int16, maxDeep int16) []domain.Spot {
	result := make([]domain.Spot, upTo)

	for i, _ := range result {
		patternsCount := rand.Intn(int(maxDeep))
		schedulePatterns := make([]domain.SchedulePattern, patternsCount)
		for i := 0; i < patternsCount; i++ {
			schedulePatterns[i] = domain.SchedulePattern{
				Id:        string(i),
				HostId:    "JUAN123",
				Day:       domain.Saturday,
				FromDate:  "2007-03-01",
				ToDate:    "2007-07-01",
				StartTime: "13:00:00",
				EndTime:   "16:00:00",
			}
		}

		result[i] = domain.Spot{
			SpotId:           string(i),
			SchedulePatterns: schedulePatterns,
		}

	}
	return result
}
