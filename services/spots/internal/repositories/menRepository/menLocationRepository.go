package menRepository

type MenSpotActivityTopic struct {
	men map[string][]byte
}

func NewMenSpotActivityTopic() *MenSpotActivityTopic {
	return &MenSpotActivityTopic{
		men: map[string][]byte{},
	}
}

func (l *MenSpotActivityTopic) AppendSpot(spotId string) error {
	return nil
}

func (l *MenSpotActivityTopic) RemoveSpot(spotId string) error {
	return nil
}
