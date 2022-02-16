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

func (repo *locationRepository) CrateLocation(spot domain.Spot) error {
	log.Printf("CrateLocation: \n\tUUID: %s,\n\tLat: %f,\n\tLon: %f", spot.EventInfo.UUID, spot.PlaceInfo.Lat, spot.PlaceInfo.Lon)

	db, err := repo.connector.CreateSession()
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
	result, err := db.Query(query, spot.EventInfo.UUID, spot.PlaceInfo.Lat, spot.PlaceInfo.Lon)
	defer result.Close()

	if err != nil {
		log.Println("An error ocoured!: ", err)
		return err
	}

	log.Printf("%+v", result)

	return nil
}
