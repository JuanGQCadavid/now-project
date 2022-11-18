package main

import (
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/services/spotsrv"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/handlers/httphdl"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/repositories/neo4jRepository"
	spotactivityservices "github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/repositories/spotActivityServices"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/pkg/uuidgen"
	"github.com/gin-gonic/gin"
)

func main() {
	repoSpot := neo4jRepository.NewNeo4jSpotRepo() //menRepository.New()
	repoLocation := spotactivityservices.NewAWSSpotActivityTopic()
	uuid := uuidgen.New()

	service := spotsrv.New(repoSpot, repoLocation, uuid)
	httpHandler := httphdl.NewHTTPHandler(service)

	router := gin.Default()
	router.GET("/spot/:id", httpHandler.GetEvent)
	router.POST("/spot/online", httpHandler.GoOnline)
	router.POST("/spot/getSpots", httpHandler.GetEvents)

	router.Run(":8000")
}
