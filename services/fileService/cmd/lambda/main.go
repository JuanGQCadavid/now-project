package main

import (
	"context"
	"os"

	"github.com/JuanGQCadavid/now-project/services/fileService/internal/adapters/objects3"
	"github.com/JuanGQCadavid/now-project/services/fileService/internal/adapters/spotscore"
	"github.com/JuanGQCadavid/now-project/services/fileService/internal/core/service"
	"github.com/JuanGQCadavid/now-project/services/fileService/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var ginLambda *ginadapter.GinLambda

var (
	bucket_name    string
	spots_core_dns string
	hdl            *handler.HttpHandler
)

const (
	bucket_name_env    string = "bucket_name"
	spots_core_dns_env string = "spots_core_dns"
)

func init() {
	var (
		router = gin.Default()
		oks    bool
	)
	bucket_name, oks = os.LookupEnv(bucket_name_env)

	if !oks {
		log.Panic().Msg("Missing bucket_name env value")
	}

	spots_core_dns, oks = os.LookupEnv(spots_core_dns_env)

	if !oks {
		log.Panic().Msg("Missing spots_core_dns env value")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Err(err).Msg("Could not start dependencies")
		panic(err)
	}

	srv := service.NewFileService(
		objects3.NewS3ObjectStorage(bucket_name, cfg),
		spotscore.NewSpotsCoreService(spots_core_dns),
	)

	hdl = handler.NewHttpHandler(srv)
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	hdl.SetRouter(router)
	ginLambda = ginadapter.New(router)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
