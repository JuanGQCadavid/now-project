package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/JuanGQCadavid/now-project/services/fileService/internal/adapters/objects3"
// 	"github.com/JuanGQCadavid/now-project/services/fileService/internal/core/domain"
// 	"github.com/aws/aws-sdk-go-v2/aws"
// 	"github.com/aws/aws-sdk-go-v2/config"
// 	"github.com/aws/aws-sdk-go-v2/service/s3"
// 	"github.com/aws/aws-sdk-go-v2/service/s3/types"
// )

// func main() {
// 	cfg, err := config.LoadDefaultConfig(context.TODO())
// 	if err != nil {
// 		log.Fatalf("unable to load SDK config, %v", err)
// 	}
// 	s3Service := objects3.NewS3ObjectStorage("pululaap-files-staging", cfg)

// 	pre, err := s3Service.GeneratePresignedURL(
// 		objects3.WithBucketAndKey("pululaap-files-staging", "test2.jpg"),
// 		objects3.WithPublicReadACL(),
// 		// WithExpireLimitInSeconds(100),
// 		objects3.WithControlAccess(domain.ControlAccess{
// 			EventId: "123",
// 			DateId:  "567",
// 		}),
// 	)

// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}

// 	log.Println(pre.ToString())
// }

// func main2() {

// 	cfg, err := config.LoadDefaultConfig(context.TODO())
// 	if err != nil {
// 		log.Fatalf("unable to load SDK config, %v", err)
// 	}

// 	s3Client := s3.NewFromConfig(cfg)

// 	out, err := s3Client.ListBuckets(context.Background(), &s3.ListBucketsInput{})

// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}

// 	for _, b := range out.Buckets {
// 		log.Println(*b.Name)
// 	}

// 	// ----

// 	presignClient := s3.NewPresignClient(s3Client)

// 	// link, err := presignClient.PresignPutObject(context.Background(),
// 	// 	&s3.PutObjectInput{
// 	// 		ACL:    types.ObjectCannedACLPublicRead,
// 	// 		Bucket: aws.String("pululaap-files-staging"),
// 	// 		Key:    aws.String("test1.jpg"),
// 	// 		// Expires: aws.Time(time.Now().Add(time.Duration(10) * time.Second)),
// 	// 		Metadata: map[string]string{
// 	// 			"x-amz-meta-works":   "hi",
// 	// 			"x-amz-meta-works-2": "hi-2",
// 	// 		},
// 	// 	}, func(opts *s3.PresignOptions) {
// 	// 		opts.Expires = time.Duration(60 * int64(time.Second))
// 	// 	})

// 	// if err != nil {
// 	// 	log.Printf("Couldn't get a presigned request to put. Here's why: %v\n",
// 	// 		err)
// 	// }

// 	// log.Println(link.Method)
// 	// log.Println(link.SignedHeader)
// 	// log.Println("------")
// 	// log.Println(link.URL)

// 	pre, err := GeneratePresignedURL(
// 		presignClient,
// 		WithBucketAndKey("pululaap-files-staging", "test2.jpg"),
// 		// WithPublicReadACL(),
// 		// WithExpireLimitInSeconds(100),
// 		WithControlAccess(domain.ControlAccess{
// 			EventId: "123",
// 			DateId:  "567",
// 		}),
// 	)

// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}

// 	log.Println(pre.ToString())

// }

// func GeneratePresignedURL(presignClient *s3.PresignClient, opts ...PresignedURLConfig) (*domain.PresignedURL, error) {
// 	var (
// 		theLink *generatedLink = NewGeneratedLink()
// 		header                 = make(map[string]string)
// 	)

// 	for _, opt := range opts {
// 		opt(theLink)
// 	}

// 	link, err := presignClient.PresignPutObject(context.Background(),
// 		theLink.params,
// 		func(opts *s3.PresignOptions) {
// 			opts.Expires = time.Duration(theLink.ttl * int64(time.Second))
// 		},
// 	)

// 	if err != nil {
// 		log.Printf("Couldn't get a presigned request to put. Here's why: %v\n",
// 			err)

// 		return nil, err
// 	}

// 	for key, val := range link.SignedHeader {
// 		if len(val) == 0 {
// 			continue
// 		}

// 		vals := val[0]
// 		for _, val_i := range val[1:] {
// 			vals = fmt.Sprintf("%s,%s", vals, val_i)
// 		}

// 		header[key] = vals

// 	}

// 	return &domain.PresignedURL{
// 		URL:     link.URL,
// 		Method:  link.Method,
// 		Headers: header,
// 	}, nil

// }

// type generatedLink struct {
// 	params *s3.PutObjectInput
// 	ttl    int64
// }

// func NewGeneratedLink() *generatedLink {
// 	return &generatedLink{
// 		ttl:    60, // 60 seconds to upload a picture should be enough
// 		params: &s3.PutObjectInput{},
// 	}
// }

// type PresignedURLConfig func(*generatedLink)

// func WithControlAccess(controlAccess domain.ControlAccess) PresignedURLConfig {
// 	return func(gl *generatedLink) {
// 		var (
// 			metadata = make(map[string]string)
// 		)

// 		if len(controlAccess.EventId) > 0 {
// 			metadata["x-amz-meta-event-id"] = controlAccess.EventId
// 		}
// 		if len(controlAccess.DateId) > 0 {
// 			metadata["x-amz-meta-date-id"] = controlAccess.DateId
// 		}

// 		gl.params.Metadata = metadata
// 	}
// }

// func WithExpireLimitInSeconds(seconds int) PresignedURLConfig {
// 	return func(gl *generatedLink) {
// 		gl.ttl = int64(seconds)
// 	}
// }

// func WithBucketAndKey(bucket, key string) PresignedURLConfig {
// 	return func(gl *generatedLink) {
// 		gl.params.Bucket = aws.String(bucket)
// 		gl.params.Key = aws.String(key)
// 	}
// }

// func WithPublicReadACL() PresignedURLConfig {
// 	return func(gl *generatedLink) {
// 		gl.params.ACL = types.ObjectCannedACLPublicRead
// 	}
// }
