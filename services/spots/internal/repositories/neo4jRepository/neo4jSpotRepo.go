package neo4jRepository

import (
	"github.com/JuanGQCadavid/now-project/services/spots/internal/core/domain"
	"github.com/neo4j/neo4j-go-driver/neo4j"
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

func (r Neo4jSpotRepo) Get(id string) (domain.Spot, error) {
	session := r.neo4jRepoDriver.driver.NewSession(neo4j.SessionConfig{})

	return domain.Spot{}, nil
}
func (r Neo4jSpotRepo) CreateOnline(spot domain.Spot) error {
	return nil
}
func (r Neo4jSpotRepo) GetSpotByUserId(personId string) (domain.Spot, error) {
	return domain.Spot{}, nil
}
func (r Neo4jSpotRepo) EndSpot(spotId string) error {
	return nil
}
