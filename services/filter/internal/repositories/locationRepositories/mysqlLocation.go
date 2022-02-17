package locationrepositories

import (
	"fmt"
	"math"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"
)

type locationRepository struct {
	connector *MysqlConnector
}

func NewLocationRepo() *locationRepository {

	return &locationRepository{
		connector: NewConectorFromEnv(),
	}
}

func (repo *locationRepository) FetchSpotsIdsByArea(pointA domain.LatLng, pointB domain.LatLng) (domain.Locations, error) {
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

	spotResult := []domain.Spot{}

	for result.Next() {

		var spotLocation SpotLocation
		err = result.Scan(&spotLocation.SpotId, &spotLocation.Lat, &spotLocation.Lng)

		spotResult = append(spotResult, domain.Spot{
			EventInfo: domain.Event{
				UUID: spotLocation.SpotId,
			},
			PlaceInfo: domain.Place{
				Lat: spotLocation.Lat,
				Lon: spotLocation.Lng,
			},
		})
	}

	return domain.Locations{
		Places: spotResult,
	}, nil
}
