package ports

import "github.com/JuanGQCadavid/now-project/services/fileService/internal/core/domain"

type SpotsCoreRepository interface {
	GetUserEventAccess(userID string, eventId string) (*domain.UserEventAccess, error)
	GetUserDateAccess(userID string, dateId string) (*domain.UserEventAccess, error)
}
