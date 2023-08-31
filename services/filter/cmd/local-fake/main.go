package main

import (
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/services/filtersrv"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/handlers/httphdl"
	dbspotservicelambda "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/dbSpotServiceLambda"
	locationrepositories "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/locationRepositories"
	sessionservice "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/sessionService"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/gin-gonic/gin"
)

func main() {

	dbDriver, err := locationrepositories.NewConectorFromEnv()

	if err != nil {
		logs.Error.Println("There were an error while attempting to create drivers")
		logs.Error.Fatalln(err.Error())
	}

	locationrepositories.Migrate(dbDriver)
	dbDriver.AutoMigrate(&dbspotservicelambda.SpotsDB{})

	// TODO -> How can we return an error from an init method ?
	locationRepo, err := locationrepositories.NewLocationRepoWithDriver(dbDriver)

	if err != nil {
		logs.Error.Println("There were an error while attempting to create the repository")
		logs.Error.Fatalln(err.Error())
	}

	spotSrv, err := dbspotservicelambda.NewDBSpotServiceLambdaWithDriver(dbDriver)

	if err != nil {
		panic(err.Error())
	}

	filterSrv := filtersrv.New(locationRepo, spotSrv)
	sessionHdl := sessionservice.NewSearchSessionDynamoDbService()
	filterHandler := httphdl.NewHTTPHandler(filterSrv, sessionHdl)

	router := gin.Default()
	router.GET("/filter/proximity", filterHandler.FilterSpots)
	router.Static("/filter/swagger", "../../swagger")

	router.Run("0.0.0.0:8000")
}
