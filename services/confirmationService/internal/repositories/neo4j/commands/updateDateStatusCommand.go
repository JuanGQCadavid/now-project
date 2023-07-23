package commands

import (
	_ "embed"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type UpdateDateStatusCommand struct {
	dateId    string
	confirmed bool
}

var (
	//go:embed queries/updateStatus.cypher
	updateStatus string
)

func NewUpdateDateStatusCommand(dateId string, confirmed bool) *UpdateDateStatusCommand {
	return &UpdateDateStatusCommand{
		dateId:    dateId,
		confirmed: confirmed,
	}
}

func (cmd *UpdateDateStatusCommand) Run(trx neo4j.Transaction) (interface{}, error) {

	queryParams := map[string]interface{}{
		"dateId":    cmd.dateId,
		"confirmed": cmd.confirmed,
	}

	_, err := trx.Run(updateStatus, queryParams)

	if err != nil {
		logs.Error.Println("Error on UpdateDateStatusCommand trx.Run", err.Error())
		return nil, err
	}

	return nil, nil
}
