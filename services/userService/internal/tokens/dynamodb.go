package tokens

import (
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// tokenId string | userId  string | longLiveRefreshToken string | shortLiveToken string | shortLiveTokenTTL number
type DynamoDBTokensRepository struct {
	svc       *dynamodb.DynamoDB
	tableName string
	separator string
	shortTTL  time.Duration
	key       []byte
}

func NewDynamoDBTokensRepository(key []byte, tableName string, session *session.Session) *DynamoDBTokensRepository {
	svc := dynamodb.New(session)
	return &DynamoDBTokensRepository{
		tableName: tableName,
		svc:       svc,
		separator: "+",
		shortTTL:  time.Duration(3 * time.Hour),
		key:       key,
	}
}

func (repo *DynamoDBTokensRepository) GenerateJWTToken(user domain.User) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId":    user.UserId,
			"userPhone": user.PhoneNumber,
			"userName":  user.Name,
		},
	)
	return token.SignedString(repo.key)
}

// tokenId string | userId  string | longLiveRefreshToken string | shortLiveToken string | shortLiveTokenTTL number

func (repo *DynamoDBTokensRepository) GeneratePairOfTokens(userID string) (*domain.Tokens, error) {
	tokenId := repo.generateId()

	token := &domain.Tokens{
		UserID:            userID,
		TokenId:           tokenId,
		RefreshToken:      repo.maskBase64(repo.genKey(tokenId)),
		ShortLiveToken:    repo.maskBase64(repo.genKey(tokenId)),
		ShortLiveTokenTTL: time.Now().Add(repo.shortTTL),
	}

	marshaled, err := dynamodbattribute.MarshalMap(token)

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

	return token, nil
}

func (repo *DynamoDBTokensRepository) maskBase64(token string) string {
	return base64.StdEncoding.EncodeToString([]byte(token))
}

func (repo *DynamoDBTokensRepository) genKey(tokenId string) string {
	return fmt.Sprintf("%s%s%s", tokenId, repo.separator, repo.generateId())
}

func (repo *DynamoDBTokensRepository) generateId() string {
	return uuid.NewString()
}
