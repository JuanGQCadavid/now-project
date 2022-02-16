package service

import (
	"log"

	"github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/core/ports"
	locationrepositories "github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/repositories/locationRepositories"
)

type LocationService struct {
	LocationRepo ports.LocationRepository
}

func (srv *LocationService) OnSpotCreation(spot domain.Spot) error {
	log.Printf("OnSpotCreation ->  %+v", spot)

	return srv.LocationRepo.CrateLocation(spot)
}

func (srv *LocationService) OnSpotDeletion(spotId string) error {
	log.Println("OnSpotDeletion -> ", spotId)
	return nil
}

func NewLocationService() *LocationService {
	return &LocationService{
		LocationRepo: locationrepositories.NewLocationRepo(),
	}
}
