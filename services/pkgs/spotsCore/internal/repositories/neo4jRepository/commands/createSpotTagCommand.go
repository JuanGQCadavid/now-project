package commands

import (
	"fmt"
	"log"

	"github.com/JuanGQCadavid/now-project/services/spotsCore/internal/core/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type CreateSpotTagsCommand struct {
	spotId        string
	principalTag  domain.Optional
	secondaryTags []string
}

func NewCreateSpotTagsCommand(spotId string, principalTag domain.Optional, secondaryTags []string) *CreateSpotTagsCommand {
	return &CreateSpotTagsCommand{
		spotId:        spotId,
		principalTag:  principalTag,
		secondaryTags: secondaryTags,
	}
}

func (cmd *CreateSpotTagsCommand) Run(tx neo4j.Transaction) (interface{}, error) {

	var cypherBaseCommand string = "MATCH (spot:Event {UUID: $event_uuid})"

	var params map[string]interface{} = make(map[string]interface{})
	params["event_uuid"] = cmd.spotId

	if cmd.principalTag.IsPresent() {
		cyperPrimaryTagCommand := "MERGE (primaryTag:Topic {tag: $primaryTag })\nMERGE (primaryTag)-[:TAGGED {isPrincipal:true}]->(spot)"

		cypherBaseCommand = fmt.Sprintf("%s\n%s", cypherBaseCommand, cyperPrimaryTagCommand)
		params["primaryTag"] = cmd.principalTag.GetValue()
	}

	if cmd.secondaryTags != nil {
		for index, tag := range cmd.secondaryTags {
			tagKey := fmt.Sprintf("secondaryTag%d", index)
			cyperSecondaryTagCommandCreation := fmt.Sprintf("MERGE (%s:Topic {tag: $%s })", tagKey, tagKey)
			cyperSecondaryTagCommandLinking := fmt.Sprintf("MERGE (%s)-[:TAGGED {isPrincipal:false}]->(spot)", tagKey)
			cypherBaseCommand = fmt.Sprintf("%s\n%s\n%s", cypherBaseCommand, cyperSecondaryTagCommandCreation, cyperSecondaryTagCommandLinking)
			params[tagKey] = tag
		}
	}

	log.Println(cypherBaseCommand)

	return tx.Run(cypherBaseCommand, params)
}
