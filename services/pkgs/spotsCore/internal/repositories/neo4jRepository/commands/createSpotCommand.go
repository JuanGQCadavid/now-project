package commands

import (
	"log"

	"github.com/JuanGQCadavid/now-project/services/spotsCore/internal/core/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type CreateSpotCommand struct {
	Spot *domain.Spot
}

func NewCreateSpotCommand(spot *domain.Spot) *CreateSpotCommand {
	return &CreateSpotCommand{
		Spot: spot,
	}
}

func (cmd *CreateSpotCommand) Run(tr neo4j.Transaction) (interface{}, error) {

}

func (r Neo4jSpotRepo) createSpot(tr neo4j.Transaction, spot domain.Spot) error {

	var cypher string = `
		MERGE (event:Event {UUID: $event_uuid })
		ON CREATE
			SET event.description = $event_desc
			SET event.maximunCapacty = $event_max_capacity
			SET event.eventType = $event_type
			SET event.name = $event_name
			SET event.emoji = $event_emoji
		MERGE (place:Place {mapProviderId: $place_provider_id})
		ON CREATE
			SET place.lat = toFloat($place_lat)
			SET place.lon = toFloat($place_lon)
			SET place.name = $place_name
		MERGE (host:Person {phoneNumber:$host_phone_number})
		ON CREATE 
			SET host.name = $host_name
		MERGE (host)<-[:CREATED_BY]-(event)-[:ON]->(place)
	
	`
	result, error := tr.Run(cypher, map[string]interface{}{
		"event_uuid":         spot.EventInfo.UUID,
		"event_desc":         spot.EventInfo.Description,
		"event_max_capacity": spot.EventInfo.MaximunCapacty,
		"event_type":         spot.EventInfo.EventType,
		"event_name":         spot.EventInfo.Name,
		"event_emoji":        spot.EventInfo.Emoji,
		"place_provider_id":  spot.PlaceInfo.MapProviderId,
		"place_lat":          spot.PlaceInfo.Lat,
		"place_lon":          spot.PlaceInfo.Lon,
		"place_name":         spot.PlaceInfo.Name,
		"host_phone_number":  spot.HostInfo.PhoneNumber,
		"host_name":          spot.HostInfo.Name,
	})

	if error != nil {
		log.Println(error)

		return error
	}
	log.Println(result)
	return nil
}
