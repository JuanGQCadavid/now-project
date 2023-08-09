package locationrepositories

import (
	"fmt"
	"math"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"gorm.io/gorm"
)

type Coordinates struct {
	LatLeftLimit  float64
	LatRigthLimit float64
	LntLeftLimit  float64
	LntRigthLimit float64
}

type locationRepository struct {
	db *gorm.DB
}

func NewLocationRepo() (*locationRepository, error) {
	connector, err := NewConectorFromEnv()

	if err != nil {
		logs.Error.Println("An error ocoured while calling NewConectorFromEnv, err -> ", err.Error())
		return nil, err
	}

	return &locationRepository{
		db: connector,
	}, nil
}

func NewLocationRepoWithDriver(db *gorm.DB) (*locationRepository, error) {
	return &locationRepository{
		db: db,
	}, nil
}

func (repo *locationRepository) FetchSpotsIdsByAreaExcludingSpots(pointA domain.LatLng, pointB domain.LatLng, spotsIdsToExclude []string) (domain.Locations, error) {
	logs.Info.Println("FetchSpotsIdsByAreaExcludingSpots. Params:", fmt.Sprintf("pointA: %+v, pointB: %+v, spotsIdsToExclude: %+v", pointA, pointB, spotsIdsToExclude))

	if len(spotsIdsToExclude) == 0 {
		logs.Info.Println("Spots to exclude are empty, calling default FetchSpotsIdsByArea")
		return repo.FetchSpotsIdsByArea(pointA, pointB)
	}

	var coord Coordinates = repo.generateCoordinates(pointA, pointB)
	var datesLocations []DatesLocation

	result := repo.db.Where(
		"? <= lat AND lat <= ? AND ? <= lon AND lon <= ? AND date_id NOT IN ?",
		coord.LatLeftLimit, coord.LatRigthLimit, coord.LntLeftLimit, coord.LntRigthLimit, spotsIdsToExclude,
	).Find(&datesLocations)

	if result.Error != nil {
		logs.Error.Println("[ERROR] FetchSpotsIdsByAreaExcludingSpots - An error occoured while runnning Query, err: ", result.Error.Error())
		return domain.Locations{}, ports.ErrQueringData
	}

	return repo.queryResultToLocations(datesLocations)
}

func (repo *locationRepository) FetchSpotsIdsByArea(pointA domain.LatLng, pointB domain.LatLng) (domain.Locations, error) {
	logs.Info.Println("FetchSpotsIdsByArea. Params:", fmt.Sprintf("pointA: %+v, pointB: %+v", pointA, pointB))

	var coord Coordinates = repo.generateCoordinates(pointA, pointB)

	var datesLocations []DatesLocation

	// Missing setting the name propperly
	result := repo.db.Where(
		"? <= lat AND lat <= ? AND ? <= lon AND lon <= ? ",
		coord.LatLeftLimit, coord.LatRigthLimit, coord.LntLeftLimit, coord.LntRigthLimit,
	).Find(&datesLocations)

	if result.Error != nil {
		logs.Error.Println("FetchSpotsIdsByArea - An error occoured while runnning Query, err: ", result.Error.Error())
		return domain.Locations{}, ports.ErrQueringData
	}
	return repo.queryResultToLocations(datesLocations)
}

func (repo *locationRepository) generateCoordinates(pointA domain.LatLng, pointB domain.LatLng) Coordinates {
	// TODO -> What if both pontA and pointB are the same ?
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

func (repo *locationRepository) queryResultToLocations(datesLocations []DatesLocation) (domain.Locations, error) {

	spotResult := make([]domain.Spot, len(datesLocations))

	for _, date := range datesLocations {
		spotResult = append(spotResult, domain.Spot{
			// TODO -> Check if this is needed, if not Just use the date
			EventInfo: domain.Event{
				UUID: date.DateID,
			},

			PlaceInfo: domain.Place{
				Lat: date.Lat,
				Lon: date.Lon,
			},
		})
	}
	return domain.Locations{
		Places: spotResult,
	}, nil
}
