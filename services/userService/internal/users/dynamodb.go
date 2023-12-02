package users

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/ports"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

// phoneNumber string Key | name string | validated bool | UserId tring | phoneSiganuture string | otp string/[]int | otp_ttl string

type DynamoDBUserRepository struct {
	svc        *dynamodb.DynamoDB
	tableName  string
	maxRetries int
}

func NewDynamoDBUserRepository(tableName string, session *session.Session) *DynamoDBUserRepository {
	svc := dynamodb.New(session)
	return &DynamoDBUserRepository{
		tableName:  tableName,
		svc:        svc,
		maxRetries: 3,
	}
}

// Fetch user data from repository

func (repo *DynamoDBUserRepository) GetUser(phoneNumber string) (*domain.User, error) {
	user := &domain.User{}

	if err := repo.getAndMapTo(phoneNumber, user); err != nil {
		return nil, err
	}

	if user == nil || len(user.PhoneNumber) == 0 {
		return nil, ports.ErrUserDoesNotExist
	}

	return user, nil
}

// Returns latest OTP generation timestap

func (repo *DynamoDBUserRepository) GetLastOTPGenerationTimestap(phoneNumber string) (*time.Time, error) {
	otpFromRepo, err := repo.getOTP(phoneNumber)
	logs.Info.Printf("OTP: %+v \n", otpFromRepo)

	if err != nil {
		logs.Error.Println("We fail to fetch the OTP")
		return nil, err
	}

	if otpFromRepo != nil {
		return &otpFromRepo.TTL, nil
	}
	return nil, nil
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

// Save OTP
func (repo *DynamoDBUserRepository) AddOTP(phoneNumber string, otp []int, ttl time.Duration) error {
	now := time.Now()

	return repo.updateSingleAttribute(phoneNumber, "OTP", &UserOTP{
		OTP:      otp,
		TTL:      now.Add(ttl),
		Attempts: 0,
	})
}

// Validate OTP, this should punish on wrong attemp

func (repo *DynamoDBUserRepository) ValidateOTP(phoneNumber string, otp []int) error {

	otpFromRepo, err := repo.getOTP(phoneNumber)
	logs.Info.Printf("OTP: %+v \n", otpFromRepo)

	if err != nil {
		return err
	}

	if otpFromRepo == nil {
		return ports.ErrThereIsNotOTP
	}

	if len(otp) != len(otpFromRepo.OTP) {
		logs.Error.Println("lengths are not the same for the OTPs")
		return repo.punishOTP(phoneNumber, otpFromRepo)
	}

	for i, code := range otp {
		if otpFromRepo.OTP[i] != code {
			logs.Error.Println("OTP does not match")
			return repo.punishOTP(phoneNumber, otpFromRepo)
		}
	}

	if err := repo.updateSingleAttribute(phoneNumber, "Validated", true); err != nil {
		logs.Error.Println("We fail to update the user, err", err.Error())
		return err
	}

	logs.Info.Println("Code validation goes well")
	return repo.cleanOTP(phoneNumber, otpFromRepo)
}

func (repo *DynamoDBUserRepository) cleanOTP(phoneNumber string, userOTP *UserOTP) error {
	return repo.updateSingleAttribute(phoneNumber, "OTP", nil)
}

func (repo *DynamoDBUserRepository) punishOTP(phoneNumber string, userOTP *UserOTP) error {
	userOTP.Attempts++

	if userOTP.Attempts >= repo.maxRetries {
		if err := repo.cleanOTP(phoneNumber, userOTP); err != nil {
			logs.Error.Println("We fail to clean the code")
			return err
		}
		logs.Warning.Println("OTP Clean")

		return ports.ErrMaxRetriesOnTOP
	}

	if err := repo.updateSingleAttribute(phoneNumber, "OTP", userOTP); err != nil {
		return err
	}

	return ports.ErrInvalidOTP
}

func (repo *DynamoDBUserRepository) updateSingleAttribute(phoneNumber string, attribute string, value any) error {
	logs.Info.Println("Updating OTP: phoneNumber: ", phoneNumber, " value: ", value)

	otpAttribute, err := dynamodbattribute.Marshal(value)
	if err != nil {
		return err
	}

	key, err := repo.genKey(phoneNumber)
	if err != nil {
		return err
	}

	if err != nil {
		logs.Error.Println("We fail casting the VALUE, err: ", err.Error())
		return err
	}

	expressionAttributeValues := map[string]*dynamodb.AttributeValue{}
	expressionAttributeValues[":VALUE"] = otpAttribute

	out, err := repo.svc.UpdateItem(&dynamodb.UpdateItemInput{
		Key:                       key,
		TableName:                 &repo.tableName,
		ExpressionAttributeValues: expressionAttributeValues,
		ReturnConsumedCapacity:    aws.String(dynamodb.ReturnConsumedCapacityIndexes),
		UpdateExpression:          aws.String(fmt.Sprintf("set %s = :VALUE", attribute)),
	})

	if err != nil {
		logs.Error.Println("We fail updating the item: ", err.Error())
		return err
	}

	logs.Info.Println(out)
	return err

}

func (repo *DynamoDBUserRepository) generateId() string {
	return uuid.NewString()
}

func (repo *DynamoDBUserRepository) getOTP(phoneNumber string) (*UserOTP, error) {

	resp := &UserOTPBody{}

	if err := repo.getAndMapTo(phoneNumber, resp); err != nil {
		return nil, err
	}

	return resp.OTP, nil
}

func (repo *DynamoDBUserRepository) getAndMapTo(phoneNumber string, mapTo any) error {
	key, err := repo.genKey(phoneNumber)
	if err != nil {
		return err
	}

	out, err := repo.svc.GetItem(&dynamodb.GetItemInput{
		// ProjectionExpression: aws.String("OTP"),
		// ConsistentRead:       aws.Bool(true),
		TableName: aws.String(repo.tableName),
		Key:       key,
	})

	if err != nil {
		logs.Error.Println("We fail to get the OTP, error: ", err.Error())
		return err
	}

	logs.Info.Printf("%+v\n", out)

	err = dynamodbattribute.UnmarshalMap(out.Item, mapTo)

	if err != nil {
		logs.Error.Println("We fail to unmarshal the OTP, error: ", err.Error())
		return err
	}

	return nil
}

func (repo *DynamoDBUserRepository) genKey(phoneNumber string) (map[string]*dynamodb.AttributeValue, error) {
	key, err := dynamodbattribute.MarshalMap(TableKey{
		PhoneNumber: phoneNumber,
	})

	if err != nil {
		log.Fatalf("Got error marshalling key item: %s", err)
		return nil, err
	}

	return key, nil
}
