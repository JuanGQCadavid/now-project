package domain

type Place struct {
	Name          string  `json:"name,omitempty"`
	Lat           float64 `json:"lat,omitempty"`
	Lon           float64 `json:"lon,omitempty"`
	MapProviderId string  `json:"mapProviderId,omitempty"`
}
