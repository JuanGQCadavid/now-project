package neo4jRepository

type AWSSpotActivityTopic struct {
}

func NewAWSSpotActivityTopic() *AWSSpotActivityTopic {
	return &AWSSpotActivityTopic{}
}
func (r AWSSpotActivityTopic) AppendSpot(spotId string) error {
	return nil
}
func (r AWSSpotActivityTopic) RemoveSpot(spotId string) error {
	return nil
}
