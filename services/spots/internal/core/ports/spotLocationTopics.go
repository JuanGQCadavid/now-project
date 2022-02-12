package ports

type SpotActivityTopic interface {
	AppendSpot(spotId string) error
	RemoveSpot(spotId string) error
}
