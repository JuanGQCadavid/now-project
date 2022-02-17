package main

import (
	"encoding/json"
	"log"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/services/filtersrv"
	locationrepositories "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/locationRepositories"
	spotservicelambda "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/spotServiceLambda"
)

func main() {

	location := locationrepositories.NewLocationRepo()
	spot := spotservicelambda.NewSpotServiceLambda()

	srv := filtersrv.New(location, spot)

	locations := srv.FilterByProximity(75.15, 32.59, 0.5)

	str, _ := json.Marshal(locations)

	log.Println(string(str))

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
