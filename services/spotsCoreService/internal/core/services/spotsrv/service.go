package spotsrv

import (
	"fmt"
	"log"
	"strings"

	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/pkg/uuidgen"
)

/*
When implementin a interface is just creating a struct with all
the methods that the interface defined.
*/
type service struct {
	spotRepository    ports.SpotRepository
	spotActivityTopic ports.SpotActivityTopic
	uuidGen           uuidgen.UUIDGen
}

func New(spotRepository ports.SpotRepository, spotActivityTopic ports.SpotActivityTopic, uuidGen uuidgen.UUIDGen) *service {
	return &service{
		spotRepository:    spotRepository,
		spotActivityTopic: spotActivityTopic,
		uuidGen:           uuidGen,
	}
}

func (s *service) Get(spotId string, format ports.OutputFormat) (domain.Spot, error) {
	log.Println("Service: Get ->", spotId)
	spotFounded, err := s.spotRepository.Get(spotId, format)

	if err != nil {
		log.Println("We get an error while getting the spot with id", spotId)
		return domain.Spot{}, err
	}

	if len(spotFounded.EventInfo.UUID) == 0 {
		log.Println("Spot not founded with id", spotId)
		return domain.Spot{}, ports.ErrSpotNotFounded
	}

	return spotFounded, nil
}

func (s *service) GetSpots(spotIds []string, format ports.OutputFormat) (domain.MultipleSpots, error) {
	// TODO -> Should we check that all elements are not empy ?
	log.Println("Service: GetSpots ->", fmt.Sprintf("%+v", spotIds))
	return s.spotRepository.GetSpots(spotIds, format)
}

/*
	The user should not be online in more tha  two events

	Procedure:

		1. Generate UUID for the spot							YES

		3. make the user Online in the repo						YES

		4. Update the location tree.							YES

	TODO -> Extract the tags from the description

*/

func (s *service) CreateSpot(spot domain.Spot) (domain.Spot, error) {

	log.Println("Service: CreateSpot -> ", fmt.Sprintf("%+v", spot))
	//TODO -> Missing body validation

	spotUuid := s.uuidGen.New()

	if len(spot.HostInfo.Id) == 0 {
		hostUuid := s.uuidGen.New()
		spot.HostInfo.Id = hostUuid
	}

	spot.EventInfo.UUID = spotUuid

	if returnedError := s.spotRepository.CreateSpot(spot); returnedError != nil {
		return domain.Spot{}, returnedError
	}

	if returnedError := s.spotActivityTopic.NotifySpotCreated(spot); returnedError != nil {
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

func (s *service) createSpotTags(spot domain.Spot) error {

	var principalTag *domain.Optional = s.sanitizeTag(spot.TopicsInfo.PrincipalTopic)
	var secondaryTopics []string = s.sanitizeTags(spot.TopicsInfo.SecondaryTopics...)

	return s.spotRepository.CreateSpotTags(spot.EventInfo.UUID, *principalTag, secondaryTopics)
}

func (s *service) sanitizeTags(tags ...string) []string {
	var response []string

	for _, tag := range tags {
		newTagOptional := s.sanitizeTag(tag)

		if newTagOptional.IsPresent() {
			response = append(response, newTagOptional.GetValue())
		}
	}

	return response
}

func (s *service) sanitizeTag(tag string) *domain.Optional {
	var newTag string = strings.ToLower(tag)
	newTag = strings.ReplaceAll(newTag, " ", "")

	return domain.NewOptional(
		newTag,
	)
}

func (s *service) GetSpotByUserId(userId string) (domain.Spot, error) {
	spot, returnedError := s.spotRepository.GetSpotByUserId(userId)

	// TODO -> Improve this.
	if returnedError != nil {
		log.Println("Error on GetSpotByUserId: ", returnedError)
		return domain.Spot{}, returnedError
	}

	return spot, nil
}

func (s *service) DeleteSpot(spotId string, requestUserId string) error {
	log.Printf("Service - DeleteSpot: Id: %s, requestUserId: %s \n", spotId, requestUserId)

	// 1. Fetch the event that the spot id belongs to.
	originalSpot, err := s.spotRepository.Get(spotId, ports.FULL_FORMAT)

	if err != nil {
		log.Println("ERROR: Service - DeleteSpot - Fetch actual spot fail: ", err.Error())
		return err
	}

	if len(originalSpot.EventInfo.UUID) == 0 {
		log.Println("Spot not founded, it is empty")
		return ports.ErrSpotNotFounded
	}

	// 2. Verify that the owner id is the same as the one that is making the request
	if originalSpot.HostInfo.Id != requestUserId {
		log.Println("HostInfo", originalSpot.HostInfo.Id)
		log.Println("requestUserId", requestUserId)
		return ports.ErrSpotUserNotOwnerWhenUpdatingSpot
	}

	// 3. Add the Delete state to the spot

	err = s.spotRepository.DeleteSpot(spotId)

	if err != nil {
		log.Println("ERROR: Service - DeleteSpot - Delete command fail: ", err.Error())
		return err
	}

	return nil
}

func (s *service) FinalizeSpot(spotId string, requestUserId string) error {
	/*
		if returnedError := s.spotRepository.EndSpot(spotId); returnedError != nil {
			log.Println("Error on EndSpot: ", returnedError)
			return returnedError
		}

		if returnedError := s.spotActivityTopic.RemoveSpot(spotId); returnedError != nil {
			log.Println("Error on EndSpot: ", returnedError)
			return returnedError
		}

		return nil
	*/
	return nil
}

/*
Only users with owner relationship could update an spot.

Body that could be changed:

	{
		"eventInfo": {
			"name": "...",
			"description": "...",
			"maximunCapacty": ##,
			"emoji": "..."
		}
	}

Procedure:

 1. Fetch the event that the spot id belongs to.

 2. Verify that the owner id is the same as the one that is making the request

 3. Verify that the data is diffent
    TRUE -> Update Event
    FALSE -> Just return
*/
func (s *service) UpdateSpotEvent(spotId string, ownerId string, spotEvent *domain.Event) error {
	log.Printf("Service - UpdateSpotEvent: Id: %s, SpotEvent: %+v \n", spotId, spotEvent)

	// 1. Fetch the event that the spot id belongs to.
	originalSpot, err := s.spotRepository.Get(spotId, ports.FULL_FORMAT)

	if err != nil {
		log.Println("ERROR: Service - UpdateSpotEvent: ", err.Error())
		return err
	}

	// 2. Verify that the owner id is the same as the one that is making the request
	if originalSpot.HostInfo.Id != ownerId {
		return ports.ErrSpotUserNotOwnerWhenUpdatingSpot
	}

	//  3. Verify that the data is diffent
	spotEvent.UUID = spotId
	if originalSpot.EventInfo.IsEquals(spotEvent) {
		return ports.ErrSpotToUpdateIsTheSameAsTheDb
	}

	// Updated the event
	return s.spotRepository.UpdateSpotEvent(*spotEvent, spotId)
}

/*
Only users with owner relationship could update an spot.

Body that could be changed:

	{
		"placeInfo": {
			"name": "...",
			"lat": ###,
			"lon": ###,
			"mapProviderId": "..."
		}
	}
*/
func (s *service) UpdateSpotPlace(spotId string, ownerId string, spotEvent *domain.Place) error {
	return nil
}

/*
Only users with owner relationship could update an spot.

Body that could be changed:

	{
		"topicInfo": {
			"principalTopic": "...",
			"secondaryTopics" : ["...", "..."]
		}
	}
*/
func (s *service) UpdateSpotTopic(spotId string, ownerId string, spotEvent *domain.Topic) error {
	return nil
}
