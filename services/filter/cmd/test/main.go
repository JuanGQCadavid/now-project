package main

import (
	"fmt"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain/session"
	sessionservice "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/sessionService"
)

func main() {

	search := sessionservice.NewSearchSessionDynamoDbService()
	session2, _ := search.CreateSession(session.SpotsReturned)

	fmt.Println(fmt.Sprintf("%+v", session2))

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
