package main

import (
	"log"

	"github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/cmd/ssm"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/services/spotsrv"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/handlers/httphdl"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/repositories/neo4jRepository"
	spotactivityservices "github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/repositories/spotActivityServices"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/pkg/uuidgen"
	"github.com/gin-gonic/gin"
)

var (
	repoSpot     *neo4jRepository.Neo4jSpotRepo
	repoLocation *spotactivityservices.AWSSpotActivityTopic
)

func init() {
	credsFinder := ssm.NewSSMCredentialsFinder()

	neo4jDriver, err := credsFinder.FindNeo4jCredentialsFromDefaultEnv()

	if err != nil {
		log.Println("There were an error while attempting to create drivers")
		log.Fatalln(err.Error())
	}

	repoSpot = neo4jRepository.NewNeo4jSpotRepoWithDriver(neo4jDriver) //menRepository.New()
	repoLocation = spotactivityservices.NewAWSSpotActivityTopic()
}

func main() {

	uuid := uuidgen.New()

	service := spotsrv.New(repoSpot, repoLocation, uuid)
	httpHandler := httphdl.NewHTTPHandler(service)

	router := gin.Default()
	router.GET("/spots/core/:id", httpHandler.GetEvent)
	router.POST("/spots/core/online", httpHandler.GoOnline)
	router.POST("/spots/core/getSpots", httpHandler.GetEvents)

	router.Run(":8000")
}
