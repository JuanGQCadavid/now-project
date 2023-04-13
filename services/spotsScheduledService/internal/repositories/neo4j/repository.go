package neo4j

import (
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/logs"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/repositories/neo4j/commands"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Neo4jRepository struct {
	driver neo4j.Driver
}

func NewNeo4jRepoWithDriver(driver neo4j.Driver) *Neo4jRepository {
	return &Neo4jRepository{
		driver: driver,
	}
}

func (repo *Neo4jRepository) GetScheduleSpot(spotId string, flags domain.ScheduleStateFlags) (*domain.ScheduledSpot, error) {

	var command commands.Command = commands.NewGetSchedulesCommand(spotId)
	records, err := repo.executeReadCommand(command)

	if err != nil {
		logs.Error.Println("GetScheduleSpot: command getsFail erro: ", err.Error())
		return nil, err
	}

	var scheduleSpot domain.ScheduledSpot = records.(domain.ScheduledSpot)
	return &scheduleSpot, nil
}
func (repo *Neo4jRepository) AssociateSpotWithSchedulePatterns(spotId string, hostId string, schedulesPattern *[]domain.SchedulePattern) error {
	return nil
}

func (repo *Neo4jRepository) UpdateScheculeStatus(spotId string, scheduleId string, status domain.State) error {
	return nil
}

func (repo *Neo4jRepository) executeWriteCommand(cmd commands.Command) error {
	session := repo.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	_, err := session.WriteTransaction(cmd.Run)

	return err
}

func (repo *Neo4jRepository) executeReadCommand(cmd commands.Command) (interface{}, error) {
	session := repo.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	return session.WriteTransaction(cmd.Run)
}
