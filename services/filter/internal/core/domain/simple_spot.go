package domain

type SimpleSpot struct {
	Id        string   `json:"spotId"`
	Type      SpotType `json:"spotType"`
	Emoji     string   `json:"spotEmoji"`
	StartTime string   `json:"spotStartTime"`
	LatLng    LatLng   `json:"latLng"`
}

type SpotType string

const (
	OnlineSpot    SpotType = "online"
	ScheduledSpot SpotType = "schedule"
	Other         SpotType = "other"
)
