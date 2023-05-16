package commands

import (
	_ "embed"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/ports"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FetchAllSchedulePatternsCommand struct {
}

func NewFetchAllSchedulePatternsCommand() *FetchAllSchedulePatternsCommand {
	return &FetchAllSchedulePatternsCommand{}
}

var (
	//go:embed queries/getAllSchedulePatterns.cypher
	getAllSchedulePatterns string
)

func (cmd *FetchAllSchedulePatternsCommand) Run(tr neo4j.Transaction) (interface{}, error) {

	result, err := tr.Run(getAllSchedulePatterns, nil)

	if err != nil {
		logs.Error.Println("Run command fail, error -> ", err.Error())
		return nil, ports.ErrFetchingData
	}

	spots := make(map[string]domain.Spot)

	for result.Next() {
		newSpot, err := cmd.processRecord(result.Record())

		if err != nil {
			logs.Error.Printf("We fail to process a record %+v \n", result.Record())
			logs.Error.Println("error -> ", err.Error())
			continue
		}

		spots[newSpot.SpotId] = newSpot

	}

	return spots, nil
}

func (cmd *FetchAllSchedulePatternsCommand) processRecord(record *neo4j.Record) (domain.Spot, error) {

	// event.UUID as event_UUID,
	// collect(
	// 	{
	//         schedulePattern_id: schedulePattern.UUID,
	// 		schedulePattern_days: schedulePattern.days,
	// 		schedulePattern_fromDate: schedulePattern.fromDate,
	// 		schedulePattern_toDate: schedulePattern.toDate,
	// 		schedulePattern_StartTime: schedulePattern.StartTime,
	// 		schedulePattern_endTime: schedulePattern.endTime,
	// 		schedulePattern_checkedUpTo: schedulePattern.checkedUpTo,
	// 		host_id: host.id
	// 	}
	// ) as schedulePatterns

	spotId, _ := record.Get("event_UUID")

	schedulePatterns, containsSchedulePatterns := record.Get("schedulePatterns")
	var patternsData []domain.SchedulePattern = nil

	if containsSchedulePatterns && schedulePatterns != nil {
		for _, pattern := range schedulePatterns.([]interface{}) {
			patternData := pattern.(map[string]interface{})

			if patternData["schedulePattern_id"] == nil {
				continue
			}

			patternId := patternData["schedulePattern_id"].(string)

			if len(patternId) == 0 {
				continue
			}

			patternDay := patternData["schedulePattern_days"].(int64)
			patternFromDate := patternData["schedulePattern_fromDate"].(string)
			patternToDate := patternData["schedulePattern_toDate"].(string)
			patternStartTime := patternData["schedulePattern_StartTime"].(string)
			patternEndTime := patternData["schedulePattern_endTime"].(string)
			var patternCheckedUpTo int64 = 0

			patternCheckedUpToInterface := patternData["schedulePattern_checkedUpTo"]

			if patternCheckedUpToInterface != nil {
				patternCheckedUpTo = patternCheckedUpToInterface.(int64)
			}

			patternHostId := patternData["host_id"].(string)

			if patternsData == nil {
				patternsData = []domain.SchedulePattern{}
			}

			patternsData = append(patternsData, domain.SchedulePattern{
				Id:          patternId,
				Day:         domain.Day(patternDay),
				FromDate:    patternFromDate,
				ToDate:      patternToDate,
				StartTime:   patternStartTime,
				EndTime:     patternEndTime,
				HostId:      patternHostId,
				CheckedUpTo: patternCheckedUpTo,
			})

		}
	}

	return domain.Spot{
		SpotId:           spotId.(string),
		SchedulePatterns: patternsData,
	}, nil
}
