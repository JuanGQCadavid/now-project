package commands

import (
	_ "embed"
	"time"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type UpdateSpotsCommand struct {
	spots       []domain.Spot
	datesStatus string
}

var (
	//go:embed queries/updateAllSpots.cypher
	updateAllSpots string
)

func NewUpdateSpotsCommand(spots []domain.Spot, datesStatus string) *UpdateSpotsCommand {
	return &UpdateSpotsCommand{
		spots:       spots,
		datesStatus: datesStatus,
	}
}

func (cmd *UpdateSpotsCommand) Run(tr neo4j.Transaction) (interface{}, error) {

	var now int64 = time.Now().Unix()
	queryParams := make(map[string]interface{})
	schedulePatterns := make([]map[string]interface{}, 0, 10)

	for _, spot := range cmd.spots {

		for _, schedulePattern := range spot.SchedulePatterns {
			props := make([]map[string]interface{}, len(schedulePattern.Dates))

			for dateIndex, date := range schedulePattern.Dates {
				props[dateIndex] = map[string]interface{}{
					"StartTime":                     date.StartTime,
					"Date":                          date.Date,
					"Id":                            date.DateId,
					"DurationApproximatedInSeconds": date.DurationApproximatedInSeconds,
					"Confirmed":                     false,
					"MaximunCapacty":                date.MaximunCapacty,
					"Status":                        cmd.datesStatus,
					"Timestamp":                     now,
				}
			}

			schedulePatterns = append(schedulePatterns, map[string]interface{}{
				"Id":          schedulePattern.Id,
				"HostId":      schedulePattern.HostId,
				"CheckedUpTo": schedulePattern.CheckedUpTo,
				"Dates":       props,
				"SpotId":      spot.SpotId,
			})
		}
	}

	queryParams["schedulePatterns"] = schedulePatterns

	logs.Info.Printf("Query %s \nParams -> %+v\n", updateAllSpots, queryParams)

	_, err := tr.Run(updateAllSpots, queryParams)

	if err != nil {
		logs.Error.Printf("Run command fail, error -> %s \n", err.Error())
	}

	return nil, err
}
