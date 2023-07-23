package neo4j

import (
	"github.com/JuanGQCadavid/now-project/services/confirmationService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/confirmationService/internal/repositories/neo4j/commands"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
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

func (repo *Neo4jRepository) FetchDate(dateId string) (*domain.Date, error) {
	records, err := repo.executeReadCommand(commands.NewFetchDateCommand(
		dateId,
	))

	if err != nil {
		logs.Error.Println("executeReadCommand Fail")
		return nil, err
	}

	if records != nil {
		return records.(*domain.Date), nil
	}

	logs.Warning.Println("Empty record were returned by executeReadCommand")
	return nil, nil
}

func (repo *Neo4jRepository) UpdateDateOnConfirmed(dateId string, confirmed bool) error {
	logs.Info.Printf("UpdateDateOnConfirmed -> dateId: %s confirmed: %v \n", dateId, confirmed)
	err := repo.executeWriteCommand(commands.NewUpdateDateStatusCommand(dateId, confirmed))

	if err != nil {
		logs.Error.Println("executeWriteCommand Fail")
		return err
	}

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
