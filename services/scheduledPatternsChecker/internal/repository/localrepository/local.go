package localrepository

import (
	"fmt"
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

func (r *LocalRepository) UpdateSpotsByBatch(spots []domain.Spot, batchSize int) map[*domain.Spot]error {
	return nil
}

func (r *LocalRepository) generateSpots(upTo int16, maxDeep int16) []domain.Spot {
	result := make([]domain.Spot, upTo)

	for i, _ := range result {
		patternsCount := rand.Intn(int(maxDeep))
		schedulePatterns := make([]domain.SchedulePattern, patternsCount)
		for i := 0; i < patternsCount; i++ {

			if i%2 == 0 {
				schedulePatterns[i] = domain.SchedulePattern{
					Id:        fmt.Sprintf("%d", i),
					HostId:    "JUAN123",
					Day:       domain.Saturday,
					FromDate:  "2023-03-01",
					ToDate:    "2023-07-01",
					StartTime: "13:00:00",
					EndTime:   "16:00:00",
				}
			} else {
				schedulePatterns[i] = domain.SchedulePattern{
					Id:        fmt.Sprintf("%d", i),
					HostId:    "JUAN123",
					Day:       domain.Sunday,
					FromDate:  "2023-03-01",
					ToDate:    "2023-07-01",
					StartTime: "13:00:00",
					EndTime:   "16:00:00",
				}
			}

		}

		result[i] = domain.Spot{
			SpotId:           fmt.Sprintf("%d", i),
			SchedulePatterns: schedulePatterns,
		}

	}
	return result
}

func (r *LocalRepository) ConditionalDatesCreation(spot domain.Spot) error {
	return nil
}
