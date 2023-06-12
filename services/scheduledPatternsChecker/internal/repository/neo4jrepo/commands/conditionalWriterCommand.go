package commands

import (
	_ "embed"
	"time"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type ConditionalWritterCommand struct {
	spot        domain.Spot
	datesStatus string
}

func NewConditionalWritterCommand(spot domain.Spot, datesStatus string) *ConditionalWritterCommand {
	return &ConditionalWritterCommand{
		spot:        spot,
		datesStatus: datesStatus,
	}
}

var (
	//go:embed queries/multipleSpConditionalWirter.cypher
	multipleConditionalWrite string
)

func (cmd *ConditionalWritterCommand) Run(tr neo4j.Transaction) (interface{}, error) {

	var now int64 = time.Now().Unix()
	queryParams := make(map[string]interface{})
	schedulePatterns := make([]map[string]interface{}, len(cmd.spot.SchedulePatterns))

	for spIndex, schedulePattern := range cmd.spot.SchedulePatterns {
		props := make([]map[string]interface{}, len(schedulePattern.Dates))

		for dateIndex, date := range schedulePattern.Dates {
			props[dateIndex] = map[string]interface{}{
				"StartTime":                     date.StartTime,
				"Date":                          date.Date,
				"UUID":                          date.DateId,
				"DurationApproximatedInSeconds": date.DurationApproximatedInSeconds,
				"Confirmed":                     false,
				"MaximunCapacty":                date.MaximunCapacty,
				"Status":                        cmd.datesStatus,
				"Timestamp":                     now,
			}
		}

		schedulePatterns[spIndex] = map[string]interface{}{
			"id":          schedulePattern.Id,
			"hostId":      schedulePattern.HostId,
			"days":        schedulePattern.Day,
			"endTime":     schedulePattern.EndTime,
			"startTime":   schedulePattern.StartTime,
			"fromDate":    schedulePattern.FromDate,
			"toDate":      schedulePattern.ToDate,
			"checkedUpTo": schedulePattern.CheckedUpTo,
			"datesProps":  props,
		}
	}
	queryParams["spotId"] = cmd.spot.SpotId
	queryParams["schedulePatterns"] = schedulePatterns

	logs.Info.Printf("Inserting %s \nParams -> %+v\n", cmd.spot.SpotId, queryParams)

	_, err := tr.Run(multipleConditionalWrite, queryParams)

	if err != nil {
		logs.Error.Printf("Run command fail, error -> %s \n", err.Error())
	}

	return nil, err

}
