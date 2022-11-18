package spotsrv

import (
	"fmt"
	"log"
	"strings"

	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/pkg/apperrors"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/pkg/uuidgen"
)

/*
	When implementin a interface is just creating a struct with all
	the methods that the interface defined.
*/
type service struct {
	spotRepository     ports.SpotRepository
	locationRepository ports.SpotActivityTopic
	uuidGen            uuidgen.UUIDGen
}

func New(spotRepository ports.SpotRepository, locationRepository ports.SpotActivityTopic, uuidGen uuidgen.UUIDGen) *service {
	return &service{
		spotRepository:     spotRepository,
		locationRepository: locationRepository,
		uuidGen:            uuidGen,
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
		1. check if the user is already in a Online event		YES

			YES -> Desvincalte it from it						YES

		2. Generate UUID for the spot							YES

		3. make the user Online in the repo						YES

		4. Update the location tree.							YES

	TODO -> Extract the tags from the description

*/

func (s *service) GoOnline(spot domain.Spot) (domain.Spot, error) {
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

func (s *service) createEvent(spot domain.Spot) error {

	if returnedError := s.spotRepository.CreateOnline(spot); returnedError != nil {
		return returnedError
	}

	if returnedError := s.locationRepository.AppendSpot(spot); returnedError != nil {
		return returnedError
	}
	return nil

}

func (s *service) GetSpotByUserId(userId string) (domain.Spot, error) {
	spot, returnedError := s.spotRepository.GetSpotByUserId(userId)

	// TODO -> Improve this.
	if returnedError != nil {
		log.Println("Error on GetSpotByUserId: ", returnedError)
		return domain.Spot{}, apperrors.Internal
	}

	return spot, nil
}

// TODO -> Fix this!
func (s *service) EndSpot(spotId string) error {
	/*
		if returnedError := s.spotRepository.EndSpot(spotId); returnedError != nil {
			log.Println("Error on EndSpot: ", returnedError)
			return returnedError
		}

		if returnedError := s.locationRepository.RemoveSpot(spotId); returnedError != nil {
			log.Println("Error on EndSpot: ", returnedError)
			return returnedError
		}

		return nil
	*/
	return nil
}
