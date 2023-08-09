package filtersrv

import (
	"fmt"
	"time"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain/session"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
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

func (srv *service) generatePoints(centralPoint domain.LatLng, radious float64) (domain.LatLng, domain.LatLng) {
	var pointA, pointB domain.LatLng

	pointA = domain.LatLng{
		Lat: centralPoint.Lat - radious,
		Lng: centralPoint.Lng + radious,
	}

	pointB = domain.LatLng{
		Lat: centralPoint.Lat + radious,
		Lng: centralPoint.Lng - radious,
	}

	return pointA, pointB

}

//	{
//		"places": [
//			{
//				"id" : <String>,
//				"type": [ "ONLINE" |  "EVENT" | "UPCOMMING" ],
//				"emoji" : <String>, // The emoji code
//				"startsIn": <DateTime> // Only if the event is a upcomming one.
//			}
//			...
//			...
//			...
//			{
//				"id" : <String>,
//				"type": [ "ONLINE" |  "EVENT" | "UPCOMMING" ],
//				"emoji" : <String>, // The emoji code
//				"startsIn": <DateTime> // Only if the event is a upcomming one.
//			}
//		]
//	}
//
// TODO -> should we add the city parameter ?
func (srv *service) FilterByProximity(centralPointLat float64, centralPointLng float64, radious float64, sessionData session.SearchSessionData, format ports.OutputFormat) (domain.Locations, error) {
	//Procedure:
	//	1. Create pointes A and B
	// 	2. Fetch the spotsIds from LocationRepository
	//	4. Remove all spots that are not in the 3 time window
	//	5. Call Spots Service in order to get the spots info
	//
	// Return:
	//	The spots info fetched by spot service but in short format

	// 1. Create pointes A and B
	logs.Info.Println(fmt.Sprintf("FilterByProximity - centralPointLat: %f, centralPointLng: %f, radious: %f, sessionData: %+v", centralPointLat, centralPointLng, radious, sessionData))

	var pointA, pointB domain.LatLng = srv.generatePoints(
		domain.LatLng{
			Lat: centralPointLat,
			Lng: centralPointLng,
		},
		radious)

	var locations domain.Locations
	var err error
	// 2. Fetch the spotsIds from LocationRepository
	if sessionData.SessionConfiguration.SessionType != session.Empty && len(sessionData.Spots) > 0 {
		dates := srv.unfoldDatesIds(sessionData.Spots)
		locations, err = srv.locationRepository.FetchSpotsIdsByAreaExcludingSpots(pointA, pointB, dates)
	} else {
		locations, err = srv.locationRepository.FetchSpotsIdsByArea(pointA, pointB)
	}

	if err != nil {
		// TODO -> Do something when it fails here.
	}

	// 4. Remove all spots that are not in the 3 time window
	placesToReturn := srv.filterByTime(locations)

	// 5. Call Spots Service in order to get the spots info
	spotsInfo, err := srv.spotService.GetSpotsCardsInfo(placesToReturn, format)

	if err != nil {
		switch err {
		case ports.ErrBodyRequestUnmarshal:
			logs.Error.Println("The process halts wile performign the body request unmarshal, returning empty spots")
			break
		case ports.ErrBodyResponseUnmarshal:
			logs.Error.Println("The process halts wile performign the body response unmarshal, returning empty spots")
			break
		case ports.ErrBodyResponseReadFail:
			logs.Error.Println("The process halts wile performign the body data read, returning empty spots")
			break
		case ports.ErrSendingRequest:
			logs.Error.Println("The process halts wile performign the spot service call, returning empty spots")
			break
		default:
			logs.Error.Println("err -> ", err)
			break
		}

		return domain.Locations{}, ports.ErrSpotServiceFail
	}

	return domain.Locations{
		Places: spotsInfo,
	}, nil
}

func (srv *service) unfoldDatesIds(sessionDataSpots map[string][]string) []string {
	logs.Info.Println("unfoldDatesIds. sessionDataSpots:", fmt.Sprintf("%+v", sessionDataSpots))

	spots := []string{}
	for key, value := range sessionDataSpots {
		logs.Info.Println("Appending sesison data from timestamp: ", key)
		spots = append(spots, value...)
	}

	return spots
}

// HERE What is this doing ?
func (srv *service) filterByTime(locations domain.Locations) []string {
	// TODO -> This seems that is not eficient as it needs to resize when the capacity has been reached.
	var placesToReturn []string

	for _, spot := range locations.Places {
		startTimeNow := time.Now().Format(time.RFC3339) // Fetch it from Spot!  spot.StartTime

		startTime, err := time.Parse(time.RFC3339, startTimeNow)
		logs.Info.Println("Time before parsed ->", startTimeNow)
		logs.Info.Println("Time Parsed ->  ", startTime.String())

		if err != nil {
			// TODO -> Do something when it fails here.
		}

		nowTime := time.Now()
		logs.Info.Println("Time Now ->  ", nowTime.String())

		// if it is > 0 then it was in the pass, if not it is on the future.
		elapsedTime := nowTime.Sub(startTime)
		logs.Info.Println("Time elapsedTime ->  ", elapsedTime.String())

		// The it was on the past
		if elapsedTime < 0 && -elapsedTime > srv.maximunTimeWindow {
			logs.Info.Println("Boom chacalaca!!!!!!!!!!!!")
			continue
		}

		logs.Info.Println("finish")
		placesToReturn = append(placesToReturn, spot.EventInfo.UUID)
	}

	return placesToReturn

}
