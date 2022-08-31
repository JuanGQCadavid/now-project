package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"
)

var (
	ErrBodyRequestUnmarshal  = errors.New("Body request unmarshal fail")
	ErrBodyResponseUnmarshal = errors.New("Body response unmarshal fail")
	ErrBodyResponseReadFail  = errors.New("Body response unable to read")
	ErrSendingRequest        = errors.New("The service is not able to send the request")
)

type SpotService interface {
	GetSpotsCardsInfo(spots []string, format OutputFormat) ([]domain.Spot, error)
}
