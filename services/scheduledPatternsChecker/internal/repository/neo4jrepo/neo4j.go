package neo4jrepo

import (
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/repository/neo4jrepo/commands"
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

// 1. Bring all schedule patterns that are in the repository that:
//   - Pattern is active
//   - Spot Id is not eliminated
//   - Current date is btw Start and end date
func (repo *Neo4jRepository) FetchActiveSchedulePatterns() ([]domain.Spot, error) {
	cmd := commands.NewFetchAllSchedulePatternsCommand()

	result, err := repo.executeReadCommand(cmd)

	if err != nil {
		logs.Error.Println("We got errors!", err.Error())
	}

	resultAsMap := result.(map[string]domain.Spot)

	if len(resultAsMap) == 0 {
		return nil, nil
	}

	spots := make([]domain.Spot, len(resultAsMap))

	index := 0
	for _, spot := range resultAsMap {
		spots[index] = spot
		index++
	}

	return spots, nil
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
