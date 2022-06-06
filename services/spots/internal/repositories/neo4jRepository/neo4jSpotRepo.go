package neo4jRepository

import (
	"fmt"
	"log"

	"github.com/JuanGQCadavid/now-project/services/spots/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spots/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/spots/internal/repositories/neo4jRepository/commands"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Neo4jSpotRepo struct {
	neo4jRepoDriver *Neo4jRepoDriver
}

func NewNeo4jSpotRepo() *Neo4jSpotRepo {

	neo4jRepoDriver := GetNeo4jRepoDriver()
	return &Neo4jSpotRepo{
		neo4jRepoDriver: neo4jRepoDriver,
	}
}

func (r Neo4jSpotRepo) Get(id string, format ports.OutputFormat) (domain.Spot, error) {
	println("Get id -> ", id)

	var command commands.Command

	switch format {
	case ports.FULL_FORMAT:
		command = commands.NewGetFullCommand(id)
	case ports.SMALL_FORMAT:
		command = commands.NewGetSmallCommand(id)
	default:
		command = commands.NewGetFullCommand(id)
	}

	session := r.neo4jRepoDriver.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	records, err := session.ReadTransaction(func(tr neo4j.Transaction) (interface{}, error) {
		return command.Run(tr)
	})

	if err != nil {
		return domain.Spot{}, err
	}

	return *records.(*domain.Spot), nil
}

func (r Neo4jSpotRepo) GetSpots(spotIds []string, format ports.OutputFormat) (domain.MultipleSpots, error) {
	log.Println("Repository: GetSpots", fmt.Sprintf("%+v", spotIds))

	var command commands.Command

	switch format {
	case ports.FULL_FORMAT:
		command = commands.NewGetFullMultipleSpotsCommand(spotIds)
	case ports.SMALL_FORMAT:
		command = commands.NewGetSmallMultipleSpotsCommand(spotIds)
	default:
		command = commands.NewGetFullMultipleSpotsCommand(spotIds)
	}

	session := r.neo4jRepoDriver.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	records, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return command.Run(tx)
	})

	if err != nil {
		return domain.MultipleSpots{}, err
	}

	return *records.(*domain.MultipleSpots), nil
}

func (r Neo4jSpotRepo) CreateOnline(spot domain.Spot) error {

	session := r.neo4jRepoDriver.driver.NewSession(neo4j.SessionConfig{})
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return nil, r.createSpot(tx, spot)
	})

	log.Println(err)

	return err
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
		MERGE (host)-[:ON_LIVE]->(event)-[:ON]->(place)
	
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

func (r Neo4jSpotRepo) GetSpotByUserId(personId string) (domain.Spot, error) {
	return domain.Spot{}, nil
}
func (r Neo4jSpotRepo) EndSpot(spotId string) error {
	return nil
}

func (r Neo4jSpotRepo) CreateSpotTags(spotId string, principalTag domain.Optional, secondaryTags []string) error {
	log.Println("Repository: CreateSpotTags", "\nspotId: ", spotId, "\nprincipalTag: ", fmt.Sprintf("%+v", principalTag), "\nsecondaryTags: ", fmt.Sprintf("%+v", secondaryTags))

	if !principalTag.IsPresent() && (secondaryTags == nil || len(secondaryTags) == 0) {
		log.Println("Avoiding process as both Principal and secondary tags are empty.")
		return nil
	}

	session := r.neo4jRepoDriver.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return nil, r.createSpotTags(tx, spotId, principalTag, secondaryTags)
	})

	return err
}

func (r Neo4jSpotRepo) createSpotTags(tr neo4j.Transaction, spotId string, principalTag domain.Optional, secondaryTags []string) error {

	var cypherBaseCommand string = "MATCH (spot:Event {UUID: $event_uuid})"

	var params map[string]interface{} = make(map[string]interface{})
	params["event_uuid"] = spotId

	if principalTag.IsPresent() {
		cyperPrimaryTagCommand := "MERGE (primaryTag:Topic {tag: $primaryTag })\nMERGE (primaryTag)-[:TAGGED {isPrincipal:true}]->(spot)"

		cypherBaseCommand = fmt.Sprintf("%s\n%s", cypherBaseCommand, cyperPrimaryTagCommand)
		params["primaryTag"] = principalTag.GetValue()
	}

	if secondaryTags != nil {
		for index, tag := range secondaryTags {
			tagKey := fmt.Sprintf("secondaryTag%d", index)
			cyperSecondaryTagCommandCreation := fmt.Sprintf("MERGE (%s:Topic {tag: $%s })", tagKey, tagKey)
			cyperSecondaryTagCommandLinking := fmt.Sprintf("MERGE (%s)-[:TAGGED {isPrincipal:false}]->(spot)", tagKey)
			cypherBaseCommand = fmt.Sprintf("%s\n%s\n%s", cypherBaseCommand, cyperSecondaryTagCommandCreation, cyperSecondaryTagCommandLinking)
			params[tagKey] = tag
		}
	}

	log.Println(cypherBaseCommand)

	return nil
}
