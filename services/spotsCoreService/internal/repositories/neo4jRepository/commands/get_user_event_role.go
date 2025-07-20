package commands

import (
	"context"
	_ "embed"

	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/ports"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	//go:embed queries/getUserEventRole.cypher
	getUserEventRoleQuery string
)

type GetUserEventRoleCommand struct {
	ctx     context.Context
	userId  string
	eventId string
}

func NewGetUserEventRoleCommand(ctx context.Context, userId, eventId string) *GetUserEventRoleCommand {
	return &GetUserEventRoleCommand{
		ctx:     ctx,
		userId:  userId,
		eventId: eventId,
	}
}

func (cmd *GetUserEventRoleCommand) Run(tr neo4j.Transaction) (interface{}, error) {

	var (
		logger = log.Ctx(cmd.ctx)
		params = map[string]interface{}{
			"user_id":  cmd.userId,
			"event_id": cmd.eventId,
		}
		result, err = tr.Run(getUserEventRoleQuery, params)
	)

	logger.Debug().Str("Query", getUserEventRoleQuery).Any("Params", params).Msg("Data for call")

	if err != nil {
		logger.Err(err).
			Str("Method", "GetUserEventRoleCommand.Run.tr.Run").
			Str("userId", cmd.userId).
			Str("eventId", cmd.eventId).
			Msg("Error while running the commmand")

		return nil, ports.ErrCallingRepository
	}

	for result.Next() {
		return cmd.recordToAccess(logger, result.Record()), nil
	}

	return nil, nil
}

func (cmd *GetUserEventRoleCommand) recordToAccess(logger *zerolog.Logger, record *neo4j.Record) *domain.Access {
	var (
		userId, _     = record.Get("user_id")
		accessType, _ = record.Get("relation_type")
		userName, _   = record.Get("user_name")
	)

	logger.Debug().Any("Record keys", record.Keys).Any("Record values", record.Values).Send()

	return &domain.Access{
		UserId:   getStringFromInterface(userId),
		UserName: getStringFromInterface(userName),
		Role: domain.AccessRoleFromString(
			getStringFromInterface(accessType),
		),
	}
}
