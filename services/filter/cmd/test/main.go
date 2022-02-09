package main

import (
	"fmt"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/models"
	fakedata "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/fakeData"
)

func main() {

	spotType := models.Online

	fmt.Println(spotType)

	spot := models.Spot{
		Id:        "DDE",
		Type:      models.Online,
		Emoji:     ":p",
		StartTime: "i dont lnow",
	}

	fmt.Printf("%+v\n", spot)

	example := models.Locations{
		Places: []models.Spot{spot, spot},
	}

	fmt.Printf("%+v\n", example)

	println("Testing random generator")

	cp := models.LatLng{
		Lat: 6.2409826,
		Lng: -75.5862183,
	}

	gen := fakedata.NewDummyDataGenerator(10, cp, 0.05)
	gen.GeneratePoints()

}
