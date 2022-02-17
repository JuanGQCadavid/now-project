package locationrepositories

type SpotLocation struct {
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
	SpotId string  `json:"spotId"`
}
