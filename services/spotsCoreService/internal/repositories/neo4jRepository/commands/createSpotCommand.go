package commands

import (
	_ "embed"

	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type CreateSpotCommand struct {
	Spot *domain.Spot
}

func NewCreateSpotCommand(spot *domain.Spot) *CreateSpotCommand {
	return &CreateSpotCommand{
		Spot: spot,
	}
}

var (
	//go:embed queries/createSport.cypher
	createSpotCypher string
)

func (cmd *CreateSpotCommand) Run(tr neo4j.Transaction) (interface{}, error) {
	cypherParams := map[string]interface{}{
		"event_uuid":         cmd.Spot.EventInfo.UUID,
		"event_desc":         cmd.Spot.EventInfo.Description,
		"event_max_capacity": cmd.Spot.EventInfo.MaximunCapacty,
		"event_name":         cmd.Spot.EventInfo.Name,
		"event_emoji":        cmd.Spot.EventInfo.Emoji,
		"place_provider_id":  cmd.Spot.PlaceInfo.MapProviderId,
		"place_lat":          cmd.Spot.PlaceInfo.Lat,
		"place_lon":          cmd.Spot.PlaceInfo.Lon,
		"place_name":         cmd.Spot.PlaceInfo.Name,
		"host_phone_number":  cmd.Spot.HostInfo.PhoneNumber,
		"host_name":          cmd.Spot.HostInfo.Name,
		"host_id":            cmd.Spot.HostInfo.Id,
	}

	return tr.Run(createSpotCypher, cypherParams)
}
