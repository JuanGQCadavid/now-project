package main

import (
	"flag"
	"os"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/cmd/ssm"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/services/spotsrv"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/handlers/httphdl"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/notifiers/topics"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/repositories/neo4jRepository"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/pkg/uuidgen"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	repoSpot *neo4jRepository.Neo4jSpotRepo
	notifier *topics.Notifier
)

const (
	TopicArnEnvName = "snsArn"
)

func init() {
	credsFinder := ssm.NewSSMCredentialsFinder()

	neo4jDriver, err := credsFinder.FindNeo4jCredentialsFromDefaultEnv()

	if err != nil {
		logs.Error.Println("There were an error while attempting to create drivers")
		logs.Error.Println(err.Error())
	}

	repoSpot = neo4jRepository.NewNeo4jSpotRepoWithDriver(neo4jDriver)
	notifier, err = topics.NewNotifierFromEnv(TopicArnEnvName)
	if err != nil {
		logs.Error.Fatalln("We have a problem seting up the server, notifer error", err.Error())
	}
}

func main() {
	var (
		uuid        = uuidgen.New()
		service     = spotsrv.New(repoSpot, notifier, uuid)
		httpHandler = httphdl.NewHTTPHandler(service)
		router      = gin.Default()
		debug       = flag.Bool("debug", false, "sets log level to debug")
	)

	flag.Parse()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	httpHandler.SetRouter(router)
	// router.Run("localhost:8000")
	router.Run(":8000")
}
