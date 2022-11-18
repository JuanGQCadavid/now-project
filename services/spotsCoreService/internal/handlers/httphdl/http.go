package httphdl

import (
	"log"
	"strings"

	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/ports"
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

	if len(id) == 0 {
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: "Missing Id param",
		})
		return
	}

	format := hdl.getOuputFormat(context.DefaultQuery("format", "empty"))

	log.Printf("\nHandler: GetEvent \n\tId: %s, \n\tFormat: %s", id, string(format))

	event, err := hdl.spotService.Get(id, format)

	if err != nil {
		log.Println(err)
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message:       "We face an error while fethcing the data",
			InternalError: err.Error(),
		})
		return
	}

	context.JSON(200, event)
}

func (hdl *HTTPHandler) GetEvents(context *gin.Context) {

	// Getting data from call
	spotIds := SpotsIdsRequest{}
	context.BindJSON(&spotIds)

	format := hdl.getOuputFormat(context.DefaultQuery("format", "empty"))

	log.Printf("\nHandler: GetEvents \n\tSpotIds: %+v, \n\tFormat: %s", spotIds, string(format))

	multipleSpots, err := hdl.spotService.GetSpots(spotIds.SpotIds, format)

	if err != nil {
		log.Println(err)
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message:       "We face an error while fethcing the data",
			InternalError: err.Error(),
		})
		return
	}

	context.JSON(200, multipleSpots)
}

func (hdl *HTTPHandler) GoOnline(context *gin.Context) {

	spot := domain.Spot{}
	context.BindJSON(&spot)

	log.Printf("\nHandler: GoOnline \n\tSpot: %+v", spot)

	if !hdl.isSpotCorrect(spot) {
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: "The spot is missing some data.",
		})
		return
	}

	spot, err := hdl.spotService.GoOnline(spot)

	if err != nil {
		log.Println(err)
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message:       "We face an error while Going online",
			InternalError: err.Error(),
		})
		return
	}

	context.JSON(200, spot)
}

func (hdl *HTTPHandler) getOuputFormat(query string) ports.OutputFormat {
	switch strings.ToUpper(query) {
	case string(ports.FULL_FORMAT):
		return ports.FULL_FORMAT
	case string(ports.SMALL_FORMAT):
		return ports.SMALL_FORMAT
	}
	return ports.FULL_FORMAT
}

func (hdl *HTTPHandler) isSpotCorrect(spot domain.Spot) bool {
	// TODO -> Could we use json schemas here ?

	return true
}
