package menRepository

type LocationRepository struct {
	men map[string][]byte
}

func NewLocationRepository() *LocationRepository {
	return &LocationRepository{
		men: map[string][]byte{},
	}
}

func (l *LocationRepository) AppendSpot(spotId string) error {
	return nil
}

func (l *LocationRepository) RemoveSpot(spotId string) error {
	return nil
}
