package main

import (
	"log"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/cmd/ssm"
	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/core/services"
	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/handlers/httphdl"
	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/notifiers/topics"
	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/repositories/neo4j"
	"github.com/gin-gonic/gin"
)

var (
	repoSpot *neo4j.Neo4jRepo
	notifier ports.Notify
)

const (
	TopicArnEnvName = "snsArn"
)

func init() {
	credsFinder := ssm.NewSSMCredentialsFinder()

	neo4jDriver, err := credsFinder.FindNeo4jCredentialsFromDefaultEnv()

	if err != nil {
		log.Println("There were an error while attempting to create drivers")
		log.Fatalln(err.Error())
	}

	repoSpot = neo4j.NewNeo4jRepoWithDriver(neo4jDriver)

	notifier, err = topics.NewNotifierFromEnv(TopicArnEnvName)

	if err != nil {
		logs.Error.Fatalln("We have a problem seting up the server, notifer error", err.Error())
	}
}

func main() {
	service := services.NewService(repoSpot, notifier)
	httpHandler := httphdl.NewHTTPHandler(service)

	router := gin.Default()
	httpHandler.InjectDefaultPaths(router)

	router.Run("localhost:8000")
}
