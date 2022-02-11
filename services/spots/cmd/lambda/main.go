package main

import (
	"context"
	"log"

	"github.com/JuanGQCadavid/now-project/services/spots/internal/core/services/spotsrv"
	"github.com/JuanGQCadavid/now-project/services/spots/internal/handlers/httphdl"
	"github.com/JuanGQCadavid/now-project/services/spots/internal/repositories/menRepository"
	"github.com/JuanGQCadavid/now-project/services/spots/internal/repositories/neo4jRepository"
	"github.com/JuanGQCadavid/now-project/services/spots/pkg/uuidgen"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")

	repoSpot := neo4jRepository.NewNeo4jSpotRepo() //menRepository.New()
	repoLocation := menRepository.NewLocationRepository()
	uuid := uuidgen.New()

	service := spotsrv.New(repoSpot, repoLocation, uuid)
	httpHandler := httphdl.NewHTTPHandler(service)

	router := gin.Default()
	router.GET("/spot/:id", httpHandler.GetEvent)
	router.POST("/spot/online", httpHandler.GoOnline)
	router.POST("/spot/getSpots", httpHandler.GetEvents)

	ginLambda = ginadapter.New(router)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
