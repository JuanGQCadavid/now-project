package core

import (
	"context"

	"github.com/JuanGQCadavid/now-project/services/fileService/internal/core/domain"
	"github.com/rs/zerolog/log"
)

type FileService struct {
}

func NewFileService() *FileService {
	return &FileService{}
}
func (fs *FileService) GeneratePresignedURL(ctx context.Context, metadata *domain.FileMetadata) (presigned *domain.PresignedURL, err error) {
	var (
		logger = log.Ctx(ctx)
	)
	logger.Info().Any("Payload", metadata).Send()
	return nil, nil
}
