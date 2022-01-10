package neo4jRepository

import (
	"fmt"

	"github.com/JuanGQCadavid/now-project/services/spots/internal/core/domain"
	"github.com/neo4j/neo4j-go-driver/neo4j"
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

func (r Neo4jSpotRepo) Get(id string) (*domain.Spot, error) {
	session := r.neo4jRepoDriver.driver.NewSession(neo4j.SessionConfig{})

	records, err := session.ReadTransaction(func(tr neo4j.TransactionConfig) (interface{}, error) {
		return r.getSpot(tr, id)
	})

	if err != nil {
		return &domain.Spot{}, err
	}

	fmt.Println(records)
	spot := records.(*domain.Spot)
	fmt.Println(spot)
	return spot, nil
}

func (r Neo4jSpotRepo) getSpot(tr *neo4j.TransactionConfig, spotId string) (*domain.Spot, error) {

	return &domain.Spot{}, nil
}

func (r Neo4jSpotRepo) CreateOnline(spot domain.Spot) error {
	session := r.neo4jRepoDriver.driver.NewSession(neo4j.SessionConfig{})
	records, err := session.WriteTransaction(func(tr neo4j.TransactionConfig) (interface {}, error) {
		return 
	})
	return nil
}

func (r Neo4jSpotRepo) createEvent( tr *neo4j.TransactionConfig, spot domain.Spot)  (*domain.Spot, error) {

}


func (r Neo4jSpotRepo) GetSpotByUserId(personId string) (*domain.Spot, error) {
	return &domain.Spot{}, nil
}
func (r Neo4jSpotRepo) EndSpot(spotId string) error {
	return nil
}
