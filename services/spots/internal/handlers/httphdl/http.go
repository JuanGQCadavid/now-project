package httphdl

import "github.com/JuanGQCadavid/now-project/services/spots/internal/core/ports"

type HTTPHandler struct {
	spotService ports.SpotService
}

func NewHTTPHandler(spotService ports.SpotService) *HTTPHandler {
	return &HTTPHandler{
		spotService: spotService,
	}
}
