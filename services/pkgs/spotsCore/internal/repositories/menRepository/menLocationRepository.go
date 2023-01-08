package menRepository

import "github.com/JuanGQCadavid/now-project/services/spotsCore/internal/core/domain"

type MenSpotActivityTopic struct {
	men map[string][]byte
}

func NewMenSpotActivityTopic() *MenSpotActivityTopic {
	return &MenSpotActivityTopic{
		men: map[string][]byte{},
	}
}

func (l *MenSpotActivityTopic) AppendSpot(spot domain.Spot) error {
	return nil
}

func (l *MenSpotActivityTopic) RemoveSpot(spotId string) error {
	return nil
}

func (l *MenSpotActivityTopic) CreateSpotTags(spotId string, principalTag domain.Optional, secondaryTags []string) error {
	return nil
}
