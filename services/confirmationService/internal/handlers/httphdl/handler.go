package httphdl

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/confirmationService/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	service ports.Service
}

var (
	ErrMissingDateIdOnParam                   = errors.New("Date id is required but it is missing on path params")
	ErrUserIsNotAllowedToPerformThisOperation = errors.New("The user is not allowed to perfom such uperation")
)

func NewHttpHandler(service ports.Service) *HttpHandler {
	return &HttpHandler{
		service: service,
	}
}

func (hdl *HttpHandler) SetRouter(router *gin.Engine) {
	router.PUT("/confirmation/date/:date_uuid/confirm", hdl.ConfirmDate)
	router.PUT("/confirmation/date/:date_uuid/unconfirm", hdl.UnconfirmDate)
}

/*
PUT /confirmation/date/:date_uuid/confirm
*/
func (hdl *HttpHandler) ConfirmDate(context *gin.Context) {
	dateId := context.Param("date_uuid")
	requesterId := context.Request.Header.Get("Authorization")

	if len(dateId) == 0 {
		logs.Error.Println(ErrMissingDateIdOnParam.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: ErrMissingDateIdOnParam.Error(),
		})
		return
	}

	if len(requesterId) == 0 {
		logs.Error.Println(ErrUserIsNotAllowedToPerformThisOperation.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: ErrUserIsNotAllowedToPerformThisOperation.Error(),
		})
		return
	}

	err := hdl.service.ConfirmDate(dateId, requesterId)

	if err != nil {
		logs.Error.Println(err.Error())
		switch err.Error() {
		case ports.ErrEmptyDate.Error():
			context.AbortWithStatusJSON(500, ErrorMessage{
				Message: err.Error(),
			})
			return
		case ports.ErrEmptyDate.Error():
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

	context.Status(204)
}

/*
PUT /confirmation/date/:date_uuid/unconfirm
*/
func (hdl *HttpHandler) UnconfirmDate(context *gin.Context) {
	dateId := context.Param("date_uuid")
	requesterId := context.Request.Header.Get("Authorization")

	if len(dateId) == 0 {
		logs.Error.Println(ErrMissingDateIdOnParam.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: ErrMissingDateIdOnParam.Error(),
		})
		return
	}

	if len(requesterId) == 0 {
		logs.Error.Println(ErrUserIsNotAllowedToPerformThisOperation.Error())
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: ErrUserIsNotAllowedToPerformThisOperation.Error(),
		})
		return
	}

	err := hdl.service.UnconfirmDate(dateId, requesterId)

	if err != nil {
		logs.Error.Println(err.Error())
		switch err.Error() {
		case ports.ErrEmptyDate.Error():
			context.AbortWithStatusJSON(500, ErrorMessage{
				Message: err.Error(),
			})
			return
		case ports.ErrEmptyDate.Error():
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

	context.Status(204)

}
