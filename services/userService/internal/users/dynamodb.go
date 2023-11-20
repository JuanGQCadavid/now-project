package users

import (
	"time"

	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/domain"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
)

// phoneNumber string Key | name string | validated bool | UserId tring | phoneSiganuture string | otp string/[]int | otp_ttl string

type DynamoDBUserRepository struct {
	svc       *dynamodb.DynamoDB
	tableName string
}

func NewDynamoDBUserRepository(tableName string, session *session.Session) *DynamoDBUserRepository {
	svc := dynamodb.New(session)
	return &DynamoDBUserRepository{
		tableName: tableName,
		svc:       svc,
	}
}

// Fetch user data from repository

func (repo *DynamoDBUserRepository) GetUser(phoneNumber string) (*domain.User, error) {
	return nil, nil
}

// Returns latest OTP generation timestap

func (repo *DynamoDBUserRepository) GetLastOTPGenerationTimestap(phoneNumber string) *time.Time {
	return nil
}

// Create a user in the repository

func (repo *DynamoDBUserRepository) CreateUser(phoneNumber, userName string) (*domain.User, error) {
	return nil, nil
}

// Save OTP

func (repo *DynamoDBUserRepository) AddOTP(phoneNumber string, otp []int, ttl time.Duration) error {
	return nil
}

// Validate OTP, this should punish on wrong attemp

func (repo *DynamoDBUserRepository) ValidateOTP(phoneNumber string, otp []int) error {
	return nil
}

func (repo *DynamoDBUserRepository) generateId() string {
	return uuid.NewString()
}
