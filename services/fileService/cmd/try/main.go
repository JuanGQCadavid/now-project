package main

import (
	"context"
	"os"

	"github.com/JuanGQCadavid/now-project/services/fileService/internal/adapters/spotscore"
	"github.com/JuanGQCadavid/now-project/services/fileService/internal/core/ports"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	var (
		adapter ports.SpotsCoreRepository = spotscore.NewSpotsCoreService("http://localhost:8000")
		eventId                           = "cf4d1ccf-dfdf-4868-81ae-c58a2f423f68"
		userId                            = "b3439f92-5d3e-458e-8e02-d2220245c524"
		dateId                            = "88fd219b-1671-4c1d-b935-130dd252b79a"
	)

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	access, err := adapter.GetUserEventAccess(log.Logger.WithContext(context.TODO()), userId, eventId)

	if err != nil {
		panic(err.Error())
	}

	log.Info().Bool("Is attending", access.IsAttending).Bool("is Hosting", access.IsHoster).Send()

	// Date

	access2, err := adapter.GetUserDateAccess(log.Logger.WithContext(context.TODO()), eventId, userId, dateId)

	if err != nil {
		panic(err.Error())
	}

	log.Info().Bool("Is attending", access2.IsAttending).Bool("is Hosting", access2.IsHoster).Send()
}
