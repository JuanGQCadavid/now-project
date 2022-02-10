package filtersrv

import (
	"time"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/models"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/ports"
)

type service struct {
	locationRepository ports.LocationRepository
	spotService        ports.SpotService
	maximunTimeWindow  time.Duration
}

func New(locationRepository ports.LocationRepository, spotService ports.SpotService) *service {
	return &service{
		locationRepository: locationRepository,
		maximunTimeWindow:  3 * time.Hour,
		spotService:        spotService,
	}
}

func (srv *service) generatePoints(centralPoint models.LatLng, radious float32) (models.LatLng, models.LatLng) {
	var pointA, pointB models.LatLng

	pointA = models.LatLng{
		Lat: centralPoint.Lat - radious,
		Lng: centralPoint.Lng + radious,
	}

	pointB = models.LatLng{
		Lat: centralPoint.Lat + radious,
		Lng: centralPoint.Lng - radious,
	}

	return pointA, pointB

}

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
// TODO -> should we add the city parameter ?
func (srv *service) FilterByProximity(centralPointLat float32, centralPointLng float32, radious float32) models.Locations {
	//Procedure:
	//	1. Create pointes A and B
	// 	2. Fetch the spotsIds from LocationRepository
	//	4. Remove all spots that are not in the 3 time window
	//	5. Call Spots Service in order to get the spots info
	//
	// Return:
	//	The spots info fetched by spot service but in short format

	// 1. Create pointes A and B
	var pointA, pointB models.LatLng = srv.generatePoints(
		models.LatLng{
			Lat: centralPointLat,
			Lng: centralPointLng,
		},
		radious)

	// 2. Fetch the spotsIds from LocationRepository

	locations, err := srv.locationRepository.FetchSpotsIdsByArea(pointA, pointB)

	if err != nil {
		// TODO -> Do something when it fails here.
	}

	// 4. Remove all spots that are not in the 3 time window

	// TODO -> This seems that is not eficient as it needs to resize when the capacity has been reached.
	// TODO -> Does golang has garbage collector ? Or should I remove it manually ?
	var placesToReturn []string

	for _, spot := range locations.Places {
		println("------------------")
		startTime, err := time.Parse(time.RFC3339, spot.StartTime)
		println("Time before parsed", spot.StartTime)
		println("Time Parsed ->  ", startTime.String())

		if err != nil {
			// TODO -> Do something when it fails here.
		}

		nowTime := time.Now()
		println("Time Now ->  ", nowTime.String())

		// if it is > 0 then it was in the pass, if not it is on the future.
		elapsedTime := nowTime.Sub(startTime)
		println("Time elapsedTime ->  ", elapsedTime.String())

		// The it was on the past
		if elapsedTime < 0 && -elapsedTime > srv.maximunTimeWindow {
			println("Boom chacalaca!!!!!!!!!!!!")
			continue
		}

		println("finish")
		placesToReturn = append(placesToReturn, spot.Id)

	}

	// 5. Call Spots Service in order to get the spots info
	spotsInfo, err := srv.spotService.GetSpotsCardsInfo(placesToReturn)

	if err != nil {
		// TODO -> Do something when it fails here.
	}

	return models.Locations{
		Places: spotsInfo,
	}
}
