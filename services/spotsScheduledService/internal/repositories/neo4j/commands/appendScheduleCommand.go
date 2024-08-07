package commands

import (
	_ "embed"
	"fmt"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/ports"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type AppendScheduleCommand struct {
	scheduleSpot domain.ScheduledSpot
}

var (
	//go:embed queries/appendSchedule.cypher
	appendQueryCypher string
)

func NewAppendScheduleCommand(scheduleSpot domain.ScheduledSpot) *AppendScheduleCommand {
	return &AppendScheduleCommand{
		scheduleSpot: scheduleSpot,
	}
}

// ERROR -> We should be able to insert more than one sp at the time

func (cmd *AppendScheduleCommand) Run(tr neo4j.Transaction) (interface{}, error) {
	logs.Info.Println("++++++++", cmd.scheduleSpot.Patterns)
	var cyperParams map[string]interface{} = map[string]interface{}{}

	var initalCypherQuery string = `
	MATCH	
		(host:Person {id: $host_id})-[:OWNS]->(event:Event {UUID: $event_uuid })
	
	`
	cyperParams["host_id"] = cmd.scheduleSpot.SpotInfo.OwnerId
	cyperParams["event_uuid"] = cmd.scheduleSpot.SpotInfo.SpotId

	for i, scheduleSpot := range cmd.scheduleSpot.Patterns {
		logs.Info.Println("-------", i)
		if len(scheduleSpot.Host.HostId) > 0 {
			cyperParams["host_id"] = scheduleSpot.Host.HostId
		} else {
			cyperParams["host_id"] = cmd.scheduleSpot.SpotInfo.OwnerId
		}

		mergeQuery := fmt.Sprintf(`
		MERGE 
		(schedulePattern_%[1]d:SchedulePattern {UUID: $schedulePattern_uuid_%[1]d})
		ON CREATE
			SET schedulePattern_%[1]d.days = $schedulePattern_days_%[1]d
			SET schedulePattern_%[1]d.fromDate = $schedulePattern_fromDate_%[1]d
			SET schedulePattern_%[1]d.toDate = $schedulePattern_toDate_%[1]d
			SET schedulePattern_%[1]d.StartTime = $schedulePattern_StartTime_%[1]d
			SET schedulePattern_%[1]d.endTime = $schedulePattern_endTime_%[1]d
		MERGE 
		(host)<-[:HOST_BY]-(schedulePattern_%[1]d)-[:AT {status: $status_%[1]d, timestamp: $timestamp_%[1]d }]->(event)
		`, i)

		initalCypherQuery += "\n" + mergeQuery

		cyperParams[fmt.Sprintf("schedulePattern_uuid_%[1]d", i)] = scheduleSpot.Id
		cyperParams[fmt.Sprintf("schedulePattern_days_%[1]d", i)] = scheduleSpot.Day
		cyperParams[fmt.Sprintf("schedulePattern_fromDate_%[1]d", i)] = scheduleSpot.FromDate
		cyperParams[fmt.Sprintf("schedulePattern_toDate_%[1]d", i)] = scheduleSpot.ToDate
		cyperParams[fmt.Sprintf("schedulePattern_StartTime_%[1]d", i)] = scheduleSpot.StartTime
		cyperParams[fmt.Sprintf("schedulePattern_endTime_%[1]d", i)] = scheduleSpot.EndTime
		cyperParams[fmt.Sprintf("status_%[1]d", i)] = string(scheduleSpot.State.Status)
		cyperParams[fmt.Sprintf("timestamp_%[1]d", i)] = scheduleSpot.State.Since

	}

	logs.Info.Printf("cyperParams: %+v\n", cyperParams)
	logs.Info.Printf("initalCypherQuery: %+v\n", initalCypherQuery)

	if _, err := tr.Run(initalCypherQuery, cyperParams); err != nil {
		logs.Error.Println("NewAppendScheduleCommand: Run: Error when running the query, error -> ", err.Error())
		return nil, ports.ErrOnRepository
	}

	return nil, nil

}
