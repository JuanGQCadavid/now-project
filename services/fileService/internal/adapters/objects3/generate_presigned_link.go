package objects3

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/JuanGQCadavid/now-project/services/fileService/internal/core/domain"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type generatedLink struct {
	params *s3.PutObjectInput
	ttl    int64
}

type PresignedURLConfig func(*generatedLink)

func newGeneratedLink() *generatedLink {
	return &generatedLink{
		ttl:    60, // 60 seconds to upload a picture should be enough
		params: &s3.PutObjectInput{},
	}
}

func generatePresignedURL(presignClient *s3.PresignClient, opts ...PresignedURLConfig) (*domain.PresignedURL, error) {
	var (
		theLink *generatedLink = newGeneratedLink()
		header                 = make(map[string]string)
	)

	for _, opt := range opts {
		opt(theLink)
	}

	link, err := presignClient.PresignPutObject(context.Background(),
		theLink.params,
		func(opts *s3.PresignOptions) {
			opts.Expires = time.Duration(theLink.ttl * int64(time.Second))
		},
	)

	if err != nil {
		log.Printf("Couldn't get a presigned request to put. Here's why: %v\n",
			err)

		return nil, err
	}

	for key, val := range link.SignedHeader {
		if len(val) == 0 {
			continue
		}

		vals := val[0]
		for _, val_i := range val[1:] {
			vals = fmt.Sprintf("%s,%s", vals, val_i)
		}

		header[key] = vals

	}

	return &domain.PresignedURL{
		URL:     link.URL,
		Method:  link.Method,
		Headers: header,
	}, nil

}

func WithControlAccess(controlAccess *domain.ControlAccess) PresignedURLConfig {
	return func(gl *generatedLink) {
		if controlAccess == nil {
			return
		}

		var (
			metadata = make(map[string]string)
		)

		if len(controlAccess.EventId) > 0 {
			metadata["x-amz-meta-event-id"] = controlAccess.EventId
		}
		if len(controlAccess.DateId) > 0 {
			metadata["x-amz-meta-date-id"] = controlAccess.DateId
		}

		gl.params.Metadata = metadata
	}
}

func WithExpireLimitInSeconds(seconds int) PresignedURLConfig {
	return func(gl *generatedLink) {
		gl.ttl = int64(seconds)
	}
}

func WithBucketAndKey(bucket, key string) PresignedURLConfig {
	return func(gl *generatedLink) {
		gl.params.Bucket = aws.String(bucket)
		gl.params.Key = aws.String(key)
	}
}

func WithPublicReadACL() PresignedURLConfig {
	return func(gl *generatedLink) {
		gl.params.ACL = types.ObjectCannedACLPublicRead
	}
}
