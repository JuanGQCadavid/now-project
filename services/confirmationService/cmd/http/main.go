package main

import (
	"github.com/JuanGQCadavid/now-project/services/confirmationService/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/confirmationService/internal/core/service"
	"github.com/JuanGQCadavid/now-project/services/confirmationService/internal/handlers/httphdl"
	"github.com/JuanGQCadavid/now-project/services/confirmationService/internal/notifiers/topics"
	"github.com/JuanGQCadavid/now-project/services/confirmationService/internal/repositories/neo4j"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/cmd/ssm"
	"github.com/gin-gonic/gin"
)

const (
	TopicArnEnvName = "snsArn"
)

var (
	repoSpot ports.Repository
	notifier ports.Notify
)

func init() {
	logs.Info.Println("Gin cold start")

	credsFinder := ssm.NewSSMCredentialsFinder()
	neo4jDriver, err := credsFinder.FindNeo4jCredentialsFromDefaultEnv()

	if err != nil {
		logs.Error.Println("There were an error while attempting to create drivers")
		logs.Error.Fatalln(err.Error())
	}

	repoSpot = neo4j.NewNeo4jRepoWithDriver(neo4jDriver)
	notifier, err = topics.NewNotifierFromEnv(TopicArnEnvName)

	if err != nil {
		logs.Error.Fatalln("We have a problem seting up the server, notifer error", err.Error())
	}

}

func main() {
	service := service.NewConfirmationService(repoSpot, notifier)
	httpHandler := httphdl.NewHttpHandler(service)

	router := gin.Default()
	httpHandler.SetRouter(router)

	router.Run("localhost:8000")
}
