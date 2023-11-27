package users

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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
	user := &domain.User{
		Name:        userName,
		PhoneNumber: phoneNumber,
		Validated:   false,
		UserId:      repo.generateId(),
	}

	marshaled, err := dynamodbattribute.MarshalMap(user)

	if err != nil {
		logs.Error.Println("err we fail marshaling the data, err: ", err.Error())
		return nil, errors.New("err we fail marshaling the data")
	}

	userInput := &dynamodb.PutItemInput{
		Item:                   marshaled,
		TableName:              &repo.tableName,
		ReturnConsumedCapacity: aws.String(dynamodb.ReturnConsumedCapacityTotal),
	}

	_, err = repo.svc.PutItem(userInput)

	if err != nil {
		logs.Error.Println("err we fail putting the data, err", err.Error())
		return nil, errors.New("err we fail putting the data")
	}

	return user, nil
}

type UserOTP struct {
	OTP      []int         `json:"otp"`
	TTL      time.Duration `json:"ttl"`
	Attempts int           `json:"attempts"`
}

// Save OTP
func (repo *DynamoDBUserRepository) AddOTP(phoneNumber string, otp []int, ttl time.Duration) error {

	newOTP := &OTP{
		PhoneNumber: phoneNumber,
	}

	key, err := dynamodbattribute.MarshalMap(newOTP)

	if err != nil {
		log.Fatalf("Got error marshalling key item: %s", err)
	}

	expressionAttributeValues := map[string]*dynamodb.AttributeValue{}
	expressionAttributesNames := map[string]*string{}

	updateExpression := "set"
	updateExpression = fmt.Sprintf("%s #%s = :%s,", updateExpression, "OTP", "OTP")
	updateExpression = fmt.Sprintf("%s #%s = :%s,", updateExpression, "OTP_TTL", "OTP_TTL")
	// Remove last " , " in the updateExpression
	updateExpression = strings.TrimSuffix(updateExpression, ",")

	expressionAttributesNames[fmt.Sprintf("#%s", "OTP")] = aws.String("OTP")
	expressionAttributesNames[fmt.Sprintf("#%s", "OTP_TTL")] = aws.String("OTP_TTL")

	otpAttribute, err := dynamodbattribute.Marshal(otp)
	if err != nil {
		logs.Error.Println("We fail casting the OTP, err: ", err.Error())
		return err
	}
	expressionAttributeValues[fmt.Sprintf(":%s", "OTP")] = otpAttribute

	ttlAttribute, err := dynamodbattribute.Marshal(ttl)
	if err != nil {
		logs.Error.Println("We fail casting the TTL, err: ", err.Error())
		return err
	}
	expressionAttributeValues[fmt.Sprintf(":%s", "OTP_TTL")] = ttlAttribute

	input := &dynamodb.UpdateItemInput{
		Key:                       key,
		TableName:                 &repo.tableName,
		ExpressionAttributeValues: expressionAttributeValues,
		ExpressionAttributeNames:  expressionAttributesNames,
		ReturnConsumedCapacity:    aws.String(dynamodb.ReturnConsumedCapacityIndexes),
		UpdateExpression:          &updateExpression,
	}

	_, err = repo.svc.UpdateItem(input)
	if err != nil {
		logs.Error.Println("We fail updating the item: ", err.Error())
		return err
	}
	return err
}

// Validate OTP, this should punish on wrong attemp

func (repo *DynamoDBUserRepository) ValidateOTP(phoneNumber string, otp []int) error {
	return nil
}

func (repo *DynamoDBUserRepository) generateId() string {
	return uuid.NewString()
}
