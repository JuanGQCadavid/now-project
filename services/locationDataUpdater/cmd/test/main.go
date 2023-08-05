package main

import (
	locationrepositories "github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/repositories/locationRepositories"
)

func main() {
	connector, err := locationrepositories.NewConector("admin", "admin", "pululapp", "localhost")
	connector.Migrate()

	if err != nil {
		panic(err)
	}

	connector.Migrate()

	repo, err := locationrepositories.NewLocationRepo(connector)

	if err != nil {
		panic(err)
	}

	// date := domain.DatesLocation{
	// 	DateID: "OTHER_TEST_123",
	// 	Lat:    123.456,
	// 	Lon:    789.123,
	// 	Type: domain.Types{
	// 		TypeID:      domain.Online,
	// 		Description: "It is online, babe",
	// 	},
	// 	State: domain.States{
	// 		StateID:     domain.OnlineDateStatus,
	// 		Description: "Active date",
	// 	},
	// }
	// repo.CrateLocation(date)
	//repo.UpdateLocationStatus("OTHER_TEST_123", domain.StoppedDateStatus)

	repo.RemoveLocation("OTHER_TEST_123")

}
