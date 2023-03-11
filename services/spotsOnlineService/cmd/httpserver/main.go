package main

import (
	"log"

	"github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/cmd/ssm"
	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/core/services"
	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/handlers/httphdl"
	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/repositories/neo4j"
	"github.com/gin-gonic/gin"
)

// import (
// 	"log"

// 	"github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/cmd/ssm"
// 	"github.com/JuanGQCadavid/now-project/services/spotsCore/internal/core/services/spotsrv"
// 	"github.com/JuanGQCadavid/now-project/services/spotsCore/internal/handlers/httphdl"
// 	"github.com/JuanGQCadavid/now-project/services/spotsCore/internal/repositories/neo4jRepository"
// 	spotactivityservices "github.com/JuanGQCadavid/now-project/services/spotsCore/internal/repositories/spotActivityServices"
// 	"github.com/JuanGQCadavid/now-project/services/spotsCore/pkg/uuidgen"
// 	"github.com/gin-gonic/gin"
// )

var (
	repoSpot *neo4j.Neo4jRepo
)

func init() {
	credsFinder := ssm.NewSSMCredentialsFinder()

	neo4jDriver, err := credsFinder.FindNeo4jCredentialsFromDefaultEnv()

	if err != nil {
		log.Println("There were an error while attempting to create drivers")
		log.Fatalln(err.Error())
	}

	repoSpot = neo4j.NewNeo4jRepoWithDriver(neo4jDriver)
}

func main() {
	service := services.NewService(repoSpot)
	httpHandler := httphdl.NewHTTPHandler(service)

	router := gin.Default()
	router.POST("/spots/online/:spot_uuid/start", httpHandler.Start)

	router.Run("localhost:8000")
}
