package commands

import (
	"log"

	"github.com/JuanGQCadavid/now-project/services/spots/internal/core/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
)

type GetFullMultipleSpotsCommand struct {
	spotIds []string
}

func (command *GetFullMultipleSpotsCommand) Run(tr neo4j.Transaction) (interface{}, error) {

	var cypherQ string = `
	MATCH
		(host:Person)-[host_relation:ON_LIVE]->(event:Event)-[location_relation:ON]->(place:Place)
	WHERE
		event.UUID IN $spotIds
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
		host.name as host_name
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

func (command *GetFullMultipleSpotsCommand) getSpotDataFromResult(record *db.Record) domain.Spot {
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

	log.Printf("%+v", record)

	return domain.Spot{
		EventInfo: domain.Event{
			Name:           getStringFromInterface(event_name),
			Description:    getStringFromInterface(event_desc),
			UUID:           getStringFromInterface(event_UUID),
			MaximunCapacty: getInt64FromInterface(event_max_capacity),
			EventType:      getStringFromInterface(event_type),
			Emoji:          getStringFromInterface(event_emoji),
		},
		HostInfo: domain.Person{
			Name:        getStringFromInterface(host_name),
			PhoneNumber: getStringFromInterface(host_phone_number),
		},
		PlaceInfo: domain.Place{
			Name:          getStringFromInterface(place_name),
			Lat:           getFloat64FromInterface(place_lat),
			Lon:           getFloat64FromInterface(place_lon),
			MapProviderId: getStringFromInterface(place_provider_id),
		},
	}

}

func NewGetFullMultipleSpotsCommand(spotIds []string) *GetFullMultipleSpotsCommand {
	return &GetFullMultipleSpotsCommand{
		spotIds: spotIds,
	}
}
