package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/domain"
)

var (
	ErrCreatingConfirmationService = errors.New("err - We face an error while creating the service")
)

type Confirmation interface {
	SendConfirmationRequestOnBatch(payload []domain.Spot, batchSize int8) map[string]error
}
