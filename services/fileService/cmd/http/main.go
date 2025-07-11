package main

import (
	"flag"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/JuanGQCadavid/now-project/services/fileService/internal/core"
	"github.com/JuanGQCadavid/now-project/services/fileService/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func main() {

	var (
		router = gin.Default()
		srv    = core.NewFileService()
		hdl    = handler.NewHttpHandler(srv)
		debug  = flag.Bool("debug", false, "sets log level to debug")
	)

	flag.Parse()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	hdl.SetRouter(router)
	router.Run(":8002")
}
