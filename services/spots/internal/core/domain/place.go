package domain

type Place struct {
	Name          string  `json:"name"`
	Lat           float32 `json:"lat"` // those nas are tags!
	Lon           float32 `json:"lon"`
	MapProviderId int32   `json:"mapProviderId"`
}
