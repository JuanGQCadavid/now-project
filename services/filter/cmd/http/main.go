package main

import (
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/handlers/httphdl"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/services/filtersrv"
	locationrepositories "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/locationRepositories"
	spotservicelambda "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/spotServiceLambda"
	"github.com/gin-gonic/gin"
)

func main() {

	locationRepo := locationrepositories.NewLocationRepo()
	spotSrv := spotservicelambda.NewSpotServiceLambda()

	filterSrv := filtersrv.New(locationRepo, spotSrv)
	filterHandler := httphdl.NewHTTPHandler(filterSrv)

	router := gin.Default()
	router.GET("/filter/proximity", filterHandler.FilterSpots)

	router.Run(":8001")
}
