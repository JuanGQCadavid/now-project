package ports

import "github.com/JuanGQCadavid/now-project/services/fileService/internal/core/domain"

type SpotsCoreRepository interface {
	GetUserEventAccess(userID string, eventOrDateId string) (*domain.UserEventAccess, error)
}
