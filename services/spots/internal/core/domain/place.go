package domain

type Place struct {
	Name          string  `json:"name"`
	Lat           float64 `json:"lat"` // those nas are tags!
	Lon           float64 `json:"lon"`
	MapProviderId string  `json:"mapProviderId"`
}
