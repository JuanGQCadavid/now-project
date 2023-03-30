package domain

type SpotInfo struct {
	SpotId  string `json:"spotId"`
	OwnerId string `json:"-"`
}
