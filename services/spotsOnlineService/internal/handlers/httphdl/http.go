package httphdl

import (
	"log"

	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/core/ports"
	"github.com/gin-gonic/gin"
)

// import (
// 	"log"
// 	"strings"

// 	"github.com/JuanGQCadavid/now-project/services/spotsCore/internal/core/domain"
// 	"github.com/JuanGQCadavid/now-project/services/spotsCore/internal/core/ports"
// 	"github.com/gin-gonic/gin"
// )

type HTTPHandler struct {
	spotOnlineService ports.SpotOnlineService
}

func NewHTTPHandler(spotOnlineService ports.SpotOnlineService) *HTTPHandler {
	return &HTTPHandler{
		spotOnlineService: spotOnlineService,
	}
}

func (hdl *HTTPHandler) InjectDefaultPaths(router *gin.Engine) {
	router.GET("/spots/online/:spot_uuid/", hdl.Get)
	router.POST("/spots/online/:spot_uuid/start", hdl.Start)
	router.POST("/spots/online/:spot_uuid/stop", hdl.Stop)
	router.POST("/spots/online/:spot_uuid/resume", hdl.Resume)
	router.POST("/spots/online/:spot_uuid/finalize", hdl.Finalize)
}

func (hdl *HTTPHandler) Get(context *gin.Context) {
	// Path Variables
	spot_uudi, is_spot_uudi_present := context.Params.Get("spot_uuid")

	if !is_spot_uudi_present {
		log.Println("Spot id is mising in the path")
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: "The spot id is missing",
		})
		// return not id sent
		return
	}

	spotWitjDates, err := hdl.spotOnlineService.GetDates(spot_uudi, domain.FlagFinalized|domain.FlagOnline|domain.FlagPaused)

	if err != nil {
		log.Println("We found an error while calling servide start \n", err.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message:       "We face an error while starting the spot online",
			InternalError: err.Error(),
		})
		return
	}

	context.JSON(200, spotWitjDates)
}

/*
/spots/online/:spot_uuid/start

Input:

	Path variable: spot uuid
	Body: Date body

Output:

	Date with uudi
*/
func (hdl *HTTPHandler) Start(context *gin.Context) {
	log.Println("HTTPHandler: Start")
	// Path Variables
	spot_uudi, is_spot_uudi_present := context.Params.Get("spot_uuid")
	userId := context.Request.Header.Get("Authorization")

	if len(userId) == 0 {
		log.Println("User id is missinbg in Authorization header")
		context.AbortWithStatusJSON(401, ErrorMessage{
			Message: "We could not found the user",
		})
		return
	}

	// Body
	var date *SpotDateRequest
	if err := context.BindJSON(&date); err != nil {
		log.Println("Spot data is mising in the body")
		log.Println(err.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: "The spot data is missing in body.",
		})
		return
	}

	if !is_spot_uudi_present {
		log.Println("Spot id is mising in the path")
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: "The spot id is missing",
		})
		// return not id sent
		return
	}
	log.Printf("\nHandler: Start \n\tSpot UUID: %s,\n\tDate: %+v", spot_uudi, date)

	// context.JSON(200, date)

	// spot, err := hdl.spotService.GoOnline(spot)
	onlineSpot, err := hdl.spotOnlineService.Start(spot_uudi, userId, date.DurationApproximated, date.MaximunCapacity)

	if err != nil {
		log.Println("We found an error while calling servide start \n", err.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message:       "We face an error while starting the spot online",
			InternalError: err.Error(),
		})
		return
	}

	context.JSON(200, onlineSpot)
}

/*
/spots/online/:spot_uuid/stop

Input:

	Path variable: spot uuid

Output:

	204 -> No content
*/
func (hdl *HTTPHandler) Stop(context *gin.Context) {
	log.Println("HTTPHandler: Stop")
	// Path Variables
	spot_uudi, is_spot_uudi_present := context.Params.Get("spot_uuid")
	userId := context.Request.Header.Get("Authorization")

	if len(userId) == 0 {
		log.Println("User id is missinbg in Authorization header")
		context.AbortWithStatusJSON(401, ErrorMessage{
			Message: "We could not found the user",
		})
		return
	}

	if !is_spot_uudi_present {
		log.Println("Spot id is mising in the path")
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: "The spot id is missing",
		})
		// return not id sent
		return
	}
	log.Printf("\nHandler: Stop \n\tSpot UUID: %s,\n\tuserId: %+v", spot_uudi, userId)

	// context.JSON(200, date)

	// spot, err := hdl.spotService.GoOnline(spot)
	err := hdl.spotOnlineService.Stop(spot_uudi, userId)

	if err != nil {
		log.Println("We found an error while calling servide stop \n", err.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message:       "We face an error while stopping the spot",
			InternalError: err.Error(),
		})
		return
	}

	context.Status(204)
}

/*
/spots/online/:spot_uuid/finalize

Input:

	Path variable: spot uuid

Output:

	204 -> No content
*/
func (hdl *HTTPHandler) Finalize(context *gin.Context) {
	log.Println("HTTPHandler: Finalize")
	// Path Variables
	spot_uudi, is_spot_uudi_present := context.Params.Get("spot_uuid")
	userId := context.Request.Header.Get("Authorization")

	if len(userId) == 0 {
		log.Println("User id is missinbg in Authorization header")
		context.AbortWithStatusJSON(401, ErrorMessage{
			Message: "We could not found the user",
		})
		return
	}

	if !is_spot_uudi_present {
		log.Println("Spot id is mising in the path")
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: "The spot id is missing",
		})
		// return not id sent
		return
	}
	log.Printf("\nHandler: Finalize \n\tSpot UUID: %s,\n\tuserId: %+v", spot_uudi, userId)

	// context.JSON(200, date)

	// spot, err := hdl.spotService.GoOnline(spot)
	err := hdl.spotOnlineService.Finalize(spot_uudi, userId)

	if err != nil {
		log.Println("We found an error while calling servide Finalize \n", err.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message:       "We face an error while finalizing the spot",
			InternalError: err.Error(),
		})
		return
	}

	context.Status(204)
}

/*
/spots/online/:spot_uuid/stop

Input:

	Path variable: spot uuid

Output:

	204 -> No content
*/
func (hdl *HTTPHandler) Resume(context *gin.Context) {
	log.Println("HTTPHandler: Resume")
	// Path Variables
	spot_uudi, is_spot_uudi_present := context.Params.Get("spot_uuid")
	userId := context.Request.Header.Get("Authorization")

	if len(userId) == 0 {
		log.Println("User id is missinbg in Authorization header")
		context.AbortWithStatusJSON(401, ErrorMessage{
			Message: "We could not found the user",
		})
		return
	}

	if !is_spot_uudi_present {
		log.Println("Spot id is mising in the path")
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: "The spot id is missing",
		})
		// return not id sent
		return
	}
	log.Printf("\nHandler: Resume \n\tSpot UUID: %s,\n\tuserId: %+v", spot_uudi, userId)

	// context.JSON(200, date)

	// spot, err := hdl.spotService.GoOnline(spot)
	err := hdl.spotOnlineService.Resume(spot_uudi, userId)

	if err != nil {
		log.Println("We found an error while calling servide Resume \n", err.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message:       "We face an error while resuming the spot",
			InternalError: err.Error(),
		})
		return
	}

	context.Status(204)
}

// func (hdl *HTTPHandler) GetEvent(context *gin.Context) {

// 	id := context.Param("id")

// 	if len(id) == 0 {
// 		context.AbortWithStatusJSON(400, ErrorMessage{
// 			Message: "Missing Id param",
// 		})
// 		return
// 	}

// 	format := hdl.getOuputFormat(context.DefaultQuery("format", "empty"))

// 	log.Printf("\nHandler: GetEvent \n\tId: %s, \n\tFormat: %s", id, string(format))

// 	event, err := hdl.spotService.Get(id, format)

// 	if err != nil {
// 		log.Println(err)
// 		context.AbortWithStatusJSON(400, ErrorMessage{
// 			Message:       "We face an error while fethcing the data",
// 			InternalError: err.Error(),
// 		})
// 		return
// 	}

// 	context.JSON(200, event)
// }

// func (hdl *HTTPHandler) GetEvents(context *gin.Context) {

// 	// Getting data from call
// 	spotIds := SpotsIdsRequest{}
// 	context.BindJSON(&spotIds)

// 	format := hdl.getOuputFormat(context.DefaultQuery("format", "empty"))

// 	log.Printf("\nHandler: GetEvents \n\tSpotIds: %+v, \n\tFormat: %s", spotIds, string(format))

// 	multipleSpots, err := hdl.spotService.GetSpots(spotIds.SpotIds, format)

// 	if err != nil {
// 		log.Println(err)
// 		context.AbortWithStatusJSON(400, ErrorMessage{
// 			Message:       "We face an error while fethcing the data",
// 			InternalError: err.Error(),
// 		})
// 		return
// 	}

// 	context.JSON(200, multipleSpots)
// }

// func (hdl *HTTPHandler) getOuputFormat(query string) ports.OutputFormat {
// 	switch strings.ToUpper(query) {
// 	case string(ports.FULL_FORMAT):
// 		return ports.FULL_FORMAT
// 	case string(ports.SMALL_FORMAT):
// 		return ports.SMALL_FORMAT
// 	}
// 	return ports.FULL_FORMAT
// }

// func (hdl *HTTPHandler) isSpotCorrect(spot domain.Spot) bool {
// 	// TODO -> Could we use json schemas here ?

// 	return true
// }
