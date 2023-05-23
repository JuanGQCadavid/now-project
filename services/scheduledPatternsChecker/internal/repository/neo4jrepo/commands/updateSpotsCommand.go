package commands

import (
	"bytes"
	_ "embed"
	"fmt"
	"time"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/ports"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type UpdateSpotsCommand struct {
	spots       []domain.Spot
	datesStatus string
}

func NewUpdateSpotsCommand(spots []domain.Spot, datesStatus string) *UpdateSpotsCommand {
	return &UpdateSpotsCommand{
		spots:       spots,
		datesStatus: datesStatus,
	}
}

func (cmd *UpdateSpotsCommand) Run(tr neo4j.Transaction) (interface{}, error) {

	var updateQuery bytes.Buffer
	queryParams := make(map[string]interface{})
	var now int64 = time.Now().Unix()

	for spotIndex, spot := range cmd.spots {

		scheduleCommand := `
			MATCH (schedulePattern_%[1]d:SchedulePattern {UUID: $schedule_pattern_id_%[1]d})
			SET schedulePattern_%[1]d.checkedUpTo = $checkedUpTo_%[1]d
			WITH schedulePattern_%[1]d
		`

		for scheduleIndex, schedulePattern := range spot.SchedulePatterns {
			fmt.Fprintf(&updateQuery, scheduleCommand, scheduleIndex)

			queryParams[fmt.Sprintf("schedule_pattern_id_%[1]d", scheduleIndex)] = schedulePattern.Id
			queryParams[fmt.Sprintf("checkedUpTo_%[1]d", scheduleIndex)] = schedulePattern.CheckedUpTo
		}

		matchCommand := `
			MATCH (host_%[1]d:Person {id: $host_id_%[1]d})-[host_relation_%[1]d:OWNS]->(event_%[1]d:Event {UUID: $event_uuid_%[1]d } )
		`
		dateCreationCommand := `
			MERGE (date_%[1]d:Date {StartTime: $date_start_time_%[1]d, Date: $date_date_%[1]d })
			ON CREATE
				SET date_%[1]d.UUID = $date_uuid_%[1]d
				SET date_%[1]d.DurationApproximatedInSeconds = $date_approximated_seconds_%[1]d
				SET date_%[1]d.Confirmed = $date_confirmed_%[1]d
				SET date_%[1]d.MaximunCapacty = $date_maximun_capacity_%[1]d
		`
		joinCommand := `
			MERGE (host_%[1]d)-[:HOST]->(date_%[2]d)-[at_%[1]d:AT {status: $status}]->(event_%[1]d)
			ON CREATE
				SET at_%[1]d.timestamp = $timestamp 
		`

		if len(spot.Dates) > 0 {

			fmt.Fprintf(&updateQuery, matchCommand, spotIndex)
			queryParams[fmt.Sprintf("event_uuid_%[1]d", spotIndex)] = spot.SpotId
			queryParams["status"] = cmd.datesStatus
			queryParams["timestamp"] = now

		}

		var hostAdded bool = false

		for dateIndex, date := range spot.Dates {

			if !hostAdded {
				queryParams[fmt.Sprintf("host_id_%[1]d", spotIndex)] = date.HostId
				hostAdded = true
			}

			baseCommand := fmt.Sprintf(dateCreationCommand, dateIndex)

			queryParams[fmt.Sprintf("date_uuid_%d", dateIndex)] = date.DateId
			queryParams[fmt.Sprintf("date_approximated_seconds_%d", dateIndex)] = date.DurationApproximatedInSeconds
			queryParams[fmt.Sprintf("date_start_time_%d", dateIndex)] = date.StartTime
			queryParams[fmt.Sprintf("date_date_%d", dateIndex)] = date.Date
			queryParams[fmt.Sprintf("date_confirmed_%d", dateIndex)] = false
			queryParams[fmt.Sprintf("date_maximun_capacity_%d", dateIndex)] = date.MaximunCapacty

			fmt.Fprintln(&updateQuery, baseCommand)
			fmt.Fprintf(&updateQuery, joinCommand, spotIndex, dateIndex)
		}

	}

	logs.Info.Printf("Command -> \n%s \nParams -> %+v", updateQuery.String(), queryParams)

	_, err := tr.Run(updateQuery.String(), queryParams)

	if err != nil {
		logs.Error.Println("Run command fail, error -> ", err.Error())
		return nil, ports.ErrFetchingData
	}

	return nil, nil
}
