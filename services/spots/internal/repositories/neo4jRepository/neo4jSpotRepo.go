package neo4jRepository

import (
	"fmt"
	"log"

	"github.com/JuanGQCadavid/now-project/services/spots/internal/core/domain"
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
)

type Neo4jSpotRepo struct {
	neo4jRepoDriver *Neo4jRepoDriver
	spotInfo        string
}

func NewNeo4jSpotRepo() *Neo4jSpotRepo {

	spotInfo := `
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

	neo4jRepoDriver := GetNeo4jRepoDriver()
	return &Neo4jSpotRepo{
		neo4jRepoDriver: neo4jRepoDriver,
		spotInfo:        spotInfo,
	}
}

func (r Neo4jSpotRepo) println(body interface{}) {
	fmt.Fprintf(gin.DefaultWriter, "%#v", body)
	fmt.Fprintln(gin.DefaultWriter, "  -> DONE")
}
func (r Neo4jSpotRepo) Get(id string) (domain.Spot, error) {
	println("Get id -> ", id)

	session := r.neo4jRepoDriver.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	records, err := session.ReadTransaction(func(tr neo4j.Transaction) (interface{}, error) {
		return r.getSpot(tr, id)
	})

	if err != nil {
		return domain.Spot{}, err
	}

	return *records.(*domain.Spot), nil
}

func (r Neo4jSpotRepo) GetSpots(spotIds []string) (domain.MultipleSpots, error) {
	log.Println("Repository: GetSpots", fmt.Sprintf("%+v", spotIds))

	session := r.neo4jRepoDriver.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	records, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		log.Println("Repository: GetSpots, Before calling GetSpots")
		return r.getSpotsTransaction(tx, spotIds)
	})

	if err != nil {
		return domain.MultipleSpots{}, err
	}

	return *records.(*domain.MultipleSpots), nil
}

func (r Neo4jSpotRepo) convertFromArrayToString(data []string) string {

	var finalResult string = ""

	for _, value := range data {
		finalResult = fmt.Sprintf("%s\"%s\",", finalResult, value)
	}

	finalResult = finalResult[0 : len(finalResult)-1]

	log.Println("Result from joining the array -> " + finalResult)

	return finalResult

}

func (r Neo4jSpotRepo) getSpotsTransaction(tr neo4j.Transaction, spotIds []string) (*domain.MultipleSpots, error) {
	//var cypherQuery string = fmt.Sprintf("MATCH %s WHERE %s RETURN %s ",
	//	"(host:Person)-[host_relation:ON_LIVE]->(event:Event)-[location_relation:ON]->(place:Place)",
	//	"event.UUID IN ["+r.convertFromArrayToString(spotIds)+"]",
	//	r.spotInfo,
	//)

	// result, err := tr.Run(cypherQuery, nill)

	var cypherQuery string = fmt.Sprintf("MATCH %s WHERE %s RETURN %s ",
		"(host:Person)-[host_relation:ON_LIVE]->(event:Event)-[location_relation:ON]->(place:Place)",
		"event.UUID IN $spotIds",
		r.spotInfo,
	)

	println(cypherQuery)

	cyperParams := map[string]interface{}{"spotIds": spotIds}

	result, err := tr.Run(cypherQuery, cyperParams)

	var spotsToReturn []domain.Spot = []domain.Spot{}

	if err != nil {
		println("Error at running!", err)
		return &domain.MultipleSpots{}, err
	}

	for result.Next() {
		record := result.Record()
		spot := r.getSpotDataFromResult(record)
		r.println(spot)
		spotsToReturn = append(spotsToReturn, spot)
	}

	return &domain.MultipleSpots{
		Spots: spotsToReturn,
	}, nil
}

func (r Neo4jSpotRepo) getSpotDataFromResult(record *db.Record) domain.Spot {
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

	r.println(record)

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
	}

}

func (r Neo4jSpotRepo) getSpot(tr neo4j.Transaction, spotId string) (*domain.Spot, error) {

	var cypherQuery string = fmt.Sprintf("MATCH %s RETURN %s ",
		"(host:Person)-[host_relation:ON_LIVE]->(event:Event {UUID : $spotId})-[location_relation:ON]->(place:Place)",
		r.spotInfo,
	)

	cyperParams := map[string]interface{}{"spotId": spotId}

	result, err := tr.Run(cypherQuery, cyperParams)

	if err != nil {
		println("Error at running!", err)
		return &domain.Spot{}, err
	}
	var spot domain.Spot = domain.Spot{}
	for result.Next() {

		record := result.Record()
		spot = r.getSpotDataFromResult(record)
		r.println(spot)
	}
	return &spot, nil
}

func (r Neo4jSpotRepo) CreateOnline(spot domain.Spot) error {

	session := r.neo4jRepoDriver.driver.NewSession(neo4j.SessionConfig{})
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return nil, r.createSpot(tx, spot)
	})

	r.println(err)

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
		r.println(error)

		return error
	}
	r.println(result)
	return nil
}

func (r Neo4jSpotRepo) GetSpotByUserId(personId string) (domain.Spot, error) {
	return domain.Spot{}, nil
}
func (r Neo4jSpotRepo) EndSpot(spotId string) error {
	return nil
}
