package locationrepositories

import (
	"fmt"
	"log"

	"github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/core/domain"
)

type locationRepository struct {
	connector *MysqlConnector
}

func NewLocationRepo() *locationRepository {

	return &locationRepository{
		connector: NewConectorFromEnv(),
	}
}

func (repo *locationRepository) CrateLocation(date domain.Date) error {
	log.Printf("CrateLocation: Date: %v\n", date)

	db, err := repo.connector.GetSession()
	defer db.Close()

	if err != nil {
		panic(err)
	}

	query := fmt.Sprintf(`
	INSERT INTO
		locations (spotId, lat, lng) 
	VALUES
		(?,?,?)`)

	println(query)

	// TODO -> check the lat and lon, they should no be empty
	result, err := db.Exec(query, date.DateId, date.Lat, date.Lon)

	if err != nil {
		log.Println("An error ocoured!: ", err)
		return err
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
