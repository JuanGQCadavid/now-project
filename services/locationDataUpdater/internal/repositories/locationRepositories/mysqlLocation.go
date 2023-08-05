package locationrepositories

import (
	"github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
)

type locationRepository struct {
	connector *MysqlConnector
}

func NewLocationRepoFromEnv() (*locationRepository, error) {

	connector, err := NewConectorFromEnv()

	if err != nil {
		return nil, err
	}

	return &locationRepository{
		connector: connector,
	}, nil
}

func NewLocationRepo(connector *MysqlConnector) (*locationRepository, error) {

	return &locationRepository{
		connector: connector,
	}, nil
}

func (repo *locationRepository) CrateLocation(date domain.DatesLocation) error {
	logs.Info.Printf("CrateLocation: Date: %v\n", date)

	result := repo.connector.session.Create(&date)

	if result.Error != nil {
		logs.Error.Println("An error ocoured!: ", result.Error)
		return result.Error
	}

	logs.Info.Printf("%+v", result)
	return nil

}

func (repo *locationRepository) RemoveLocation(string) error {
	return nil
}

func (repo *locationRepository) UpdateLocationStatus(dateID string, state domain.DateState) error {

	logs.Info.Printf("UpdateLocationStatus: dateID: %v, status: %v \n", dateID, state)

	date := domain.DatesLocation{}

	result := repo.connector.session.First(&date, &dateID)

	if result.Error != nil {
		logs.Error.Println("Error while Fetching date: ", result.Error)
		return result.Error
	}

	date.State = domain.States{
		StateID: state,
	}

	result = repo.connector.session.Save(&date)

	if result.Error != nil {
		logs.Error.Println("Error while Saving new state date: ", result.Error)
		return result.Error
	}

	logs.Info.Printf("%+v", result)
	return nil
}
