package objects3

import (
	"github.com/JuanGQCadavid/now-project/services/fileService/internal/core/domain"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3ObjectStorage struct {
	presignClient *s3.PresignClient
	s3Client      *s3.Client
	bucketName    string
}

func NewS3ObjectStorage(bucketName string, cfg aws.Config) *S3ObjectStorage {
	s3Client := s3.NewFromConfig(cfg)

	return &S3ObjectStorage{
		s3Client:      s3Client,
		presignClient: s3.NewPresignClient(s3Client),
		bucketName:    bucketName,
	}
}

func (repo *S3ObjectStorage) GeneratePresignedURL(path string, controlAccess *domain.ControlAccess) (*domain.PresignedURL, error) {
	args := []PresignedURLConfig{
		WithBucketAndKey(repo.bucketName, path),
		WithControlAccess(controlAccess),
		WithExpireLimitInSeconds(100),
	}

	if controlAccess == nil {
		args = append(args, WithPublicReadACL())
	}

	return generatePresignedURL(repo.presignClient, args...)
}
