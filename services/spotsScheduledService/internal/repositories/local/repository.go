package local

import "github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/domain"

type LocalRepository struct {
}

func (repo *LocalRepository) GetScheduleSpot(spotId string, flags domain.ScheduleStateFlags) (*domain.ScheduledSpot, error) {
	patterns := []domain.SchedulePattern{
		{
			Id: "3",
			State: domain.State{
				Status: domain.ACTIVATE,
			},
			Host: domain.Host{
				HostId:   "JUAN",
				HostName: "JUAN",
			},
			Day:       domain.Monday,
			FromDate:  "2007-03-01",
			ToDate:    "2007-07-01",
			StartTime: "13:00:00",
			EndTime:   "16:00:00",
		},

		{
			Id: "2",
			State: domain.State{
				Status: domain.ACTIVATE,
			},
			Host: domain.Host{
				HostId:   "JUAN",
				HostName: "JUAN",
			},
			Day:       domain.Friday,
			FromDate:  "2007-03-01",
			ToDate:    "2007-07-01",
			StartTime: "13:00:00",
			EndTime:   "16:00:00",
		},

		{
			Id: "1",
			State: domain.State{
				Status: domain.ACTIVATE,
			},
			Host: domain.Host{
				HostId:   "JUAN",
				HostName: "JUAN",
			},
			Day:       domain.Saturday,
			FromDate:  "2007-03-01",
			ToDate:    "2007-07-01",
			StartTime: "13:00:00",
			EndTime:   "16:00:00",
		},
	}
	return &domain.ScheduledSpot{
		SpotInfo: domain.SpotInfo{
			SpotId:  spotId,
			OwnerId: "JUAN",
		},
		Patterns: patterns,
	}, nil
}
func (repo *LocalRepository) AssociateSpotWithSchedulePatterns(spotId string, hostId string, schedulesPattern *[]domain.SchedulePattern) error {

	return nil
}

func (repo *LocalRepository) UpdateScheculeStatus(spotId string, scheduleId string, status domain.State) error {
	return nil
}
