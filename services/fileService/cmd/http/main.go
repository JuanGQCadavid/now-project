package main

import (
	"context"
	"flag"
	"os"

	"github.com/JuanGQCadavid/now-project/services/fileService/internal/adapters/objects3"
	"github.com/JuanGQCadavid/now-project/services/fileService/internal/adapters/spotscore"
	"github.com/JuanGQCadavid/now-project/services/fileService/internal/core/service"
	"github.com/JuanGQCadavid/now-project/services/fileService/internal/handler"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

var (
	bucket_name    string
	spots_core_dns string
)

const (
	bucket_name_env    string = "bucket_name"
	spots_core_dns_env string = "spots_core_dns"
)

func init() {
	var (
		oks bool
	)
	bucket_name, oks = os.LookupEnv(bucket_name_env)

	if !oks {
		log.Panic().Msg("Missing bucket_name env value")
	}

	spots_core_dns, oks = os.LookupEnv(spots_core_dns_env)

	if !oks {
		log.Panic().Msg("Missing spots_core_dns env value")
	}
}

func main() {

	var (
		router = gin.Default()
		debug  = flag.Bool("debug", false, "sets log level to debug")
	)

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Err(err).Msg("Could not start dependencies")
		panic(err)
	}

	srv := service.NewFileService(
		objects3.NewS3ObjectStorage(bucket_name, cfg),
		spotscore.NewSpotsCoreService(spots_core_dns),
	)

	hdl := handler.NewHttpHandler(srv)

	flag.Parse()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	hdl.SetRouter(router)
	router.Run(":8002")
}
