package dbspotservicelambda

import "github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"

type SpotResponse struct {
	Spots []domain.Spot `json:"Spots"`
}
