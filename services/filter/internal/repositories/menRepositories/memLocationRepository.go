package menrepositories

import "github.com/JuanGQCadavid/now-project/services/filter/internal/core/models"

type locationRepository struct {
	memory map[string]byte
}

func NewLocationRepo() *locationRepository {
	// TODO -> Populate data with ramdon points btw an area
	// Distributed points

	return &locationRepository{
		memory: map[string]byte{},
	}
}

// This function returns all spots that are btw the rectangle that is formed
// btw the to points ( point A and Point B), in this case the data is dummy data
// which one was created at the instanciation of the struct
func (repo *locationRepository) FetchSpotsIdsByArea(pointA models.LatLng, pointB models.LatLng) (models.Locations, error) {
	// Procedure:
	// 1. Iterate over the memory map
	// 	- If the spot is btw the two points then added it to the repsonse
	//		min(Ax,Bx) <= X1 <= max(Ax,Bx) AND min(Ay,By) <= Y1 <= max(Ay,By)
	//	- If it is not then continue with the next one
	// Return
	//	Locations
	return models.Locations{}, nil
}
