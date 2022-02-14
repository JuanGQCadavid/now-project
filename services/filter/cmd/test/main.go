package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/models"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/services/filtersrv"
	fakedata "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/fakeData"
	locationrepositories "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/locationRepositories"
	menrepositories "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/menRepositories"
)

func main() {

	location := locationrepositories.NewLocationRepo()

	a := models.LatLng{
		Lat: 6.26174,
		Lng: -75.60846,
	}

	b := models.LatLng{
		Lat: 6.24706,
		Lng: -75.5961,
	}
	places, err := location.FetchSpotsIdsByArea(a, b)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", places)
}

func oldTest() {
	cp := models.LatLng{
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
