package neo4j

import (
	"log"

	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/repositories/neo4j/commands"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Neo4jRepo struct {
	driver neo4j.Driver
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
	var cmd commands.Command = commands.NewFetchSpotWithDatesCommand(spotId)
	records, err := repo.executeReadCommand(cmd)

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

	var cmd commands.Command = commands.NewCreateDateAssociationCommand(spot, domain.ONLINE_SPOT)
	err := repo.executeWriteCommand(cmd)

	if err != nil {
		log.Println("The associate command fail")
		return ports.ErrAssociatingDate
	}

	return nil
}

func (repo *Neo4jRepo) FetchSpotWithStatus(spotId string, status domain.SpotStatus) (domain.OnlineSpot, error) {
	log.Printf("FetchSpotWithStatus: status: %s, spotId:%s \n", status, spotId)
	var cmd commands.Command = commands.NewFetchSpotCommandWithStatus(spotId, status)
	records, err := repo.executeReadCommand(cmd)

	if err != nil {
		log.Println("There is an error after command, errr:", err.Error())
		return domain.OnlineSpot{}, ports.ErrFetchingData
	}

	spots := records.(domain.OnlineSpot)
	log.Printf("spots: %+v", spots)
	return spots, nil
}

func (repo *Neo4jRepo) FetchSpots(spotId string) (domain.OnlineSpot, error) {
	log.Printf("FetchSpots:  spotId:%s \n", spotId)
	var cmd commands.Command = commands.NewFetchSposCommand(spotId)
	records, err := repo.executeReadCommand(cmd)

	if err != nil {
		log.Println("There is an error after command, errr:", err.Error())
		return domain.OnlineSpot{}, ports.ErrFetchingData
	}

	spots := records.(domain.OnlineSpot)
	log.Printf("spots: %+v", spots)
	return spots, nil
}

func (repo *Neo4jRepo) StopDateOnSpot(spotId string, dateId string) error {
	log.Printf("StopDateOnSpot, spotId: %s dateId: %s \n", spotId, dateId)
	return repo.changeDateStatus(spotId, dateId, domain.PAUSED_SPOT)
}

func (repo *Neo4jRepo) ResumeDateOnSpo(spotId string, dateId string) error {
	log.Printf("ResumeDateOnSpo, spotId: %s dateId: %s \n", spotId, dateId)
	return repo.changeDateStatus(spotId, dateId, domain.ONLINE_SPOT)
}

func (repo *Neo4jRepo) FinalizeDateOnSpot(spotId string, dateId string) error {
	log.Printf("FinalizeDateOnSpot, spotId: %s dateId: %s \n", spotId, dateId)
	return repo.changeDateStatus(spotId, dateId, domain.FINALIZED_SPOT)
}

func (repo *Neo4jRepo) changeDateStatus(spotId string, dateId string, status domain.SpotStatus) error {
	var cmd commands.Command = commands.NewChangeAtStatusCommand(spotId, dateId, status)
	err := repo.executeWriteCommand(cmd)

	if err != nil {
		log.Println("changeDateStatus: Error while using command")
		return ports.ErrUpdatingAtStatus
	}

	return nil
}

func (repo *Neo4jRepo) executeWriteCommand(cmd commands.Command) error {
	session := repo.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	_, err := session.WriteTransaction(cmd.Run)

	return err
}

func (repo *Neo4jRepo) executeReadCommand(cmd commands.Command) (interface{}, error) {
	session := repo.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	return session.WriteTransaction(cmd.Run)
}
