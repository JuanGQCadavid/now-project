package httphdl

import (
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

func (hdl *HTTPHandler) GoOnline(context *gin.Context) {
	spot := domain.Spot{}
	context.BindJSON(&spot)

	spot, err := hdl.spotService.GoOnline(spot)
	if err != nil {
		context.AbortWithStatusJSON(500, gin.H{"message": err})
		return
	}

	context.JSON(200, spot)
}
