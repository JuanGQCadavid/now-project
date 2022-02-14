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

	id := context.Param("id")
	format := hdl.getOuputFormat(context.DefaultQuery("format", "empty"))

	log.Println("Calling GetEvent with -> Id:", id, ", Format: ", string(format))

	event, err := hdl.spotService.Get(id, format)

	if err != nil {
		context.AbortWithStatusJSON(500, gin.H{"message": err})
		return
	}

	context.JSON(200, event)
}

func (hdl *HTTPHandler) GetEvents(context *gin.Context) {

	spotIds := SpotsIdsRequest{}
	context.BindJSON(&spotIds)

	format := hdl.getOuputFormat(context.DefaultQuery("format", "empty"))

	log.Println("Calling GetEvent with -> spotIds:", fmt.Sprintf("%+v", spotIds), ", Format: ", string(format))

	multipleSpots, err := hdl.spotService.GetSpots(spotIds.SpotIds, format)
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

func (hdl *HTTPHandler) getOuputFormat(query string) ports.OutputFormat {
	switch query {
	case string(ports.FULL_FORMAT):
		return ports.FULL_FORMAT
	case string(ports.SMALL_FORMAT):
		return ports.SMALL_FORMAT
	}
	return ports.FULL_FORMAT
}
