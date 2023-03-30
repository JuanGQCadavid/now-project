package httphdl

import "github.com/gin-gonic/gin"

type HttpHandler struct {
}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (hdl *HttpHandler) SetRouter(router *gin.Engine) {
	router.POST("/spots/schedule/:spot_uuid/append", hdl.AppendSchedule)
	router.GET("/spots/schedule/:spot_uuid/", hdl.GetSchedule)
	router.PUT("/spots/schedule/:spot_uuid/scheduled/:scheduled_uuid/resume", hdl.Resume)
	router.PUT("/spots/schedule/:spot_uuid/scheduled/:scheduled_uuid/freeze", hdl.Freeze)
	router.PUT("/spots/schedule/:spot_uuid/scheduled/:scheduled_uuid/conclude", hdl.Conclude)

}

/*
GET /spots/schedule/<spot_UUID>/
*/
func (hdl *HttpHandler) GetSchedule(context *gin.Context) {

}

/*
POST /spots/schedule/<spot_UUID>/append
*/
func (hdl *HttpHandler) AppendSchedule(context *gin.Context) {

}

/*
PUT /spots/schedule/<spot_UUID>/sheduled/<scheduled_uuid>/resume
*/
func (hdl *HttpHandler) Resume(context *gin.Context) {

}

/*
PUT /spots/schedule/<spot_UUID>/sheduled/<scheduled_uuid>/freeze
*/
func (hdl *HttpHandler) Freeze(context *gin.Context) {

}

/*
PUT /spots/schedule/<spot_UUID>/sheduled/<scheduled_uuid>/conclude
*/
func (hdl *HttpHandler) Conclude(context *gin.Context) {

}
