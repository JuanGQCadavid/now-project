package locationrepositories

type SpotLocation struct {
	Lat    float32 `json:"lat"`
	Lng    float32 `json:"lng"`
	SpotId string  `json:"spotId"`
}
