package commands

import (
	_ "embed"

	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/logs"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/ports"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type GetSchedulesCommand struct {
	spotId string
}

var (
	//go:embed queries/getSchedulePattern.cypher
	getSchedulePatternQueryCypher string
)

func NewGetSchedulesCommand(spotId string) *GetSchedulesCommand {
	return &GetSchedulesCommand{
		spotId: spotId,
	}
}

func (cmd *GetSchedulesCommand) Run(tr neo4j.Transaction) (interface{}, error) {

	var params map[string]interface{} = map[string]interface{}{
		"spot_uuid": cmd.spotId,
	}

	output, err := tr.Run(getSchedulePatternQueryCypher, params)

	if err != nil {
		logs.Error.Println("GetSchedulesCommand: Run: Error when running the query, error -> ", err.Error())
		logs.Error.Printf("Cypher: %s \nParams: %s\n", getSchedulePatternQueryCypher, params)
		return nil, err
	}

	for output.Next() {
		record := output.Record()
		return cmd.castOutput(record)
	}

	return nil, ports.ErrNotRecordsToProcess
}

func (cmd *GetSchedulesCommand) castOutput(record *neo4j.Record) (domain.ScheduledSpot, error) {
	spotId, _ := record.Get("event_UUID")
	ownerId, _ := record.Get("owner_id")

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

			pattenrState := patternData["state"].(map[string]interface{})
			patternSince := pattenrState["since"].(int64)
			var patternStatus domain.Status

			switch patternStatusType := pattenrState["status"].(string); patternStatusType {
			case string(domain.ACTIVATE):
				patternStatus = domain.ACTIVATE
			case string(domain.CONCLUDE):
				patternStatus = domain.CONCLUDE
			case string(domain.FREEZE):
				patternStatus = domain.FREEZE
			default:
				return domain.ScheduledSpot{}, ports.ErrSpotStatusUndefined
			}

			pattenrHostedBy := patternData["hosted_by"].(map[string]interface{})
			patternHostId := pattenrHostedBy["host_id"].(string)
			patternHostName := pattenrHostedBy["host_name"].(string)

			if patternsData == nil {
				patternsData = []domain.SchedulePattern{}
			}

			patternsData = append(patternsData, domain.SchedulePattern{
				Id:        patternId,
				Day:       domain.Day(patternDay),
				FromDate:  patternFromDate,
				ToDate:    patternToDate,
				StartTime: patternStartTime,
				EndTime:   patternEndTime,
				State: domain.State{
					Status: patternStatus,
					Since:  patternSince,
				},
				Host: domain.Host{
					HostId:   patternHostId,
					HostName: patternHostName,
				},
			})

		}
	}

	return domain.ScheduledSpot{
		SpotInfo: domain.SpotInfo{
			SpotId:  spotId.(string),
			OwnerId: ownerId.(string),
		},
		Patterns: patternsData,
	}, nil
}
