package ports

type LocationRepository interface {
	AppendSpot(spotId string) error
	RemoveSpot(spotId string) error
}
