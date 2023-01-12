package spotsrv

import (
	"fmt"
	"log"
	"strings"

	"github.com/JuanGQCadavid/now-project/services/spotsCore/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsCore/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/spotsCore/pkg/uuidgen"
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
	return s.spotRepository.Get(spotId, format)
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

	uuid := s.uuidGen.New()
	spot.EventInfo.UUID = uuid

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

// TODO -> Fix this!
func (s *service) FinalizeSpot(spotId string) error {
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

// TODO Implement
func (s *service) UpdateSpot() error {
	return nil
}
