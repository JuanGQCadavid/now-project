package queue

import (
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/sqsqueue"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/ports"
)

type SQSConfirmation struct {
	sqsAction *sqsqueue.SQSQueueActions
}

func NewSQSConfirmationFromEnv(sqsName string) (*SQSConfirmation, error) {

	sqsAction, err := sqsqueue.NewSQSQueueActionsFromEnv(sqsName)

	if err != nil {
		logs.Error.Println("We were unable to create the confirmation queue")
		return nil, ports.ErrCreatingConfirmationService
	}

	return &SQSConfirmation{
		sqsAction: sqsAction,
	}, nil
}

func (srv *SQSConfirmation) SendConfirmationRequestOnBatch(payload []domain.Spot, batchSize int8) map[string]error {
	logs.Info.Println("Sending data: ", payload)

	dataAsInterface := make([]interface{}, len(payload))

	for i, data := range payload {
		dataAsInterface[i] = data
	}

	errors := srv.sqsAction.SendBulkMessages(dataAsInterface)

	if errors != nil {
		sendErrors := make(map[string]error)
		logs.Error.Println("Some messages fail to be processed")

		for spotI, err := range errors {
			spot := spotI.(domain.Spot)
			sendErrors[spot.SpotId] = err
		}

		return sendErrors
	}

	return nil
}
