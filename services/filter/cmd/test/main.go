package main

import (
	"fmt"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/models"
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

}
