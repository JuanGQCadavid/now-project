package dbspotservicelambda

import (
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/ports"
	"gorm.io/gorm"
)

type DBSpotServiceLambda struct {
	db *gorm.DB
}

func NewDBSpotServiceLambdaWithDriver(db *gorm.DB) (*DBSpotServiceLambda, error) {
	return &DBSpotServiceLambda{
		db: db,
	}, nil
}

const (
	SPOT_URL = "spotServiceURL"
)

func (srv *DBSpotServiceLambda) GetSpotsCardsInfo(datesIds []string, format ports.OutputFormat) ([]domain.Spot, error) {
	return nil, nil
}
