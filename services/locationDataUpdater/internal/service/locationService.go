package service

import (
	"log"

	"github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/core/domain"
)

type LocationService struct {
}

func (srv *LocationService) OnSpotCreation(spot domain.Spot) error {
	log.Printf("OnSpotCreation ->  %+v", spot)

	return nil
}

func (srv *LocationService) OnSpotDeletion(spotId string) error {
	log.Println("OnSpotDeletion -> ", spotId)
	return nil
}

func NewLocationService() *LocationService {
	return &LocationService{}
}
