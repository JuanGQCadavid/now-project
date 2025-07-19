package commands

import (
	"context"
	_ "embed"

	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/ports"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/zerolog/log"
)

var (
	// go:embed queries/getDateAttendantsWithRole.cypher
	getDateAttendantsWithRoleQuery string
)

type GetDateAttendantsWithRoleCommand struct {
	ctx     context.Context
	dateId  string
	eventId string
}

func NewGetDateAttendantsWithRoleCommand(ctx context.Context, dateId, eventId string) *GetDateAttendantsWithRoleCommand {
	return &GetDateAttendantsWithRoleCommand{
		ctx:     ctx,
		dateId:  dateId,
		eventId: eventId,
	}
}

func (cmd *GetDateAttendantsWithRoleCommand) Run(tr neo4j.Transaction) (interface{}, error) {

	var (
		logger      = log.Ctx(cmd.ctx)
		result, err = tr.Run(getDateAttendantsWithRoleQuery, map[string]interface{}{
			"date_id":  cmd.dateId,
			"event_id": cmd.eventId,
		})
	)

	if err != nil {
		logger.Err(err).
			Str("Method", "GetDateAttendantsWithRoleCommand.Run.tr.Run").
			Str("dateId", cmd.dateId).
			Str("eventId", cmd.eventId).
			Msg("Error while running the commmand")

		return nil, ports.ErrCallingRepository
	}

	records, err := result.Collect()

	if err != nil {
		logger.Err(err).
			Str("Method", "GetDateAttendantsWithRoleCommand.Run.result.Collect()").
			Str("dateId", cmd.dateId).
			Str("eventId", cmd.eventId).
			Msg("Error while  collecting the results")

		return nil, ports.ErrCallingRepository
	}

	if len(records) == 0 {
		return []*domain.Access{}, nil
	}

	var (
		userAccess = make([]*domain.Access, len(records))
	)

	for i, record := range records {
		userAccess[i] = cmd.recordToAccess(record)
	}

	return userAccess, nil
}

func (cmd *GetDateAttendantsWithRoleCommand) recordToAccess(record *neo4j.Record) *domain.Access {
	var (
		userId, _     = record.Get("user_id")
		accessType, _ = record.Get("access_type")
		userName, _   = record.Get("user_name")
	)

	return &domain.Access{
		UserId:   getStringFromInterface(userId),
		UserName: getStringFromInterface(userName),
		Role: domain.AccessRoleFromString(
			getStringFromInterface(accessType),
		),
	}
}
