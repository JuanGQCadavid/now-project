package main

import (
	"context"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"

	"github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/cmd/ssm"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/services/spotsrv"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/handlers/httphdl"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/notifiers/topics"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/repositories/neo4jRepository"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/pkg/uuidgen"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

const (
	TopicArnEnvName = "snsArn"
)

func init() {
	// stdout and stderr are sent to AWS CloudWatch  Logs
	logs.Info.Printf("Gin cold start")

	credsFinder := ssm.NewSSMCredentialsFinder()

	neo4jDriver, err := credsFinder.FindNeo4jCredentialsFromDefaultEnv()

	if err != nil {
		logs.Error.Println("There were an error while attempting to create drivers")
		logs.Error.Fatalln(err.Error())
	}

	repoSpot := neo4jRepository.NewNeo4jSpotRepoWithDriver(neo4jDriver) //menRepository.New()
	notifier, err := topics.NewNotifierFromEnv(TopicArnEnvName)

	if err != nil {
		logs.Error.Fatalln("We have a problem seting up the server, notifer error", err.Error())
	}

	uuid := uuidgen.New()

	service := spotsrv.New(repoSpot, notifier, uuid)
	httpHandler := httphdl.NewHTTPHandler(service)
	router := gin.Default()
	router.POST("/spots/core/", httpHandler.CreateSpot)                 // OK
	router.POST("/spots/core/bulk/fetch", httpHandler.GetMultipleSpots) // OK
	router.GET("/spots/core/:id", httpHandler.GetSpot)                  // OK
	router.PUT("/spots/core/:id/event", httpHandler.UpdateSpotEvent)    // OK
	router.PUT("/spots/core/:id/topic", httpHandler.UpdateSpotTopic)    // OK
	router.PUT("/spots/core/:id/place", httpHandler.UpdateSpotPlace)    // OK
	router.DELETE("/spots/core/:id", httpHandler.DeleteSpot)            // OK

	ginLambda = ginadapter.New(router)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
