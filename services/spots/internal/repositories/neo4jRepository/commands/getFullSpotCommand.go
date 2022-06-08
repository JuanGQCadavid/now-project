package commands

import (
	"log"

	"github.com/JuanGQCadavid/now-project/services/spots/internal/core/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
)

type GetFullSpotCommand struct {
	spotId string
}

func (command *GetFullSpotCommand) Run(tr neo4j.Transaction) (interface{}, error) {

	var cypherQ string = `
	MATCH
		(host:Person)-[host_relation:ON_LIVE]->(event:Event {UUID : $spotId})-[location_relation:ON]->(place:Place)
	OPTIONAL MATCH 
		(tags:Topic)-[tagged:TAGGED]->(event)
	RETURN
		event.description as event_desc,
		event.name as event_name,
		event.eventType as event_type,
		event.maximunCapacty as event_max_capacity,
		event.UUID as event_UUID,
		event.emoji as event_emoji,
		place.name as place_name,
		place.lon as place_lon,
		place.mapProviderId as place_provider_id,
		place.lat as place_lat,
		host.phoneNumber as host_phone_number,
		host.name as host_name,
		collect(tags.tag) as tag_tags,
		collect(tagged.isPrincipal) as tag_principals
	`
	cyperParams := map[string]interface{}{"spotId": command.spotId}

	result, err := tr.Run(cypherQ, cyperParams)

	if err != nil {
		println("Error at running!", err)
		return &domain.Spot{}, err
	}
	var spot domain.Spot = domain.Spot{}
	for result.Next() {
		record := result.Record()
		spot = command.getSpotDataFromResult(record)
		log.Printf("%+v", spot)
	}
	return &spot, nil

}

func (command *GetFullSpotCommand) getSpotDataFromResult(record *db.Record) domain.Spot {
	// Event
	event_desc, _ := record.Get("event_desc")
	event_name, _ := record.Get("event_name")
	event_type, _ := record.Get("event_type")
	event_max_capacity, _ := record.Get("event_max_capacity")
	event_UUID, _ := record.Get("event_UUID")
	event_emoji, _ := record.Get("event_emoji")

	// Place
	place_name, _ := record.Get("place_name")
	place_lon, _ := record.Get("place_lon")
	place_provider_id, _ := record.Get("place_provider_id")
	place_lat, _ := record.Get("place_lat")

	// Host
	host_phone_number, _ := record.Get("host_phone_number")
	host_name, _ := record.Get("host_name")

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
			Name:           event_name.(string),
			Description:    event_desc.(string),
			UUID:           event_UUID.(string),
			MaximunCapacty: event_max_capacity.(int64),
			EventType:      event_type.(string),
			Emoji:          event_emoji.(string),
		},
		HostInfo: domain.Person{
			Name:        host_name.(string),
			PhoneNumber: host_phone_number.(string),
		},
		PlaceInfo: domain.Place{
			Name:          place_name.(string),
			Lat:           place_lat.(float64),
			Lon:           place_lon.(float64),
			MapProviderId: place_provider_id.(string),
		},
		TopicsInfo: domain.Topic{
			PrincipalTopic:  primary_tag,
			SecondaryTopics: secondary_tag,
		},
	}

}

func NewGetFullCommand(spotId string) *GetFullSpotCommand {
	return &GetFullSpotCommand{
		spotId: spotId,
	}
}
