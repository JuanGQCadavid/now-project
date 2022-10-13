package main

import (
	"log"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/services/filtersrv"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/handlers/httphdl"
	locationrepositories "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/locationRepositories"
	sessionservice "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/sessionService"
	spotservicelambda "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/spotServiceLambda"
	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.LstdFlags)
	locationRepo, err := locationrepositories.NewLocationRepo()

	if err != nil {
		panic(err.Error())
	}
	spotSrv, err := spotservicelambda.NewSpotServiceLambda()

	if err != nil {
		panic(err.Error())
	}

	filterSrv := filtersrv.New(locationRepo, spotSrv)
	sessionHdl := sessionservice.NewSearchSessionDynamoDbService()
	filterHandler := httphdl.NewHTTPHandler(filterSrv, sessionHdl)

	router := gin.Default()
	router.GET("/filter/proximity", filterHandler.FilterSpots)
	router.Static("/swagger", "../../swagger")

	router.Run(":8001")
}
