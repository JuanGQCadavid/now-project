package ports

import "github.com/JuanGQCadavid/now-project/services/confirmationService/internal/core/domain"

type Repository interface {
	FetchDate(dateId string, hostId string) (*domain.Date, error)
	UpdateDateOnConfirmed(dateId string, confirmed bool) error
}
