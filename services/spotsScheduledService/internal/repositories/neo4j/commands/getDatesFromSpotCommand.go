package commands

import (
	_ "embed"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/ports"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type GetDatesFromSpotCommand struct {
	spotId string
}

var (
	//go:embed queries/getDatesFromSpot.cypher
	getDatesFromSpot string
)

func NewGetDatesFromSpotCommand(spotId string) *GetDatesFromSpotCommand {
	return &GetDatesFromSpotCommand{
		spotId: spotId,
	}
}

func (cmd *GetDatesFromSpotCommand) Run(tr neo4j.Transaction) (interface{}, error) {

	var params map[string]interface{} = map[string]interface{}{
		"spot_uuid": cmd.spotId,
	}

	output, err := tr.Run(getDatesFromSpot, params)

	if err != nil {
		logs.Error.Println("GetDatesFromSpotCommand: Run: Error when running the query, error -> ", err.Error())
		logs.Error.Printf("Cypher: %s \nParams: %s\n", getSchedulePatternQueryCypher, params)
		return nil, err
	}

	for output.Next() {
		record := output.Record()
		return cmd.castOutput(record)
	}

	return nil, ports.ErrNotRecordsToProcess
}

func (cmd *GetDatesFromSpotCommand) castOutput(record *neo4j.Record) ([]domain.Date, error) {

	dates, containsDates := record.Get("dates")

	var datesData []domain.Date = nil

	if containsDates && dates != nil {
		for _, pattern := range dates.([]interface{}) {
			patternData := pattern.(map[string]interface{})

			if patternData["dateUUID"] == nil {
				continue
			}

			dateUUID := patternData["dateUUID"].(string)

			if len(dateUUID) == 0 {
				continue
			}

			dateMaximunCapacity := patternData["dateMaximunCapacity"].(int64)
			dateDurationApproximatedInSeconds := patternData["dateDurationApproximatedInSeconds"].(int64)
			dateStartTime := patternData["dateStartTime"].(string)
			dateConfirmed := patternData["dateConfirmed"].(bool)
			dateDate := patternData["dateDate"].(string)
			schedulePatternId := patternData["schedulePatternId"].(string)

			hostedBy := patternData["hostedBy"].(map[string]interface{})
			hostId := hostedBy["hostId"].(string)
			hostName := hostedBy["hostName"].(string)

			dateState := patternData["dateState"].(map[string]interface{})
			status := dateState["status"].(string)
			timestamp := dateState["timestamp"].(int64)

			var dateStatus domain.Status

			switch status {
			case string(domain.SCHEDULED):
				dateStatus = domain.ACTIVATE
			default:
				return nil, ports.ErrSpotStatusUndefined
			}

			if datesData == nil {
				datesData = []domain.Date{}
			}

			datesData = append(datesData, domain.Date{
				MaximunCapacty:                dateMaximunCapacity,
				DurationApproximatedInSeconds: dateDurationApproximatedInSeconds,
				StartTime:                     dateStartTime,
				Confirmed:                     dateConfirmed,
				Id:                            dateUUID,
				DateStamp:                     dateDate,
				FromSchedulePattern: domain.SchedulePattern{
					Id: schedulePatternId,
				},
				State: domain.State{
					Status: dateStatus,
					Since:  timestamp,
				},
				Host: domain.Host{
					HostId:   hostId,
					HostName: hostName,
				},
			})

		}
	}

	return datesData, nil
}
