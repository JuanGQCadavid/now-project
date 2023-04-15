package domain

type SpotInfo struct {
	SpotId  string `json:"spotId,omitempty"`
	OwnerId string `json:"-"`
}
