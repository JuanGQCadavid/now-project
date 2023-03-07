package domain

type Spot struct {
	SpotId  string `json:"SpotId"`
	OwnerId string `json:"-"`
}
