package main

import (
	"context"
	"log"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/services/filtersrv"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/handlers/httphdl"
	locationrepositories "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/locationRepositories"
	spotservicelambda "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/spotServiceLambda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	log.Println("Filter service")

	locationRepo := locationrepositories.NewLocationRepo()
	spotSrv := spotservicelambda.NewSpotServiceLambda()

	filterSrv := filtersrv.New(locationRepo, spotSrv)
	filterHandler := httphdl.NewHTTPHandler(filterSrv)

	router := gin.Default()
	router.GET("/filter/proximity", filterHandler.FilterSpots)

	ginLambda = ginadapter.New(router)

}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
