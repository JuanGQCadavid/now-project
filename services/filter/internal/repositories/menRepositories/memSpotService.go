package menrepositories

import (
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"
)

type MemSpotService struct {
	data map[string]domain.SimpleSpot
}

func NewMenSpotService(data []domain.SimpleSpot) *MemSpotService {
	dataInMap := map[string]domain.SimpleSpot{}

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
