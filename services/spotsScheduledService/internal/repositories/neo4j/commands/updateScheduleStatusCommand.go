package commands

import (
	_ "embed"

	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/logs"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/ports"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var (
	//go:embed queries/updateScheduleStatus.cypher
	updateScheduleStatusQueryCypher string
)

type UpdateScheduleStatusCommand struct {
	schedulePatternId string
	spotId            string
	status            domain.State
}

func NewUpdateScheduleStatusCommand(schedulePatternId string, spotId string, status domain.State) *UpdateScheduleStatusCommand {
	return &UpdateScheduleStatusCommand{
		schedulePatternId: schedulePatternId,
		spotId:            spotId,
		status:            status,
	}
}

func (cmd *UpdateScheduleStatusCommand) Run(tr neo4j.Transaction) (interface{}, error) {

	var queryParams map[string]interface{} = map[string]interface{}{
		"spot_uuid":            cmd.spotId,
		"schedulePattern_uuid": cmd.schedulePatternId,
		"status":               string(cmd.status.Status),
		"timestamp":            cmd.status.Since,
	}

	_, err := tr.Run(updateScheduleStatusQueryCypher, queryParams)

	if err != nil {
		logs.Error.Println("We found an error whule updating the status, ", err.Error())
		return nil, ports.ErrOnRepository
	}

	return nil, nil
}
