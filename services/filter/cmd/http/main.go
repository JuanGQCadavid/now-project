package main

import (
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/services/filtersrv"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/handlers/httphdl"
	locationrepositories "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/locationRepositories"
	sessionservice "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/sessionService"
	spotservicelambda "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/spotServiceLambda"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/cmd/ssm"
	"github.com/gin-gonic/gin"
)

func main() {
	credsFinder := ssm.NewSSMCredentialsFinder()
	credentials, err := credsFinder.GetDBCredentialsFromDefaultEnv()

	if err != nil {
		logs.Error.Fatalln("we fail to Fetch the envs")
	}

	dbDriver, err := locationrepositories.NewConector(credentials.User, credentials.Password, credentials.Name, credentials.Url)

	if err != nil {
		logs.Error.Println("There were an error while attempting to create drivers")
		logs.Error.Fatalln(err.Error())
	}

	// TODO -> How can we return an error from an init method ?
	locationRepo, err := locationrepositories.NewLocationRepoWithDriver(dbDriver)

	if err != nil {
		logs.Error.Println("There were an error while attempting to create the repository")
		logs.Error.Fatalln(err.Error())
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
	router.Static("/filter/swagger", "../../swagger")

	router.Run(":8000")
}
