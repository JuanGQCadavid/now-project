package spotsrv

import (
	"github.com/JuanGQCadavid/now-project/services/spots/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spots/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/spots/pkg/uuidgen"
)

/*
	When implementin a interface is just creating a struct with all
	the methods that the interface defined.
*/
type service struct {
	spotRepository ports.SpotRepository
	uuidGen        uuidgen.UUIDGen
}

func New(spotRepository ports.SpotRepository, uuidGen uuidgen.UUIDGen) *service {
	return &service{
		spotRepository: spotRepository,
		uuidGen:        uuidGen,
	}
}

func (s service) Get(spotId string) (domain.Spot, error) {
	return nil, nil
}

func (s service) GoOnline(spot domain.Spot) (domain.Spot, error) {

}
