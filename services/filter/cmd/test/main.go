package main

import (
	"fmt"
	"log"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain/session"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/ports"
	sessionservice "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/sessionService"
)

func main() {

	// var time1 = time.Now().Format(time.RFC3339Nano)

	// var time2 = time.Now().Format(time.RFC3339Nano)

	// log.Println(time1, time2)

	// log.Println(time1 == time2)

	var search ports.SearchSessionService = sessionservice.NewSearchSessionDynamoDbService()
	session2, _ := search.CreateSession(session.SpotsReturned)
	log.Println(fmt.Sprintf("%+v", session2))

	spotsIds := []string{
		"spotId1", "SpotId2", "SpotId3", "spotId4", "SpotId5", "SpotId6", "spotId7", "SpotId8", "SpotId9", "spotId10", "SpotId11", "SpotId12", "spotId13", "SpotId14", "SpotId15",
	}

	err := search.AddSpotsToSession(session2.SessionId, session2.SessionType, spotsIds)
	log.Println(fmt.Sprintf("%+v", err))

	// time.Sleep(time.Duration(2) * time.Second)

	// err = search.AddSpotsToSession(session2.SessionId, session2.SessionType, spotsIds)
	// log.Println(fmt.Sprintf("%+v", err))

	// log.Println(session2.SessionId)

	search.GetSessionData(session2.SessionId, session.SpotsReturned)

	// for {
	// 	err := search.AddSpotsToSession(session2.SessionId, session2.SessionType, spotsIds)
	// 	time.Sleep(time.Duration(2) * time.Second)
	// 	log.Println(fmt.Sprintf("%+v", err))
	// }

	// location := locationrepositories.NewLocationRepo()
	// spot := spotservicelambda.NewSpotServiceLambda()

	// srv := filtersrv.New(location, spot)

	// locations := srv.FilterByProximity(75.15, 32.59, 0.5)

	// str, _ := json.Marshal(locations)

	// log.Println(string(str))

	//log.Printf("%+v", locations)

	/*
		a := domain.LatLng{
			Lat: 78.00,
			Lng: 39.00,
		}

		b := domain.LatLng{
			Lat: 72.00,
			Lng: 30.00,
		}
		places, err := location.FetchSpotsIdsByArea(a, b)

		if err != nil {
			panic(err)
		}

		fmt.Printf("%+v", places)
	*/
}
