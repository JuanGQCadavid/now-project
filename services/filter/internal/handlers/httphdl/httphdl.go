package httphdl

import (
	"fmt"
	"strconv"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	service        ports.FilterService
	defaultRadious float64
}

func NewHTTPHandler(service ports.FilterService) *HTTPHandler {
	return &HTTPHandler{
		service:        service,
		defaultRadious: 0.5,
	}
}

func (hdl *HTTPHandler) FilterSpots(context *gin.Context) {

	queryLat, isLatPresent := context.GetQuery("cpLat")
	queryLon, isLonPresent := context.GetQuery("cpLon")
	queryRadious := context.DefaultQuery("radious", fmt.Sprintf("%f", hdl.defaultRadious))

	if !isLatPresent || !isLonPresent {
		context.AbortWithStatusJSON(400, map[string]interface{}{
			"errorMessage": "Missing Query params (cpLat, cpLon)",
		})
		return
	}

	cpLat, errLat := strconv.ParseFloat(queryLat, 64)
	cpLon, errLon := strconv.ParseFloat(queryLon, 64)
	radious, errRad := strconv.ParseFloat(queryRadious, 64)

	if errLat != nil || errLon != nil || errRad != nil {
		context.AbortWithStatusJSON(400, map[string]interface{}{
			"errorMessage": "Bad format on query params",
		})
		return
	}

	response := hdl.service.FilterByProximity(cpLat, cpLon, radious)

	context.JSON(200, response)
}
