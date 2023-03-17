package commands

import (
	_ "embed"
	"log"
	"time"

	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/core/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type CreateDateAssociationCommand struct {
	spot   domain.OnlineSpot
	status domain.SpotStatus
}

var (
	//go:embed queries/append_date.cypher
	associateCypherQuery string
)

func NewCreateDateAssociationCommand(spot domain.OnlineSpot, status domain.SpotStatus) *CreateDateAssociationCommand {
	return &CreateDateAssociationCommand{
		spot:   spot,
		status: status,
	}
}

func (cmd *CreateDateAssociationCommand) Run(tr neo4j.Transaction) (interface{}, error) {
	cyperParams := map[string]interface{}{
		"host_id":                   cmd.spot.SpotInfo.OwnerId,
		"event_uuid":                cmd.spot.SpotInfo.SpotId,
		"date_uuid":                 cmd.spot.DatesInfo[0].DateId,
		"date_approximated_seconds": cmd.spot.DatesInfo[0].DurationApproximatedInSeconds,
		"date_start_time":           cmd.spot.DatesInfo[0].StartTime,
		"date_date":                 cmd.spot.DatesInfo[0].Date,
		"date_confirmed":            cmd.spot.DatesInfo[0].Confirmed,
		"date_maximun_capacity":     cmd.spot.DatesInfo[0].MaximunCapacty,
		"status":                    cmd.status,
		"timestamp":                 time.Now().Unix(),
	}

	_, err := tr.Run(associateCypherQuery, cyperParams)

	if err != nil {
		log.Println("CreateDateAssociationCommand: Run: Error when running the query, error -> ", err.Error())
	}

	return nil, err
}
