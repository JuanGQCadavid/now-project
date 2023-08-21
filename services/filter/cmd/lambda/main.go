package main

import (
	"context"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/services/filtersrv"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/handlers/httphdl"
	locationrepositories "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/locationRepositories"
	sessionservice "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/sessionService"
	spotservicelambda "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/spotServiceLambda"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	dbDriver, err := locationrepositories.NewConectorFromEnv()

	if err != nil {
		logs.Error.Println("There were an error while attempting to create drivers")
		logs.Error.Fatalln(err.Error())
	}

	// TODO -> How can we return an error from an init method ?
	locationRepo, err := locationrepositories.NewLocationRepoWithDriver(dbDriver)

	if err != nil {
		logs.Error.Println("There were an error while attempting to create the repository")
		logs.Error.Fatalln(err.Error())
	}

	spotSrv, err := spotservicelambda.NewSpotServiceLambda()

	if err != nil {
		panic(err.Error())
	}

	filterSrv := filtersrv.New(locationRepo, spotSrv)
	sessionHdl := sessionservice.NewSearchSessionDynamoDbService()
	filterHandler := httphdl.NewHTTPHandler(filterSrv, sessionHdl)

	router := gin.Default()
	router.GET("/filter/proximity", filterHandler.FilterSpots)
	router.Static("/filter/swagger", "./swagger")

	ginLambda = ginadapter.New(router)

}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
