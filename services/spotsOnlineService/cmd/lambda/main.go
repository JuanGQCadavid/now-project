package main

import (
	"context"
	"log"

	"github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/cmd/ssm"
	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/core/services"
	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/handlers/httphdl"
	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/repositories/neo4j"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	// stdout and stderr are sent to AWS CloudWatch  Logs
	log.Printf("Gin cold start")

	credsFinder := ssm.NewSSMCredentialsFinder()

	neo4jDriver, err := credsFinder.FindNeo4jCredentialsFromDefaultEnv()

	if err != nil {
		log.Println("There were an error while attempting to create drivers")
		log.Fatalln(err.Error())
	}

	repoSpot := neo4j.NewNeo4jRepoWithDriver(neo4jDriver)

	service := services.NewService(repoSpot)
	httpHandler := httphdl.NewHTTPHandler(service)

	router := gin.Default()
	router.POST("/spots/online/:spot_uuid/start", httpHandler.Start)
	router.POST("/spots/online/:spot_uuid/stop", httpHandler.Stop)
	router.POST("/spots/online/:spot_uuid/resume", httpHandler.Resume)
	router.POST("/spots/online/:spot_uuid/finalize", httpHandler.Finalize)

	ginLambda = ginadapter.New(router)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
