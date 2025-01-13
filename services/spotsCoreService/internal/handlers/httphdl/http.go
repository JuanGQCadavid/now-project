package httphdl

import (
	"fmt"
	"net/http"
	"strings"

	authDomain "github.com/JuanGQCadavid/now-project/services/authService/core/core/domain"
	authUtils "github.com/JuanGQCadavid/now-project/services/authService/core/utils"
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

func (hdl *HTTPHandler) SetRouter(router *gin.Engine) {
	router.POST("/spots/core/", hdl.CreateSpot)                 // OK
	router.POST("/spots/core/bulk/fetch", hdl.GetMultipleSpots) // OK
	router.GET("/spots/core/:id", hdl.GetSpot)                  // OK
	router.PUT("/spots/core/:id/event", hdl.UpdateSpotEvent)    // OK
	router.PUT("/spots/core/:id/topic", hdl.UpdateSpotTopic)    // OK
	router.PUT("/spots/core/:id/place", hdl.UpdateSpotPlace)    // OK
	router.DELETE("/spots/core/:id", hdl.DeleteSpot)            // OK
}

// /spots/core/:id
func (hdl *HTTPHandler) GetSpot(context *gin.Context) {

	id := context.Param("id")
	if len(id) == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{
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
			context.AbortWithStatusJSON(http.StatusNotFound, ErrorMessage{
				Message:       "The spot does not exist",
				InternalError: err.Error(),
			})
			return
		}

		context.AbortWithStatusJSON(http.StatusInternalServerError, ErrorMessage{
			Message:       "We face an error while fetching the data",
			InternalError: err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, event)
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
		context.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{
			Message:       "We face an error while fethcing the data",
			InternalError: err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, multipleSpots)
}

// /spots/core/
func (hdl *HTTPHandler) CreateSpot(context *gin.Context) {

	spot := domain.Spot{}
	context.BindJSON(&spot)

	var userDetails authDomain.UserDetails = *authUtils.GetHeaders(context.Request.Header)

	logs.Info.Printf("\nHandler: CreateSpot \n\tSpot: %+v", spot)

	if !hdl.isSpotCorrect(spot) {
		context.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{
			Message: "The spot is missing some data.",
		})
		return
	}

	spot, err := hdl.spotService.CreateSpot(spot, userDetails)

	if err != nil {
		logs.Error.Println(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{
			Message:       "We face an error while creating spot",
			InternalError: err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, spot)
}

// /spots/core/:id/event
func (hdl *HTTPHandler) UpdateSpotEvent(context *gin.Context) {
	var (
		id                                 = context.Param("id")
		spotEvent                          = domain.Spot{}
		userDetails authDomain.UserDetails = *authUtils.GetHeaders(context.Request.Header)
	)

	context.BindJSON(&spotEvent)

	logs.Info.Printf("Handler - UpdateSpotEvent: UserId %s,  Id %s, spotEvent %+v\n", userDetails.UserID, id, fmt.Sprintf("%+v", spotEvent.EventInfo))

	if err := hdl.spotService.UpdateSpotEvent(id, userDetails.UserID, &spotEvent.EventInfo); err != nil {
		logs.Info.Println("Hanlder - UpdateSpotEvent - Error", err.Error())

		if err == ports.ErrSpotToUpdateIsTheSameAsTheDb {
			context.Status(http.StatusNoContent)
			return
		}
		context.AbortWithStatusJSON(http.StatusInternalServerError, ErrorMessage{
			Message:       "We face an error while updating the spot event info",
			InternalError: err.Error(),
		})
		return
	}

	context.Status(http.StatusNoContent)
}

// /spots/core/:id/place
func (hdl *HTTPHandler) UpdateSpotPlace(context *gin.Context) {
	id := context.Param("id")
	context.JSON(http.StatusNotImplemented, map[string]string{
		"method": "UpdateSpotPlace",
		"id":     id,
	})
}

// /spots/core/:id/topic
func (hdl *HTTPHandler) UpdateSpotTopic(context *gin.Context) {
	id := context.Param("id")
	context.JSON(http.StatusNotImplemented, map[string]string{
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

		context.AbortWithStatusJSON(http.StatusInternalServerError, ErrorMessage{
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
