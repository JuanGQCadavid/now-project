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

const (
	DATES_STATUS = "SCHEDULED"
)

func NewNeo4jRepoWithDriver(driver neo4j.Driver) *Neo4jRepository {
	return &Neo4jRepository{
		driver: driver,
	}
}

func (repo *Neo4jRepository) ConditionalDatesCreation(spot domain.Spot) error {
	logs.Info.Printf("ConditionalDatesCreation: Spot id %s, total ScheduePatterns: %d \n", spot.SpotId, len(spot.SchedulePatterns))
	cmd := commands.NewConditionalWritterCommand(spot, DATES_STATUS)
	return repo.executeWriteCommand(cmd)
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

func (repo *Neo4jRepository) UpdateSpotsByBatch(spots []domain.Spot, batchSize int) map[*domain.Spot]error {
	logs.Info.Printf("UpdateSpotsByBatch wiht %d spots, batchsize of %d  \n", len(spots), batchSize)
	errors := make(map[*domain.Spot]error)

	var startPointer int
	var threshold int

	if len(spots) == 0 {
		logs.Info.Println("Empty spots, aborting job")
		return nil
	}

	for {
		threshold = startPointer + batchSize
		var errUpdate error

		if threshold < len(spots) {
			errUpdate = repo.updateSpots(spots[startPointer:threshold])

			if errUpdate != nil {
				repo.saveSpotsOnError(spots[startPointer:threshold], errUpdate, errors)
			}

			startPointer = threshold
		} else if (len(spots)-startPointer) > 0 && (len(spots)-startPointer) <= batchSize {
			errUpdate = repo.updateSpots(spots[startPointer:])

			if errUpdate != nil {
				repo.saveSpotsOnError(spots[startPointer:], errUpdate, errors)
			}

			startPointer = len(spots)
		} else {
			break
		}

	}

	if errors != nil && len(errors) > 0 {
		logs.Error.Println("We found erros while processing the bacth request!")
		for key, err := range errors {
			logs.Error.Printf("%+v - %s \n", key, err.Error())
		}
		return errors
	}

	logs.Info.Println("All spots where updated successfully")
	return nil
}

func (repo *Neo4jRepository) saveSpotsOnError(spots []domain.Spot, err error, errorsMap map[*domain.Spot]error) {
	if err != nil {
		logs.Error.Println("We got an error while processing the bulk operation on bacth, spots id with error: ", err.Error())
		for _, spot := range spots {
			logs.Error.Print(spot.SpotId, ", ")
			errorsMap[&spot] = err
		}
		logs.Error.Println()
	}
}

func (repo *Neo4jRepository) updateSpots(spots []domain.Spot) error {
	logs.Info.Printf("updateSpots wiht %d spots\n", len(spots))
	return repo.executeWriteCommand(commands.NewUpdateSpotsCommand(spots, DATES_STATUS))
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
