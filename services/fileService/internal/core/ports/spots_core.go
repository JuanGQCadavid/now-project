package ports

import (
	"context"
	"errors"

	"github.com/JuanGQCadavid/now-project/services/fileService/internal/core/domain"
)

var (
	ErrCallingBackend error = errors.New("err calling backend service")
)

type SpotsCoreRepository interface {
	GetUserEventAccess(ctx context.Context, userID string, eventId string) (*domain.UserEventAccess, error)
	GetUserDateAccess(ctx context.Context, eventId, userID, dateId string) (*domain.UserEventAccess, error)
}
