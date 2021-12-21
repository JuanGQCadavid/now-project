package neo4jRepository

import "github.com/JuanGQCadavid/now-project/services/spots/internal/core/domain"

type Neo4jSpotRepo struct {
}

func (r Neo4jSpotRepo) Get(id string) (domain.Spot, error) {
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
