package localconfirmation

import (
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/domain"
)

type LocalConfirmation struct {
}

func NewLocalConfirmation() *LocalConfirmation {
	return &LocalConfirmation{}
}

func (l *LocalConfirmation) SendConfirmationRequestOnBatch(payload []domain.Spot, batchSize int8) map[string]error {
	logs.Info.Println("Sending data: ", payload)
	return nil
}
