package menrepositories

import (
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/models"
)

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

func (sr *MemSpotService) GetSpotsCardsInfo(spots []string) ([]domain.Spot, error) {
	var response = make([]domain.Spot, len(spots))

	arrayCounter := 0

	for _, spotId := range spots {

		data := sr.data[spotId]

		if len(data.Id) != 0 {
			response[arrayCounter] = domain.Spot{
				EventInfo: domain.Event{
					UUID: data.Id,
				},
				PlaceInfo: domain.Place{
					Lat: float64(data.LatLng.Lat),
					Lon: float64(data.LatLng.Lng),
				},
			}
			arrayCounter++
		}

	}

	return response, nil
}
