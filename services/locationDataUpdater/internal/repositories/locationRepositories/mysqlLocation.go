package locationrepositories

import (
	"log"

	"github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/core/domain"
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
	log.Printf("CrateLocation: Date: %v\n", date)

	result := repo.connector.session.Create(&date)

	if result.Error != nil {
		log.Println("An error ocoured!: ", result.Error)
		return result.Error
	}

	log.Printf("%+v", result)
	return nil

}

func (repo *locationRepository) RemoveLocation(string) error {
	return nil
}

func (repo *locationRepository) UpdateLocationStatus(string, domain.DateStatus) error {
	return nil
}
