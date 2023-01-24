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

// /spots/core/:id
func (hdl *HTTPHandler) GetSpot(context *gin.Context) {

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

// /spots/core/bulk/fetch
func (hdl *HTTPHandler) GetMultipleSpots(context *gin.Context) {

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

// /spots/core/
func (hdl *HTTPHandler) CreateSpot(context *gin.Context) {

	spot := domain.Spot{}
	context.BindJSON(&spot)

	log.Printf("\nHandler: CreateSpot \n\tSpot: %+v", spot)

	if !hdl.isSpotCorrect(spot) {
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: "The spot is missing some data.",
		})
		return
	}

	spot, err := hdl.spotService.CreateSpot(spot)

	if err != nil {
		log.Println(err)
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message:       "We face an error while creating spot",
			InternalError: err.Error(),
		})
		return
	}

	context.JSON(200, spot)
}

// /spots/core/:id/finalize
func (hdl *HTTPHandler) FinalizeSpot(context *gin.Context) {
	id := context.Param("id")
	context.JSON(200, map[string]string{
		"method": "FinalizeSpot",
		"id":     id,
	})
}

// /spots/core/:id
func (hdl *HTTPHandler) UpdateSpot(context *gin.Context) {
	id := context.Param("id")
	context.JSON(200, map[string]string{
		"method": "UpdateSpot",
		"id":     id,
	})
}

// /spots/core/:id
func (hdl *HTTPHandler) DeleteSpot(context *gin.Context) {
	id := context.Param("id")
	context.JSON(200, map[string]string{
		"method": "DeleteSpot",
		"id":     id,
	})
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
