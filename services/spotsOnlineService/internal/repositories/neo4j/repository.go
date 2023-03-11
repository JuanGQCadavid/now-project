package neo4j

import (
	"log"

	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/repositories/neo4j/commands"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Neo4jRepo struct {
	neo4jRepoDriver *Neo4jRepoDriver
	driver          neo4j.Driver
}

func NewNeo4jRepo() *Neo4jRepo {

	neo4jRepoDriver := GetNeo4jRepoDriver()
	return &Neo4jRepo{
		driver: neo4jRepoDriver.driver,
	}
}

func NewNeo4jRepoWithDriver(driver neo4j.Driver) *Neo4jRepo {
	return &Neo4jRepo{
		driver: driver,
	}
}

func (repo *Neo4jRepo) FetchOnlineSpot(spotId string) (domain.OnlineSpot, error) {
	log.Printf("FetchOnlineSpot %s \n", spotId)
	var command commands.Command = commands.NewFetchSpotCommand(spotId)

	session := repo.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	records, err := session.ReadTransaction(func(tr neo4j.Transaction) (interface{}, error) {
		return command.Run(tr)
	})

	if err != nil {
		log.Println("There is an error after command, errr:", err.Error())
		return domain.OnlineSpot{}, ports.ErrFetchingData
	}

	onlineSpot := records.(domain.OnlineSpot)
	log.Printf("%+v", onlineSpot)
	return onlineSpot, nil
}

func (repo *Neo4jRepo) AssociateDateWithSpot(spot domain.OnlineSpot) error {
	log.Printf("AssociateDateWithSpot,  online spot: %+v \n", spot)

	var cmd commands.Command = commands.NewCreateDateAssociationCommand(spot)

	session := repo.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	_, err := session.WriteTransaction(cmd.Run)

	if err != nil {
		log.Println("The associate command fail")
		return ports.ErrAssociatingDate
	}

	return nil
}
