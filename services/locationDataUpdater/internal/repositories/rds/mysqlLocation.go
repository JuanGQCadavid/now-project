package rds

import (
	"github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
)

type RDSRepository struct {
	connector *MysqlConnector
}

func NewRDSRepoFromEnv() (*RDSRepository, error) {

	connector, err := NewConectorFromEnv()

	if err != nil {
		return nil, err
	}

	return &RDSRepository{
		connector: connector,
	}, nil
}

func NewRDSRepo(connector *MysqlConnector) (*RDSRepository, error) {

	return &RDSRepository{
		connector: connector,
	}, nil
}

func (repo *RDSRepository) CrateLocation(date domain.DatesLocation) error {
	logs.Info.Printf("CrateLocation: Date: %v\n", date)

	result := repo.connector.session.Create(&date)

	if result.Error != nil {
		logs.Error.Println("An error ocoured!: ", result.Error)
		return result.Error
	}

	logs.Info.Printf("%+v", result)
	return nil

}

func (repo *RDSRepository) RemoveLocation(dateID string) error {
	logs.Info.Printf("RemoveLocation: dateID: %v\n", dateID)

	result := repo.connector.session.Unscoped().Delete(&domain.DatesLocation{}, &dateID)

	if result.Error != nil {
		logs.Error.Println("Error while deleting date: ", result.Error)
		return result.Error
	}

	logs.Info.Printf("%+v", result)
	return nil
}

func (repo *RDSRepository) UpdateLocationStatus(dateID string, state domain.DateState) error {
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

func (repo *RDSRepository) UpdateLocationType(dateID string, dateType domain.DateType) error {
	logs.Info.Printf("UpdateLocationStatus: dateID: %v, dateType: %v \n", dateID, dateType)

	date := domain.DatesLocation{}

	result := repo.connector.session.First(&date, &dateID)

	if result.Error != nil {
		logs.Error.Println("Error while Fetching date: ", result.Error)
		return result.Error
	}

	date.Type = domain.Types{
		TypeID: dateType,
	}

	result = repo.connector.session.Save(&date)

	if result.Error != nil {
		logs.Error.Println("Error while Saving new state date: ", result.Error)
		return result.Error
	}

	logs.Info.Printf("%+v", result)
	return nil
}
