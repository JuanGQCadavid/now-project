package service

import (
	"github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
)

type LocationService struct {
	locationRepo ports.LocationRepository
}

func NewLocationService(locationRepo ports.LocationRepository) *LocationService {
	return &LocationService{
		locationRepo: locationRepo,
	}
}

func (srv *LocationService) OnDateCreation(date domain.DatesLocation) error {
	logs.Info.Printf("OnDateCreation: date: %v\n", date)
	return srv.locationRepo.CrateLocation(date)
}

func (srv *LocationService) OnDateRemoved(dateId string) error {
	logs.Info.Printf("OnDateRemoved: date: %v\n", dateId)
	return srv.locationRepo.RemoveLocation(dateId)
}

func (srv *LocationService) OnDateStatusChanged(dateId string, newStatus domain.DateStatus) error {
	logs.Info.Printf("OnDateStatusChanged: date: %v, new status: %v\n", dateId, string(newStatus))
	return srv.locationRepo.UpdateLocationStatus(dateId, newStatus)
}
