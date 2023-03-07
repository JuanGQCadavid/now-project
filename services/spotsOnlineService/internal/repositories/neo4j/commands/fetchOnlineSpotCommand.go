package commands

import (
	_ "embed"
	"log"

	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/core/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
)

type FetchSpotCommand struct {
	spotUUID string
}

func NewFetchSpotCommand(spotUUID string) *FetchSpotCommand {
	return &FetchSpotCommand{
		spotUUID: spotUUID,
	}
}

var (
	//go:embed queries/fetch_online.cypher
	cypherQuery string
)

func (cmd *FetchSpotCommand) Run(tr neo4j.Transaction) (interface{}, error) {
	log.Printf("FetchSpotCommand - Run")

	cypherParams := map[string]interface{}{
		"spot_uuid": cmd.spotUUID,
	}

	log.Println(cypherParams, cypherQuery)

	result, err := tr.Run(cypherQuery, cypherParams)

	if err != nil {
		println("Error at running!", err)
		return domain.OnlineSpot{}, err
	}

	var onlineSpot *domain.OnlineSpot = &domain.OnlineSpot{}

	for result.Next() {
		onlineSpot = cmd.getSpotDataFromResult(result.Record())
	}

	return *onlineSpot, nil
}

func (command *FetchSpotCommand) getSpotDataFromResult(record *db.Record) *domain.OnlineSpot {
	log.Printf("getSpotDataFromResult -> \n\t%+v", record)

	// Event
	event_UUID, _ := record.Get("event_UUID")

	// Place
	place_name, _ := record.Get("place_name")
	place_lon, _ := record.Get("place_lon")
	place_provider_id, _ := record.Get("place_provider_id")
	place_lat, _ := record.Get("place_lat")

	// Host
	host_id, _ := record.Get("host_id")

	// Dates
	date_online, _ := record.Get("dates_online")
	dates := make([]domain.SpotDate, 0)
	if date_online != nil {
		date_online_array := date_online.([]map[string]interface{})
		dates := make([]domain.SpotDate, len(date_online_array))

		for i, date := range date_online_array {
			date_uuid, _ := date["date_uuid"].(string)
			date_duration_in_seconds, _ := date["date_duration_in_seconds"].(int64)
			date_start_time, _ := date["date_start_time"].(string)
			date_date, _ := date["date_date"].(string)
			date_confirmed, _ := date["date_confirmed"].(bool)
			date_maximun_capacity, _ := date["date_maximun_capacity"].(int)
			hosted_by, _ := date["hosted_by"].(map[string]string)

			dates[i] = domain.SpotDate{
				DateId:                        date_uuid,
				DurationApproximatedInSeconds: date_duration_in_seconds,
				StartTime:                     date_start_time,
				Date:                          date_date,
				Confirmed:                     date_confirmed,
				MaximunCapacty:                date_maximun_capacity,
				HostInfo: domain.HostInfo{
					HostId:   hosted_by["host_id"],
					HostName: hosted_by["host_name"],
				},
			}
		}
	}

	return &domain.OnlineSpot{
		SpotInfo: domain.Spot{
			SpotId:  getStringFromInterface(event_UUID),
			OwnerId: getStringFromInterface(host_id),
		},
		DatesInfo: dates,
		PlaceInfo: domain.Place{
			Name:          getStringFromInterface(place_name),
			Lat:           getFloat64FromInterface(place_lat),
			Lon:           getFloat64FromInterface(place_lon),
			MapProviderId: getStringFromInterface(place_provider_id),
		},
	}
}
