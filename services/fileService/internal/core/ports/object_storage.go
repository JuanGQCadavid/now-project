package ports

import "github.com/JuanGQCadavid/now-project/services/fileService/internal/core/domain"

type ObjectRepository interface {
	GeneratePresignedURL(path string, controlAccess *domain.ControlAccess) (*domain.PresignedURL, error)
}
