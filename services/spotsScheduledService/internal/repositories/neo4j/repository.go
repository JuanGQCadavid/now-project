package neo4j

import (
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/domain"
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

	var filteredPatterns []domain.SchedulePattern

	for _, pattern := range scheduleSpot.Patterns {

		actualStatus := domain.ActivateFlag

		switch pattern.State.Status {
		case domain.ACTIVATE:
			actualStatus = domain.ActivateFlag
		case domain.CONCLUDE:
			actualStatus = domain.ConcludeFlag
		case domain.FREEZE:
			actualStatus = domain.FreezeFlag
		}

		if (flags & actualStatus) == actualStatus {
			filteredPatterns = append(filteredPatterns, pattern)
		}
	}

	scheduleSpot.Patterns = filteredPatterns

	return &scheduleSpot, nil
}
func (repo *Neo4jRepository) AssociateSpotWithSchedulePatterns(spotId string, hostId string, schedulesPattern []domain.SchedulePattern) error {

	return repo.executeWriteCommand(commands.NewAppendScheduleCommand(domain.ScheduledSpot{
		SpotInfo: domain.SpotInfo{
			SpotId:  spotId,
			OwnerId: hostId,
		},
		Patterns: schedulesPattern,
	}))
}

func (repo *Neo4jRepository) UpdateScheculeStatus(spotId string, scheduleId string, status domain.State) error {
	return repo.executeWriteCommand(commands.NewUpdateScheduleStatusCommand(
		scheduleId,
		spotId,
		status,
	))
}

func (repo *Neo4jRepository) GetDatesFromSpot(spotId string) ([]domain.Date, error) {
	records, err := repo.executeReadCommand(commands.NewGetDatesFromSpotCommand(
		spotId,
	))

	if err != nil {
		logs.Error.Println("GetDatesFromSpot: command GetDatesFromSpot error: ", err.Error())
		return nil, err
	}

	if records == nil {
		logs.Info.Println("Empty records where found, returning empty dates")
		return make([]domain.Date, 0, 0), nil
	}

	result := records.([]domain.Date)

	return result, nil
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
