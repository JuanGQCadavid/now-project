package main

import (
	"context"
	"log"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/cmd/ssm"
	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/core/services"
	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/handlers/httphdl"
	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/notifiers/topics"
	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/repositories/neo4j"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

const (
	TopicArnEnvName = "snsArn"
)

var ginLambda *ginadapter.GinLambda

func init() {
	credsFinder := ssm.NewSSMCredentialsFinder()

	neo4jDriver, err := credsFinder.FindNeo4jCredentialsFromDefaultEnv()

	if err != nil {
		log.Println("There were an error while attempting to create drivers")
		log.Fatalln(err.Error())
	}

	repoSpot := neo4j.NewNeo4jRepoWithDriver(neo4jDriver)

	notifier, err := topics.NewNotifierFromEnv(TopicArnEnvName)

	if err != nil {
		logs.Error.Fatalln("We have a problem seting up the server, notifer error", err.Error())
	}

	service := services.NewService(repoSpot, notifier)
	httpHandler := httphdl.NewHTTPHandler(service)

	router := gin.Default()
	httpHandler.InjectDefaultPaths(router)

	ginLambda = ginadapter.New(router)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
