package locationrepositories

import (
	"database/sql"
	"fmt"
	"log"
	"math"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"
)

type Coordinates struct {
	LatLeftLimit  float64
	LatRigthLimit float64
	LntLeftLimit  float64
	LntRigthLimit float64
}

type locationRepository struct {
	connector *MysqlConnector
	db        *sql.DB
}

func NewLocationRepo() *locationRepository {
	connector := NewConectorFromEnv()
	db, err := connector.CreateSession()

	if err != nil {
		panic(err)
	}

	return &locationRepository{
		connector: connector,
		db:        db,
	}
}
func (repo *locationRepository) FetchSpotsIdsByAreaExcludingSpots(pointA domain.LatLng, pointB domain.LatLng, spotsIdsToExclude []string) (domain.Locations, error) {
	log.Println("FetchSpotsIdsByAreaExcludingSpots. Params:", fmt.Sprintf("pointA: %+v, pointB: %+v, spotsIdsToExclude: %+v", pointA, pointB, spotsIdsToExclude))
	if len(spotsIdsToExclude) == 0 {
		log.Println("Spots to exclude are empty, calling default FetchSpotsIdsByArea")
		return repo.FetchSpotsIdsByArea(pointA, pointB)
	}

	var inStatement string = ""

	for index, value := range spotsIdsToExclude {
		if index == 0 {
			inStatement = fmt.Sprintf("(")
		}
		inStatement = fmt.Sprintf("%s \"%s\"", inStatement, value)

		if index == (len(spotsIdsToExclude) - 1) {
			inStatement = fmt.Sprintf("%s)", inStatement)
		} else {
			inStatement = fmt.Sprintf("%s,", inStatement)
		}
	}

	log.Println("FetchSpotsIdsByAreaExcludingSpots. In Statement", inStatement)

	var coord Coordinates = repo.generateCoordinates(pointA, pointB)

	query := fmt.Sprintf(`
	SELECT 
		spotId, lat, lng 
	FROM 
		locations
	WHERE
		%f <= lat AND lat <= %f
		AND
		%f <= lng AND lng <= %f
		AND spotId NOT IN %s`, coord.LatLeftLimit, coord.LatRigthLimit, coord.LntLeftLimit, coord.LntRigthLimit, inStatement)

	log.Println("FetchSpotsIdsByAreaExcludingSpots. Query", query)

	result, err := repo.db.Query(query)
	if err != nil {
		panic(err)
	}
	return repo.queryResultToLocations(result)
}

func (repo *locationRepository) FetchSpotsIdsByArea(pointA domain.LatLng, pointB domain.LatLng) (domain.Locations, error) {
	var coord Coordinates = repo.generateCoordinates(pointA, pointB)

	query := fmt.Sprintf(`
	SELECT 
		spotId, lat, lng 
	FROM 
		locations
	WHERE
		%f <= lat AND lat <= %f
		AND
		%f <= lng AND lng <= %f`, coord.LatLeftLimit, coord.LatRigthLimit, coord.LntLeftLimit, coord.LntRigthLimit)

	println(query)
	result, err := repo.db.Query(query)
	if err != nil {
		panic(err)
	}
	return repo.queryResultToLocations(result)
}

func (repo *locationRepository) generateCoordinates(pointA domain.LatLng, pointB domain.LatLng) Coordinates {
	pointALatFloat := float64(pointA.Lat)
	pointBLatFloat := float64(pointB.Lat)

	pointALngFloat := float64(pointA.Lng)
	pointBLngFloat := float64(pointB.Lng)

	latLeftLimit, latRigthLimit := math.Min(pointALatFloat, pointBLatFloat), math.Max(pointALatFloat, pointBLatFloat)
	lntLeftLimit, lntRigthLimit := math.Min(pointALngFloat, pointBLngFloat), math.Max(pointALngFloat, pointBLngFloat)

	return Coordinates{
		LatLeftLimit:  latLeftLimit,
		LatRigthLimit: latRigthLimit,
		LntLeftLimit:  lntLeftLimit,
		LntRigthLimit: lntRigthLimit,
	}
}

func (repo *locationRepository) queryResultToLocations(result *sql.Rows) (domain.Locations, error) {

	spotResult := []domain.Spot{}

	for result.Next() {

		var spotLocation SpotLocation
		err := result.Scan(&spotLocation.SpotId, &spotLocation.Lat, &spotLocation.Lng)

		if err != nil {
			log.Println("There where a problem while trying to scan the query row")
			return domain.Locations{}, err
		}

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
