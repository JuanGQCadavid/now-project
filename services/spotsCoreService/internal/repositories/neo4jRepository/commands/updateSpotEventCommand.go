package commands

import (
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
)

type UpdateSpotEventCommand struct {
	event  *domain.Event
	spotId string
}

func NewUpdateSpotEventCommand(event *domain.Event, spotId string) *UpdateSpotEventCommand {
	return &UpdateSpotEventCommand{
		event:  event,
		spotId: spotId,
	}
}

func (cmd *UpdateSpotEventCommand) Run(trx neo4j.Transaction) (interface{}, error) {
	var cypherQ string = `
		Match (spot:Event {UUID:$spotId})
		SET spot.description=$spotDescription
		SET spot.emoji=$spotEmoji
		SET spot.maximunCapacty=$spotCapacity
		SET spot.name=$spotName
		RETURN spot.description as spot_description,
			spot.emoji as spot_emoji,
			spot.maximunCapacty as spot_maximunCapacty,
			spot.name as spot_name,
			spot.UUID as spot_UUID
	`
	var params map[string]interface{} = map[string]interface{}{
		"spotId":          cmd.spotId,
		"spotDescription": cmd.event.Description,
		"spotEmoji":       cmd.event.Emoji,
		"spotCapacity":    cmd.event.MaximunCapacty,
		"spotName":        cmd.event.Name,
	}

	result, err := trx.Run(cypherQ, params)

	if err != nil {
		logs.Error.Println("commands - UpdateSpotEventCommand - Run: We face an error when updating the spot event, err: ", err.Error())
		return nil, err
	}

	var record *db.Record
	var event *domain.Event
	for result.NextRecord(&record) {
		event = cmd.mapRecord(record)
		break
	}
	return event, nil
}

func (cmd *UpdateSpotEventCommand) mapRecord(record *db.Record) *domain.Event {
	spot_description, _ := record.Get("spot_description")
	spot_emoji, _ := record.Get("spot_emoji")
	spot_maximunCapacty, _ := record.Get("spot_maximunCapacty")
	spot_name, _ := record.Get("spot_name")
	spot_UUID, _ := record.Get("spot_UUID")

	return &domain.Event{
		Description:    getStringFromInterface(spot_description),
		Name:           getStringFromInterface(spot_name),
		UUID:           getStringFromInterface(spot_UUID),
		Emoji:          getStringFromInterface(spot_emoji),
		MaximunCapacty: getInt64FromInterface(spot_maximunCapacty),
	}
}
