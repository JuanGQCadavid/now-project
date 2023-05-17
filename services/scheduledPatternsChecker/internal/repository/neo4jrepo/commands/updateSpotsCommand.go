package commands

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
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

var (
	//go:embed queries/updateSchedulePattern.cypher
	updateSchedulePattern string

	//go:embed queries/addDatesToSpot.cypher
	addDatesToSpot string
)

func (cmd *UpdateSpotsCommand) Run(tr neo4j.Transaction) (interface{}, error) {

	var updateQuery bytes.Buffer
	var queryParams map[string]interface{}
	var now int64 = time.Now().Unix()

	for _, spot := range cmd.spots {

		for scheduleIndex, schedulePattern := range spot.SchedulePatterns {

			// MATCH (event:Event {UUID: $%s } )<-[at:AT]-(schedulePattern:SchedulePattern {UUID: $%s})
			// SET schedulePattern.checkedUpTo = $%d

			spotId := fmt.Sprintf("event_id_%d", scheduleIndex)
			spId := fmt.Sprintf("schedule_pattern_id_%d", scheduleIndex)
			checkedUpTo := fmt.Sprintf("checked_up_to_%d", scheduleIndex)

			fmt.Fprintf(&updateQuery, updateSchedulePattern, spotId, spId, checkedUpTo)

			queryParams[spotId] = spot.SpotId
			queryParams[spId] = schedulePattern.Id
			queryParams[checkedUpTo] = schedulePattern.CheckedUpTo
		}

		for dateIndex, date := range spot.Dates {
			baseCommand := addDatesToSpot

			// MATCH
			// 	(host:Person {id: $host_id})-[host_relation:OWNS]->(event:Event {UUID: $event_uuid } )
			// MERGE
			// (date:Date {UUID: $date_uuid})
			// ON CREATE
			// 	SET date.DurationApproximatedInSeconds = $date_approximated_seconds
			// 	SET date.StartTime = $date_start_time
			// 	SET date.Date = $date_date
			// 	SET date.Confirmed = $date_confirmed
			// 	SET date.MaximunCapacty = $date_maximun_capacity
			// MERGE
			// (host)-[:HOST]->(date)-[:AT {status: $status, timestamp: $timestamp }]->(event)

			spotId := fmt.Sprintf("$event_uuid_%d", dateIndex)
			baseCommand = strings.ReplaceAll(baseCommand, "$event_uuid", spotId)
			queryParams[spotId] = spot.SpotId

			hostId := fmt.Sprintf("$host_id_%d", dateIndex)
			baseCommand = strings.ReplaceAll(baseCommand, "$host_id", hostId)
			queryParams[hostId] = date.HostId

			dateId := fmt.Sprintf("$date_uuid_%d", dateIndex)
			baseCommand = strings.ReplaceAll(baseCommand, "$date_uuid", dateId)
			queryParams[dateId] = date.DateId

			approxSeconds := fmt.Sprintf("$date_approximated_seconds_%d", dateIndex)
			baseCommand = strings.ReplaceAll(baseCommand, "$date_approximated_seconds", approxSeconds)
			queryParams[approxSeconds] = date.DurationApproximatedInSeconds

			startTime := fmt.Sprintf("$date_start_time_%d", dateIndex)
			baseCommand = strings.ReplaceAll(baseCommand, "$date_start_time", startTime)
			queryParams[startTime] = date.StartTime

			dateDate := fmt.Sprintf("$date_date_%d", dateIndex)
			baseCommand = strings.ReplaceAll(baseCommand, "$date_date", dateDate)
			queryParams[dateDate] = date.Date

			dateConfirmed := fmt.Sprintf("$date_confirmed_%d", dateIndex)
			baseCommand = strings.ReplaceAll(baseCommand, "$date_confirmed", dateConfirmed)
			queryParams[dateConfirmed] = false

			maximunCapacity := fmt.Sprintf("$date_maximun_capacity_%d", dateIndex)
			baseCommand = strings.ReplaceAll(baseCommand, "$date_maximun_capacity", maximunCapacity)
			queryParams[maximunCapacity] = date.MaximunCapacty

			dateStatus := fmt.Sprintf("$status_%d", dateIndex)
			baseCommand = strings.ReplaceAll(baseCommand, "$status", dateStatus)
			queryParams[dateStatus] = cmd.datesStatus

			dateTimeStamp := fmt.Sprintf("$timestamp_%d", dateIndex)
			baseCommand = strings.ReplaceAll(baseCommand, "$timestamp", dateTimeStamp)
			queryParams[dateTimeStamp] = now

			updateQuery.WriteString(baseCommand)
		}

	}

	logs.Info.Printf("Command -> %s \nParams -> %+v", updateQuery.String(), queryParams)

	_, err := tr.Run(updateQuery.String(), queryParams)

	if err != nil {
		logs.Error.Println("Run command fail, error -> ", err.Error())
		return nil, ports.ErrFetchingData
	}

	return nil, nil
}
