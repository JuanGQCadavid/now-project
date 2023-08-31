package dbspotservicelambda

import (
	"time"

	"gorm.io/gorm"
)

type SpotsDB struct {
	EventId                           string `gorm:"primaryKey"`
	DateId                            string `gorm:"index"`
	DateDateTime                      string
	DateDurationApproximatedInSeconds int64
	DateStartTime                     string
	EventName                         string
	EventDescription                  string
	EventMaximunCapacty               int64
	EventEmoji                        string
	PersonId                          string
	PersonName                        string
	PlaceName                         string
	PlaceLat                          float64
	PlaceLon                          float64
	PlaceMapProviderId                string
	TopicPrincipalTopic               string
	TopicSecondaryTopics              string

	// GORM Variables
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
