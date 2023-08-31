package dbspotservicelambda

import (
	"fmt"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"gorm.io/gorm"
)

type DBSpotServiceLambda struct {
	db *gorm.DB
}

func NewDBSpotServiceLambdaWithDriver(db *gorm.DB) (*DBSpotServiceLambda, error) {
	return &DBSpotServiceLambda{
		db: db,
	}, nil
}

const (
	SPOT_URL = "spotServiceURL"
)

func (srv *DBSpotServiceLambda) GetSpotsCardsInfo(datesIds []string, format ports.OutputFormat) ([]domain.Spot, error) {
	logs.Info.Println("GetSpotsCardsInfo. Params:", fmt.Sprintf("datesIds: %+v, format: %+v", datesIds, format))

	var results []SpotsDB

	result := srv.db.Where("event_id IN ?", datesIds).Find(&results)

	if result.Error != nil {
		logs.Error.Println("[ERROR] GetSpotsCardsInfo - An error occoured while runnning Query, err: ", result.Error.Error())
		return nil, ports.ErrQueringData
	}

	toReturn := make([]domain.Spot, len(results))

	for i, toMap := range results {
		toReturn[i] = domain.Spot{
			DateInfo: domain.Date{
				DateTime:                      toMap.DateDateTime,
				Id:                            toMap.DateId,
				DurationApproximatedInSeconds: toMap.DateDurationApproximatedInSeconds,
				StartTime:                     toMap.DateStartTime,
			},
			EventInfo: domain.Event{
				Name:        toMap.EventName,
				UUID:        toMap.EventId,
				Description: toMap.EventDescription,
				Emoji:       toMap.EventEmoji,
			},
			HostInfo: domain.Person{
				Id:   toMap.PersonId,
				Name: toMap.PersonName,
			},
			PlaceInfo: domain.Place{
				Name:          toMap.PlaceName,
				Lat:           toMap.PlaceLat,
				Lon:           toMap.PlaceLon,
				MapProviderId: toMap.PlaceMapProviderId,
			},
			TopicsInfo: domain.Topic{
				PrincipalTopic:  toMap.TopicPrincipalTopic,
				SecondaryTopics: []string{toMap.TopicSecondaryTopics},
			},
		}
	}

	return toReturn, nil
}
