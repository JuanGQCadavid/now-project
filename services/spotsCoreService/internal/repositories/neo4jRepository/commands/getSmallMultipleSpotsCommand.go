package commands

import (
	_ "embed"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
)

type GetSmallMultipleSpotsCommand struct {
	datesIds []string
}

var (
	//go:embed queries/smallBulkFetch.cypher
	smallBulkFetch string
)

func NewGetSmallMultipleSpotsCommand(datesIds []string) *GetSmallMultipleSpotsCommand {
	return &GetSmallMultipleSpotsCommand{
		datesIds: datesIds,
	}
}

func (command *GetSmallMultipleSpotsCommand) Run(tr neo4j.Transaction) (interface{}, error) {
	cyperParams := map[string]interface{}{"datesIds": command.datesIds}

	result, err := tr.Run(smallBulkFetch, cyperParams)

	var spotsToReturn []domain.Spot = []domain.Spot{}

	if err != nil {
		logs.Error.Println("Error at running!", err)
		return &domain.MultipleSpots{}, err
	}

	for result.Next() {
		record := result.Record()
		spot := command.getSpotDataFromResult(record)
		spotsToReturn = append(spotsToReturn, spot)
	}

	return &domain.MultipleSpots{
		Spots: spotsToReturn,
	}, nil

}

func (command *GetSmallMultipleSpotsCommand) getSpotDataFromResult(record *db.Record) domain.Spot {
	// Event
	event_name, _ := record.Get("event_name")
	event_UUID, _ := record.Get("event_UUID")
	event_emoji, _ := record.Get("event_emoji")

	// Place
	place_lon, _ := record.Get("place_lon")
	place_provider_id, _ := record.Get("place_provider_id")
	place_lat, _ := record.Get("place_lat")

	// Tags
	tags_ids, _ := record.Get("tag_tags")
	tags_principals, _ := record.Get("tag_principals")
	tags_principals_array := tags_principals.([]interface{})

	secondary_tag := make([]string, 0, len(tags_principals_array))
	primary_tag := ""

	// Date
	// date_confirmed, isConfirmed := record.Get("date_confirmed")
	date_date, _ := record.Get("date_date")
	date_durationApproximatedInSeconds, _ := record.Get("date_durationApproximatedInSeconds")
	date_startTime, _ := record.Get("date_startTime")
	date_UUID, _ := record.Get("date_UUID")

	for index, tag := range tags_ids.([]interface{}) {
		if tags_principals_array[index].(bool) {
			primary_tag = tag.(string)
		} else {
			secondary_tag = append(secondary_tag, tag.(string))
		}

	}

	logs.Info.Printf("%+v", record)

	return domain.Spot{
		EventInfo: domain.Event{
			Name:  getStringFromInterface(event_name),
			UUID:  getStringFromInterface(event_UUID),
			Emoji: getStringFromInterface(event_emoji),
		},
		PlaceInfo: domain.Place{
			Lat:           getFloat64FromInterface(place_lat),
			Lon:           getFloat64FromInterface(place_lon),
			MapProviderId: getStringFromInterface(place_provider_id),
		},
		TopicsInfo: domain.Topic{
			PrincipalTopic:  primary_tag,
			SecondaryTopics: secondary_tag,
		},
		DateInfo: domain.Date{
			DateTime:                      getStringFromInterface(date_date),
			DurationApproximatedInSeconds: getInt64FromInterface(date_durationApproximatedInSeconds),
			Id:                            getStringFromInterface(date_UUID),
			StartTime:                     getStringFromInterface(date_startTime),
		},
	}

}
