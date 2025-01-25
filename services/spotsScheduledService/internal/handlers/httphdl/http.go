package httphdl

import (
	"errors"

	authDomain "github.com/JuanGQCadavid/now-project/services/authService/core/core/domain"
	authUtils "github.com/JuanGQCadavid/now-project/services/authService/core/utils"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	service ports.Service
}

var (
	ErrMissingSpotIdOnParam                   = errors.New("Spot id is required but it is missing on path params")
	ErrMissingScheduleIdOnParam               = errors.New("Schedule id is required but it is missing on path params")
	ErrUserIsNotAllowedToPerformThisOperation = errors.New("The user is not allowed to perfom such uperation")
	ErrEmptySchedulePatterns                  = errors.New("Avoiding process as the schedule patterns is empty")
)

func NewHttpHandler(service ports.Service) *HttpHandler {
	return &HttpHandler{
		service: service,
	}
}

func (hdl *HttpHandler) SetRouter(router *gin.Engine) {
	router.POST("/spots/scheduled/:spot_uuid/append", hdl.AppendSchedule)
	router.GET("/spots/scheduled/:spot_uuid/", hdl.GetSchedule)
	router.GET("/spots/scheduled/:spot_uuid/dates", hdl.GetDates)
	router.PUT("/spots/scheduled/:spot_uuid/scheduled/:scheduled_uuid/resume", hdl.Resume)
	router.PUT("/spots/scheduled/:spot_uuid/scheduled/:scheduled_uuid/freeze", hdl.Freeze)
	router.PUT("/spots/scheduled/:spot_uuid/scheduled/:scheduled_uuid/conclude", hdl.Conclude)
}

/*
GET /spots/schedule/<spot_UUID>/
*/
func (hdl *HttpHandler) GetSchedule(context *gin.Context) {
	spotId := context.Param("spot_uuid")
	userDetails := authUtils.GetHeaders(context.Request.Header)

	if len(spotId) == 0 {
		logs.Error.Println(ErrMissingSpotIdOnParam.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: ErrMissingSpotIdOnParam.Error(),
		})
		return
	}

	spot, err := hdl.service.GetSchedules(spotId, userDetails.UserID, domain.ActivateFlag|domain.ConcludeFlag|domain.FreezeFlag)

	if err != nil {
		logs.Error.Println(err.Error())
		switch err.Error() {
		case ports.ErrOnRepository.Error():
			context.AbortWithStatusJSON(500, ErrorMessage{
				Message: err.Error(),
			})
			return
		case ports.ErrSpotNotFound.Error():
			context.AbortWithStatusJSON(404, ErrorMessage{
				Message: err.Error(),
			})
			return
		default:
			context.AbortWithStatusJSON(500, ErrorMessage{
				Message: err.Error(),
			})
			return
		}
	}
	context.JSON(200, spot)

}

/*
GET /spots/schedule/<spot_UUID>/
*/
func (hdl *HttpHandler) GetDates(context *gin.Context) {
	spotId := context.Param("spot_uuid")
	userDetails := authUtils.GetHeaders(context.Request.Header)

	if len(spotId) == 0 {
		logs.Error.Println(ErrMissingSpotIdOnParam.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: ErrMissingSpotIdOnParam.Error(),
		})
		return
	}

	dates, err := hdl.service.GetDates(spotId, userDetails.UserID)

	if err != nil {
		logs.Error.Println(err.Error())
		switch err.Error() {
		case ports.ErrOnRepository.Error():
			context.AbortWithStatusJSON(500, ErrorMessage{
				Message: err.Error(),
			})
			return
		case ports.ErrSpotNotFound.Error():
			context.AbortWithStatusJSON(404, ErrorMessage{
				Message: err.Error(),
			})
			return
		default:
			context.AbortWithStatusJSON(500, ErrorMessage{
				Message: err.Error(),
			})
			return
		}
	}
	context.JSON(200, dates)

}

/*
POST /spots/schedule/<spot_UUID>/append
*/
func (hdl *HttpHandler) AppendSchedule(context *gin.Context) {
	spotId := context.Param("spot_uuid")
	userDetails := authUtils.GetHeaders(context.Request.Header)

	var scheduleSpotRequest domain.ScheduledSpot = domain.ScheduledSpot{}
	context.BindJSON(&scheduleSpotRequest)

	if len(spotId) == 0 {
		logs.Error.Println(ErrMissingSpotIdOnParam.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: ErrMissingSpotIdOnParam.Error(),
		})
		return
	}

	if userDetails.UserID == authDomain.AnonymousUser.UserID {
		logs.Error.Println(ErrUserIsNotAllowedToPerformThisOperation.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: ErrUserIsNotAllowedToPerformThisOperation.Error(),
		})
		return
	}

	if len(scheduleSpotRequest.Patterns) == 0 {
		logs.Error.Println(ErrEmptySchedulePatterns.Error())
		context.AbortWithStatusJSON(202, ErrorMessage{
			Message: ErrEmptySchedulePatterns.Error(),
		})
		return
	}

	scheduleSpot, timeConflicts, err := hdl.service.AppendSchedule(spotId, userDetails.UserID, scheduleSpotRequest.Patterns)

	if timeConflicts != nil && len(*timeConflicts) > 0 {
		logs.Error.Println(err.Error())
		context.AbortWithStatusJSON(400, TimeErrorsConflictsMessage{
			Message:       err.Error(),
			TimeConflicts: *timeConflicts,
		})

		return
	} else if err != nil {
		logs.Error.Println(err.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: err.Error()})
		return
	}

	context.JSON(202, scheduleSpot)

}

/*
PUT /spots/schedule/<spot_UUID>/sheduled/<scheduled_uuid>/resume
*/
func (hdl *HttpHandler) Resume(context *gin.Context) {
	spotId := context.Param("spot_uuid")
	scheduleId := context.Param("scheduled_uuid")
	userDetails := authUtils.GetHeaders(context.Request.Header)

	if len(spotId) == 0 {
		logs.Error.Println(ErrMissingSpotIdOnParam.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: ErrMissingSpotIdOnParam.Error(),
		})
		return
	} else if len(scheduleId) == 0 {
		logs.Error.Println(ErrMissingScheduleIdOnParam.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: ErrMissingScheduleIdOnParam.Error(),
		})
		return
	} else if userDetails.UserID == authDomain.AnonymousUser.UserID {
		logs.Error.Println(ErrUserIsNotAllowedToPerformThisOperation.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: ErrUserIsNotAllowedToPerformThisOperation.Error(),
		})
		return
	}

	err := hdl.service.ResumeSchedule(spotId, scheduleId, userDetails.UserID)

	if err != nil {
		logs.Error.Println(err.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: err.Error(),
		})

		return
	}

	context.Status(204)
}

/*
PUT /spots/schedule/<spot_UUID>/sheduled/<scheduled_uuid>/freeze
*/
func (hdl *HttpHandler) Freeze(context *gin.Context) {
	spotId := context.Param("spot_uuid")
	scheduleId := context.Param("scheduled_uuid")
	userDetails := authUtils.GetHeaders(context.Request.Header)

	if len(spotId) == 0 {
		logs.Error.Println(ErrMissingSpotIdOnParam.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: ErrMissingSpotIdOnParam.Error(),
		})
		return
	} else if len(scheduleId) == 0 {
		logs.Error.Println(ErrMissingScheduleIdOnParam.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: ErrMissingScheduleIdOnParam.Error(),
		})
		return
	} else if userDetails.UserID == authDomain.AnonymousUser.UserID {
		logs.Error.Println(ErrUserIsNotAllowedToPerformThisOperation.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: ErrUserIsNotAllowedToPerformThisOperation.Error(),
		})
		return
	}

	err := hdl.service.FreezeSchedule(spotId, scheduleId, userDetails.UserID)

	if err != nil {
		logs.Error.Println(err.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: err.Error(),
		})

		return
	}

	context.Status(204)
}

/*
PUT /spots/schedule/<spot_UUID>/sheduled/<scheduled_uuid>/conclude
*/
func (hdl *HttpHandler) Conclude(context *gin.Context) {
	spotId := context.Param("spot_uuid")
	scheduleId := context.Param("scheduled_uuid")
	userDetails := authUtils.GetHeaders(context.Request.Header)

	if len(spotId) == 0 {
		logs.Error.Println(ErrMissingSpotIdOnParam.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: ErrMissingSpotIdOnParam.Error(),
		})
		return
	} else if len(scheduleId) == 0 {
		logs.Error.Println(ErrMissingScheduleIdOnParam.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: ErrMissingScheduleIdOnParam.Error(),
		})
		return
	} else if userDetails.UserID == authDomain.AnonymousUser.UserID {
		logs.Error.Println(ErrUserIsNotAllowedToPerformThisOperation.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: ErrUserIsNotAllowedToPerformThisOperation.Error(),
		})
		return
	}

	err := hdl.service.ConcludeSchedule(spotId, scheduleId, userDetails.UserID)

	if err != nil {
		logs.Error.Println(err.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: err.Error(),
		})

		return
	}

	context.Status(204)
}
