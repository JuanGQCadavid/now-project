package commands

import (
	"fmt"
	"log"

	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type AddSelfRelationship struct {
	state  domain.SpotStates
	spotId string
}

func NewAddSelfRelationship(spotId string, state domain.SpotStates) *AddSelfRelationship {
	return &AddSelfRelationship{
		state:  state,
		spotId: spotId,
	}
}

func (cmd *AddSelfRelationship) Run(trx neo4j.Transaction) (interface{}, error) {
	log.Println("AddSelfRelationship: Run ")
	var cypherQ string = fmt.Sprintf(`
		MATCH (e:Event {UUID: $spotId})
		MERGE (e)<-[:%s]-(e)
		return e
		`, cmd.state)

	log.Println("cypherQ: ", cypherQ)
	var params map[string]interface{} = map[string]interface{}{
		"spotId": cmd.spotId,
	}
	_, err := trx.Run(cypherQ, params)

	if err != nil {
		log.Println("commands - AddSelfRelationship - Run: We face an error when adding relationshipt the spot event, err: ", err.Error())
		return nil, err
	}
	return nil, nil
}
