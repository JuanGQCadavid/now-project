package commands

import (
	_ "embed"
	"log"
	"time"

	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/core/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type ChangeAtStatusCommand struct {
	spotId string
	dateId string
	status domain.SpotStatus
}

var (
	//go:embed queries/change_at_status.cypher
	changeAtStatusQuery string
)

func NewChangeAtStatusCommand(spotId string, dateId string, status domain.SpotStatus) *ChangeAtStatusCommand {
	return &ChangeAtStatusCommand{
		spotId: spotId,
		dateId: dateId,
		status: status,
	}
}

func (cmd *ChangeAtStatusCommand) Run(tr neo4j.Transaction) (interface{}, error) {

	queryParams := map[string]interface{}{
		"spot_uuid": cmd.spotId,
		"date_uuid": cmd.dateId,
		"status":    cmd.status,
		"timestamp": time.Now().Unix(),
	}

	_, err := tr.Run(changeAtStatusQuery, queryParams)

	if err != nil {
		log.Println("ChangeAtStatusCommand: Run: Error when running the query, error -> ", err.Error())
	}

	return nil, err
}
