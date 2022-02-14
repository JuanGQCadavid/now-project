package locationrepositories

import (
	"fmt"
	"math"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/models"
)

type locationRepository struct {
	connector *MysqlConnector
}

func NewLocationRepo() *locationRepository {

	return &locationRepository{
		connector: NewConectorFromEnv(),
	}
}

func (repo *locationRepository) FetchSpotsIdsByArea(pointA models.LatLng, pointB models.LatLng) (models.Locations, error) {
	db, err := repo.connector.CreateSession()
	defer db.Close()

	if err != nil {
		panic(err)
	}

	pointALatFloat := float64(pointA.Lat)
	pointBLatFloat := float64(pointB.Lat)

	pointALngFloat := float64(pointA.Lng)
	pointBLngFloat := float64(pointB.Lng)

	latLeftLimit, latRigthLimit := math.Min(pointALatFloat, pointBLatFloat), math.Max(pointALatFloat, pointBLatFloat)
	lntLeftLimit, lntRigthLimit := math.Min(pointALngFloat, pointBLngFloat), math.Max(pointALngFloat, pointBLngFloat)

	query := fmt.Sprintf(`
	SELECT 
		spotId, lat, lng 
	FROM 
		locations
	WHERE
		%f <= lat AND lat <= %f
		AND
		%f <= lng AND lng <= %f`, latLeftLimit, latRigthLimit, lntLeftLimit, lntRigthLimit)

	println(query)

	result, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	var spotResult []models.Spot = []models.Spot{}

	for result.Next() {

		var spotLocation SpotLocation
		err = result.Scan(&spotLocation.SpotId, &spotLocation.Lat, &spotLocation.Lng)
		spotResult = append(spotResult, models.Spot{
			Id: spotLocation.SpotId,
			LatLng: models.LatLng{
				Lat: spotLocation.Lat,
				Lng: spotLocation.Lng,
			},
		})
	}

	return models.Locations{
		Places: spotResult,
	}, nil
}
