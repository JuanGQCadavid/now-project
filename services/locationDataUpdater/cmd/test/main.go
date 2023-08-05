package main

import (
	"github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/repositories/rds"
)

func main() {
	connector, err := rds.NewConector("admin", "admin", "pululapp", "localhost")
	connector.Migrate()

	if err != nil {
		panic(err)
	}

	connector.Migrate()

	repo, err := rds.NewRDSRepo(connector)

	if err != nil {
		panic(err)
	}

	date := domain.DatesLocation{
		DateID: "OTHER_TEST_123",
		Lat:    123.456,
		Lon:    789.123,
		Type: domain.Types{
			TypeID:      domain.Online,
			Description: "It is online, babe",
		},
		State: domain.States{
			StateID:     domain.OnlineDateStatus,
			Description: "Active date",
		},
	}
	repo.CrateLocation(date)
	repo.UpdateLocationStatus("OTHER_TEST_123", domain.StoppedDateStatus)
	repo.UpdateLocationType("OTHER_TEST_123", domain.Scheduled)

	// repo.RemoveLocation("OTHER_TEST_123")

}
