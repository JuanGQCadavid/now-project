
/*
	The user should not be online in more tha  two events

	Procedure:
		1. check if the user is already in a Online event		YES

			YES -> Desvincalte it from it						YES

		2. Generate UUID for the spot							YES

		3. make the user Online in the repo						YES

		4. Update the location tree.							YES

	TODO -> Extract the tags from the description

*/

func (s *service) CreateSpot(spot domain.Spot) (domain.Spot, error) {
	// TODO -> Missing refactor from go online to createSpot

	log.Println("GoOnline -> ", fmt.Sprintf("%+v", spot))
	//TODO -> Missing body validation

	// TODO -> I think this is not working as I spect, we need to check this.
	if returnedSpot, returnedError := s.GetSpotByUserId(spot.HostInfo.Id); returnedError == nil {
		if err := s.EndSpot(returnedSpot.EventInfo.UUID); err != nil {
			return domain.Spot{}, err
		}
	}

	uuid := s.uuidGen.New()
	spot.EventInfo.UUID = uuid

	if returnedError := s.createEvent(spot); returnedError != nil {
		return domain.Spot{}, returnedError
	}

	// Check if method contains tags for the event

	if spot.TopicsInfo.PrincipalTopic != "" || spot.TopicsInfo.SecondaryTopics != nil {
		if returnedError := s.createSpotTags(spot); returnedError != nil {
			return domain.Spot{}, returnedError
		}
	}

	return spot, nil
}