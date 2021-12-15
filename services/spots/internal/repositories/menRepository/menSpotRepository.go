package menRepository

import (
	"encoding/json"
	"fmt"

	"github.com/JuanGQCadavid/now-project/services/spots/internal/core/domain"
	"github.com/matiasvarela/errors"
)

/*

	This will save all Objects in a men array
	just for testing / working propuse.

*/
type MenSpotRepository struct {
	repo map[string][]byte
}

func New() *MenSpotRepository {
	return &MenSpotRepository{
		repo: map[string][]byte{},
	}
}

func (r *MenSpotRepository) Get(id string) (domain.Spot, error) {

	if element, ok := r.repo[id]; ok {
		fmt.Println(element)
		spot := domain.Spot{}

		if err := json.Unmarshal(element, &spot); err != nil {
			return domain.Spot{}, errors.New(errors.DefaultError, err, "Error while doing unmarshal", "Dude")
		}

		return spot, nil
	}
	return domain.Spot{}, errors.Define("NOT_FOUND")
}

func (r *MenSpotRepository) CreateOnline(spot domain.Spot) error {

	spotToSave, err := json.Marshal(&spot)

	if err != nil {
		return errors.Define("pailas bebe, no se pudo guardar cuando lo estab marsheliando")
	}

	r.repo[spot.EventInfo.UUID] = spotToSave

	return nil
}

func (r *MenSpotRepository) GetSpotByUserId(personId string) (domain.Spot, error) {
	return domain.Spot{}, nil
}

func (r *MenSpotRepository) EndSpot(spotId string) error {
	return nil
}
