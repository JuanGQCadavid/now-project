package neo4jRepository

import (
	"fmt"

	"github.com/JuanGQCadavid/now-project/services/spots/internal/core/domain"
	"github.com/gin-gonic/gin"
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

func (r Neo4jSpotRepo) println(body interface{}) {
	fmt.Fprintf(gin.DefaultWriter, "%#v", body)
}
func (r Neo4jSpotRepo) Get(id string) (domain.Spot, error) {
	println("Get id -> ", id)

	session := r.neo4jRepoDriver.driver.NewSession(neo4j.SessionConfig{})

	records, err := session.ReadTransaction(func(tr neo4j.Transaction) (interface{}, error) {
		return r.getSpot(tr, id)
	})

	if err != nil {
		return domain.Spot{}, err
	}

	println(records)
	spot := records.(*domain.Spot)
	println(spot)
	return domain.Spot{}, nil
}

func (r Neo4jSpotRepo) getSpot(tr neo4j.Transaction, spotId string) (*domain.Spot, error) {

	var cypherQuery string = "MATCH (e:Event {UUID : $spotId})-[r]-(s) RETURN e,r,s"
	cyperParams := map[string]interface{}{"spotId": spotId}

	result, err := tr.Run(cypherQuery, cyperParams)

	if err != nil {
		println("Error at running!", err)
		return &domain.Spot{}, err
	}

	resultRecords, err := result.Collect()

	if err != nil {
		println("Error at collecting !", err)
		return &domain.Spot{}, err
	}

	var spot domain.Spot

	for _, record := range resultRecords {

		// print("Values -> ")
		// println(record.Values...)

		print("Values other way -> ")
		r.println(record.Values)

		print("Keys -> ")
		r.println(record.Keys)

		// DUDE We need to unmarshal that !!

	}

	return &domain.Spot{}, nil
}

func (r Neo4jSpotRepo) CreateOnline(spot domain.Spot) error {
	// session := r.neo4jRepoDriver.driver.NewSession(neo4j.SessionConfig{})
	// records, err := session.WriteTransaction(func(tr neo4j.Transaction) (interface{}, error) {
	// 	return
	// })
	return nil
}

func (r Neo4jSpotRepo) createEvent(tr *neo4j.TransactionConfig, spot domain.Spot) (*domain.Spot, error) {
	return &domain.Spot{}, nil
}

func (r Neo4jSpotRepo) GetSpotByUserId(personId string) (domain.Spot, error) {
	return domain.Spot{}, nil
}
func (r Neo4jSpotRepo) EndSpot(spotId string) error {
	return nil
}
