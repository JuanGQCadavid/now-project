package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/services/filtersrv"
	fakedata "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/fakeData"
	locationrepositories "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/locationRepositories"
	menrepositories "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/menRepositories"
)

func main() {

	location := locationrepositories.NewLocationRepo()

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
}

func oldTest() {
	cp := domain.LatLng{
		Lat: 6.2409826,
		Lng: -75.5862183,
	}

	var maxItem int32 = 4
	var r float32 = 0.05

	gen := fakedata.NewDummyDataGenerator(maxItem, cp, r)
	gen.GeneratePoints()

	menLocationService := menrepositories.NewLocationRepo(gen.GetAllData())
	menSpotService := menrepositories.NewMenSpotService(gen.GetAllData())

	service := filtersrv.New(menLocationService, menSpotService)

	locations := service.FilterByProximity(cp.Lat, cp.Lng, r)

	println("RESULT  ***********************************************************")
	marsh, _ := json.Marshal(locations)
	fmt.Printf("%+v\n", string(marsh))

	println("Size -> ", len(locations.Places))
	println("Actual time -> ", time.Now().String())
}
