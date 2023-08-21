package httphdl

import (
	"fmt"
	"strings"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
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

	logs.Info.Printf("\nHandler: GetEvent \n\tId: %s, \n\tFormat: %s", id, string(format))

	event, err := hdl.spotService.Get(id, format)

	if err != nil {
		logs.Error.Println(err)

		if err == ports.ErrSpotNotFounded {
			context.AbortWithStatusJSON(404, ErrorMessage{
				Message:       "The spot does not exist",
				InternalError: err.Error(),
			})
			return
		}

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
	datesIds := SpotsIdsRequest{}
	context.BindJSON(&datesIds)

	format := hdl.getOuputFormat(context.DefaultQuery("format", "empty"))

	logs.Info.Printf("\nHandler: GetEvents \n\tDateIds: %+v, \n\tFormat: %s", datesIds, string(format))

	multipleSpots, err := hdl.spotService.GetSpotsByDatesIds(datesIds.DatesIds, format)

	if err != nil {
		logs.Error.Println(err)
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

	logs.Info.Printf("\nHandler: CreateSpot \n\tSpot: %+v", spot)

	if !hdl.isSpotCorrect(spot) {
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: "The spot is missing some data.",
		})
		return
	}

	spot, err := hdl.spotService.CreateSpot(spot)

	if err != nil {
		logs.Error.Println(err.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message:       "We face an error while creating spot",
			InternalError: err.Error(),
		})
		return
	}

	context.JSON(200, spot)
}

// /spots/core/:id/event
func (hdl *HTTPHandler) UpdateSpotEvent(context *gin.Context) {
	id := context.Param("id")
	userId := context.Request.Header.Get("Authorization")
	spotEvent := domain.Spot{}
	context.BindJSON(&spotEvent)

	logs.Info.Printf("Handler - UpdateSpotEvent: UserId %s,  Id %s, spotEvent %+v\n", userId, id, fmt.Sprintf("%+v", spotEvent.EventInfo))

	if err := hdl.spotService.UpdateSpotEvent(id, userId, &spotEvent.EventInfo); err != nil {
		logs.Info.Println("Hanlder - UpdateSpotEvent - Error", err.Error())

		if err == ports.ErrSpotToUpdateIsTheSameAsTheDb {
			context.Status(204)
			return
		}
		context.AbortWithStatusJSON(500, ErrorMessage{
			Message:       "We face an error while updating the spot event info",
			InternalError: err.Error(),
		})
		return
	}

	context.Status(204)
}

// /spots/core/:id/place
func (hdl *HTTPHandler) UpdateSpotPlace(context *gin.Context) {
	id := context.Param("id")
	context.JSON(200, map[string]string{
		"method": "UpdateSpotPlace",
		"id":     id,
	})
}

// /spots/core/:id/topic
func (hdl *HTTPHandler) UpdateSpotTopic(context *gin.Context) {
	id := context.Param("id")
	context.JSON(200, map[string]string{
		"method": "UpdateSpotTopic",
		"id":     id,
	})
}

// /spots/core/:id
func (hdl *HTTPHandler) DeleteSpot(context *gin.Context) {
	id := context.Param("id")
	userRequestedId := context.Request.Header.Get("Authorization")
	logs.Info.Printf("Handler - DeleteSpot: userRequestedId %s,  Id %s", userRequestedId, id)

	if err := hdl.spotService.DeleteSpot(id, userRequestedId); err != nil {
		logs.Info.Println("Hanlder - DeleteSpot - Error", err.Error())

		if err == ports.ErrSpotUserNotOwnerWhenUpdatingSpot {
			context.AbortWithStatusJSON(401, ErrorMessage{
				Message:       "The user is not the owner of the spot.",
				InternalError: err.Error(),
			})
			return
		}

		context.AbortWithStatusJSON(500, ErrorMessage{
			Message:       "We face an error while deleting the spot",
			InternalError: err.Error(),
		})
		return
	}

	context.Status(204)
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
