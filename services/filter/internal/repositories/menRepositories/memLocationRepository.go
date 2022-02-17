package menrepositories

import (
	"math"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/models"
)

type locationRepository struct {
	memory []models.Spot
}

func NewLocationRepo(data []models.Spot) *locationRepository {
	return &locationRepository{
		memory: data,
	}
}

// This function returns all spots that are btw the rectangle that is formed
// btw the to points ( point A and Point B), in this case the data is dummy data
// which one was created at the instanciation of the struct
func (repo *locationRepository) FetchSpotsIdsByArea(pointA domain.LatLng, pointB domain.LatLng) (domain.Locations, error) {
	// Procedure:
	// 1. Iterate over the memory map
	// 	- If the spot is btw the two points then added it to the repsonse
	//		min(Ax,Bx) <= X1 <= max(Ax,Bx) AND min(Ay,By) <= Y1 <= max(Ay,By)
	//	- If it is not then continue with the next one
	// Return
	//	Locations

	var response = []domain.Spot{}
	pointALatFloat := float64(pointA.Lat)
	pointBLatFloat := float64(pointB.Lat)

	pointALngFloat := float64(pointA.Lng)
	pointBLngFloat := float64(pointB.Lng)

	for _, spot := range repo.memory {

		spotLatFloat := float64(spot.LatLng.Lat)
		spotLngFloat := float64(spot.LatLng.Lng)

		if math.Min(pointALatFloat, pointBLatFloat) <= spotLatFloat && spotLatFloat <= math.Max(pointALatFloat, pointBLatFloat) {
			if math.Min(pointALngFloat, pointBLngFloat) <= spotLngFloat && spotLngFloat <= math.Max(pointALngFloat, pointBLngFloat) {
				response = append(response, domain.Spot{
					EventInfo: domain.Event{
						UUID: spot.Id,
					},
					PlaceInfo: domain.Place{
						Lat: float64(spot.LatLng.Lat),
						Lon: float64(spot.LatLng.Lng),
					},
				})
			}
		}
	}
	return domain.Locations{
		Places: response,
	}, nil
}
