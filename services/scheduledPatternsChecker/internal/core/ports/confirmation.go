package ports

import "github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/domain"

type Confirmation interface {
	SendConfirmationRequestOnBatch(payload []domain.Spot, batchSize int8) error
}
