package models

import "github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"

type Spot struct {
	Id        string          `json:"spotId"`
	Type      domain.SpotType `json:"spotType"`
	Emoji     string          `json:"spotEmoji"`
	StartTime string          `json:"spotStartTime"`
	LatLng    domain.LatLng   `json:"latLng"`
}
