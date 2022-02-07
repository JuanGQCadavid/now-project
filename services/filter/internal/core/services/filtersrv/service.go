package filtersrv

import (
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/models"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/ports"
)

type service struct {
	locationRepository ports.LocationRepository
}

func New(locationRepository ports.LocationRepository) *service {
	return &service{
		locationRepository: locationRepository,
	}
}

// TODO -> should we add the city parameter ?
func (srv *service) FilterByProximity(centralPointLat float32, centralPointLng float32, radious float32) models.Locations {
	//Procedure:
	//	1. Create pointes A and B
	// 	2. Fetch the spotsIds from LocationRepository
	//	4. Remove all spots that are not in the 3 time window
	//	5. Call Spots Service in order to get the spots info
	//
	//
	// Return:
	//	The spots info fetched by spot service but in short format
	//{
	// 	"places": [
	// 		{
	// 			"id" : <String>,
	// 			"type": [ "ONLINE" |  "EVENT" | "UPCOMMING" ],
	// 			"emoji" : <String>, // The emoji code
	// 			"startsIn": <DateTime> // Only if the event is a upcomming one.
	// 		}
	// 		...
	// 		...
	// 		...
	// 		{
	// 			"id" : <String>,
	// 			"type": [ "ONLINE" |  "EVENT" | "UPCOMMING" ],
	// 			"emoji" : <String>, // The emoji code
	// 			"startsIn": <DateTime> // Only if the event is a upcomming one.
	// 		}
	// 	]
	// }
	return models.Locations{}
}
