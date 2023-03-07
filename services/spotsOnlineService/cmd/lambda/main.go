// package main

// import (
// 	"context"
// 	"log"

// 	"github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/cmd/ssm"

// 	"github.com/JuanGQCadavid/now-project/services/spotsCore/internal/core/services/spotsrv"
// 	"github.com/JuanGQCadavid/now-project/services/spotsCore/internal/handlers/httphdl"
// 	"github.com/JuanGQCadavid/now-project/services/spotsCore/internal/repositories/neo4jRepository"
// 	spotactivityservices "github.com/JuanGQCadavid/now-project/services/spotsCore/internal/repositories/spotActivityServices"
// 	"github.com/JuanGQCadavid/now-project/services/spotsCore/pkg/uuidgen"
// 	"github.com/aws/aws-lambda-go/events"
// 	"github.com/aws/aws-lambda-go/lambda"
// 	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
// 	"github.com/gin-gonic/gin"
// )

// var ginLambda *ginadapter.GinLambda

// func init() {
// 	// stdout and stderr are sent to AWS CloudWatch  Logs
// 	log.Printf("Gin cold start")

// 	credsFinder := ssm.NewSSMCredentialsFinder()

// 	neo4jDriver, err := credsFinder.FindNeo4jCredentialsFromDefaultEnv()

// 	if err != nil {
// 		log.Println("There were an error while attempting to create drivers")
// 		log.Fatalln(err.Error())
// 	}

// 	repoSpot := neo4jRepository.NewNeo4jSpotRepoWithDriver(neo4jDriver) //menRepository.New()
// 	repoLocation := spotactivityservices.NewAWSSpotActivityTopic()
// 	uuid := uuidgen.New()

// 	service := spotsrv.New(repoSpot, repoLocation, uuid)
// 	httpHandler := httphdl.NewHTTPHandler(service)

// 	router := gin.Default()
// 	router.GET("/spots/core/:id", httpHandler.GetEvent)
// 	router.POST("/spots/core/online", httpHandler.GoOnline)
// 	router.POST("/spots/core/getSpots", httpHandler.GetEvents)

// 	ginLambda = ginadapter.New(router)
// }

// func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	// If no name is provided in the HTTP request body, throw an error
// 	return ginLambda.ProxyWithContext(ctx, req)
// }

// func main() {
// 	lambda.Start(Handler)
// }
