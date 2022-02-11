package menrepositories

import "github.com/JuanGQCadavid/now-project/services/filter/internal/core/models"

type MemSpotService struct {
	data map[string]models.Spot
}

func NewMenSpotService(data []models.Spot) *MemSpotService {
	dataInMap := map[string]models.Spot{}

	for _, spotToAdd := range data {
		dataInMap[spotToAdd.Id] = spotToAdd
	}

	return &MemSpotService{
		data: dataInMap,
	}
}

func (sr *MemSpotService) GetSpotsCardsInfo(spots []string) ([]models.Spot, error) {
	var response = make([]models.Spot, len(spots))

	arrayCounter := 0

	for _, spotId := range spots {

		data := sr.data[spotId]

		if len(data.Id) != 0 {
			response[arrayCounter] = data
			arrayCounter++
		}

	}

	return response, nil
}
