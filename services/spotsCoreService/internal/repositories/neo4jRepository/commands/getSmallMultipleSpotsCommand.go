package commands

import (
	"log"

	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
)

type GetSmallMultipleSpotsCommand struct {
	spotIds []string
}

func (command *GetSmallMultipleSpotsCommand) Run(tr neo4j.Transaction) (interface{}, error) {

	var cypherQ string = `
	MATCH
		(host:Person)-[host_relation:OWNS]->(event:Event)-[location_relation:ON]->(place:Place)
	WHERE
		event.UUID IN $spotIds
		AND NOT (event)-[:IS_DELETED]->(event)
	OPTIONAL MATCH (tags:Topic)-[tagged:TAGGED]->(event)
	RETURN
		event.name as event_name,
		event.UUID as event_UUID,
		event.emoji as event_emoji,
		place.lon as place_lon,
		place.mapProviderId as place_provider_id,
		place.lat as place_lat,
		collect(tags.tag) as tag_tags,
		collect(tagged.isPrincipal) as tag_principals
	`

	println(cypherQ)

	cyperParams := map[string]interface{}{"spotIds": command.spotIds}

	result, err := tr.Run(cypherQ, cyperParams)

	var spotsToReturn []domain.Spot = []domain.Spot{}

	if err != nil {
		println("Error at running!", err)
		return &domain.MultipleSpots{}, err
	}

	for result.Next() {
		record := result.Record()
		spot := command.getSpotDataFromResult(record)
		spotsToReturn = append(spotsToReturn, spot)
	}

	return &domain.MultipleSpots{
		Spots: spotsToReturn,
	}, nil

}

func (command *GetSmallMultipleSpotsCommand) getSpotDataFromResult(record *db.Record) domain.Spot {
	// Event
	event_name, _ := record.Get("event_name")
	event_UUID, _ := record.Get("event_UUID")
	event_emoji, _ := record.Get("event_emoji")

	// Place
	place_lon, _ := record.Get("place_lon")
	place_provider_id, _ := record.Get("place_provider_id")
	place_lat, _ := record.Get("place_lat")

	// Tags
	tags_ids, _ := record.Get("tag_tags")
	tags_principals, _ := record.Get("tag_principals")
	tags_principals_array := tags_principals.([]interface{})

	secondary_tag := make([]string, 0, len(tags_principals_array))
	primary_tag := ""

	for index, tag := range tags_ids.([]interface{}) {
		if tags_principals_array[index].(bool) {
			primary_tag = tag.(string)
		} else {
			secondary_tag = append(secondary_tag, tag.(string))
		}

	}

	log.Printf("%+v", record)

	return domain.Spot{
		EventInfo: domain.Event{
			Name:  getStringFromInterface(event_name),
			UUID:  getStringFromInterface(event_UUID),
			Emoji: getStringFromInterface(event_emoji),
		},
		PlaceInfo: domain.Place{
			Lat:           getFloat64FromInterface(place_lat),
			Lon:           getFloat64FromInterface(place_lon),
			MapProviderId: getStringFromInterface(place_provider_id),
		},
		TopicsInfo: domain.Topic{
			PrincipalTopic:  primary_tag,
			SecondaryTopics: secondary_tag,
		},
	}

}

func NewGetSmallMultipleSpotsCommand(spotIds []string) *GetSmallMultipleSpotsCommand {
	return &GetSmallMultipleSpotsCommand{
		spotIds: spotIds,
	}
}
