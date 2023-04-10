package commands

import (
	_ "embed"

	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/domain"
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

	return nil, nil
}
