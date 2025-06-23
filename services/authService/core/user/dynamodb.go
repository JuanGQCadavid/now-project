package user

import (
	"github.com/JuanGQCadavid/now-project/services/authService/core/core/domain"
	"github.com/JuanGQCadavid/now-project/services/authService/core/core/ports"
	"github.com/JuanGQCadavid/now-project/services/authService/core/utils"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBUserRepository struct {
	tableName string
	keyName   string
	indexName string
	svc       *dynamodb.DynamoDB
}

const (
	defaultTableKey string = "UserId"
)

func NewDynamoDBUserRepository(tableName, indexName string, session *session.Session) *DynamoDBUserRepository {
	svc := dynamodb.New(session)
	return &DynamoDBUserRepository{
		tableName: tableName,
		keyName:   defaultTableKey,
		indexName: indexName,
		svc:       svc,
	}
}

// Fetch user data from repository

func (repo *DynamoDBUserRepository) GetUserData(useId string) (*domain.User, error) {
	user := &domain.User{}

	if err := utils.DynamoQueryOneAndMapTo(repo.keyName, useId, repo.tableName, repo.indexName, user, repo.svc); err != nil {
		return nil, err
	}

	if user == nil || len(user.PhoneNumber) == 0 {
		return nil, ports.ErrUserDoesNotExist
	}

	return user, nil
}
