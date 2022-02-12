package httphdl

import (
	"fmt"
	"log"

	"github.com/JuanGQCadavid/now-project/services/spots/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spots/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	spotService ports.SpotService
}

func NewHTTPHandler(spotService ports.SpotService) *HTTPHandler {
	return &HTTPHandler{
		spotService: spotService,
	}
}

func (hdl *HTTPHandler) GetEvent(context *gin.Context) {

	event, err := hdl.spotService.Get(context.Param("id"))

	if err != nil {
		context.AbortWithStatusJSON(500, gin.H{"message": err})
		return
	}

	context.JSON(200, event)
}

func (hdl *HTTPHandler) GetEvents(context *gin.Context) {

	spotIds := SpotsIdsRequest{}
	context.BindJSON(&spotIds)

	log.Printf("%+v", spotIds)

	multipleSpots, err := hdl.spotService.GetSpots(spotIds.SpotIds)
	if err != nil {
		context.AbortWithStatusJSON(500, gin.H{"message": err})
		return
	}

	context.JSON(200, multipleSpots)
}

func (hdl *HTTPHandler) GoOnline(context *gin.Context) {
	log.Println("Context -> ", fmt.Sprintf("%+v", context))
	log.Println("Calling GoOnline")

	spot := domain.Spot{}
	context.BindJSON(&spot)

	log.Println("Spot -> ", fmt.Sprintf("%+v", spot))

	spot, err := hdl.spotService.GoOnline(spot)
	if err != nil {
		context.AbortWithStatusJSON(500, gin.H{"message": err})
		return
	}

	context.JSON(200, spot)
}
