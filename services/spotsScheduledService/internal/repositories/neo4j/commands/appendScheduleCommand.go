package commands

import (
	_ "embed"
	"fmt"

	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/logs"
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

	var cyperParams map[string]interface{} = map[string]interface{}{}

	var initalCypherQuery string = `
	MATCH	
		(host:Person {id: $host_id})-[:OWNS]->(event:Event {UUID: $event_uuid })
	
	`
	cyperParams["host_id"] = cmd.scheduleSpot.SpotInfo.SpotId
	cyperParams["event_uuid"] = cmd.scheduleSpot.SpotInfo.OwnerId

	for i, scheduleSpot := range cmd.scheduleSpot.Patterns {
		mergeQuery := fmt.Sprintf(`
		MERGE 
		(schedulePattern_%[1]d:SchedulePattern {UUID: $schedulePattern_uuid_%[1]d})
		ON CREATE
			SET schedulePattern.days = $schedulePattern_days_%[1]d
			SET schedulePattern.fromDate = $schedulePattern_fromDate_%[1]d
			SET schedulePattern.toDate = $schedulePattern_toDate_%[1]d
			SET schedulePattern.StartTime = $schedulePattern_StartTime_%[1]d
			SET schedulePattern.endTime = $schedulePattern_endTime_%[1]d
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
		cyperParams[fmt.Sprintf("status_%[1]d", i)] = scheduleSpot.State.Status
		cyperParams[fmt.Sprintf("timestamp_%[1]d", i)] = scheduleSpot.State.Since

	}

	logs.Info.Printf("cyperParams: %+v\n", cyperParams)
	logs.Info.Printf("initalCypherQuery: %+v\n", initalCypherQuery)

	if _, err := tr.Run(appendQueryCypher, cyperParams); err != nil {
		logs.Error.Println("NewAppendScheduleCommand: Run: Error when running the query, error -> ", err.Error())
		return nil, err
	}

	return nil, nil

}
