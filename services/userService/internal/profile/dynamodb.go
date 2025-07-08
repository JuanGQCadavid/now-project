package profile

import (
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/utils"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type ProfileRepositoryDynamoDB struct {
	svc       *dynamodb.DynamoDB
	tableName string
}

const (
	DEFAULT_USER_ID_KEY_NAME string = "UserId"
)

func NewProfileRepositoryDynamoDB(tableName string, session *session.Session) *ProfileRepositoryDynamoDB {
	svc := dynamodb.New(session)
	return &ProfileRepositoryDynamoDB{
		tableName: tableName,
		svc:       svc,
	}
}

func (repo *ProfileRepositoryDynamoDB) CreateProfile(profile *domain.UserProfile) error {
	return repo.UpdateProfile(profile)
}

func (repo *ProfileRepositoryDynamoDB) UpdateProfile(profile *domain.UserProfile) error {
	return utils.UpdateItem(profile, repo.tableName, repo.svc)
}

// Fetch User profile from repository
// Returns:
//   - ErrUserDoesNotExist
//   - UserProfile
func (repo *ProfileRepositoryDynamoDB) GetUserProfile(userId string) (*domain.UserProfile, error) {
	var (
		userProfile *domain.UserProfile = &domain.UserProfile{}
	)
	if err := utils.DynamoGetAndMapTo(DEFAULT_USER_ID_KEY_NAME, userId, repo.tableName, userProfile, repo.svc); err != nil {
		logs.Error.Println("We fail to fetch user profile, dynamodb: ", err.Error())
		return nil, ports.ErrOnDynamoDB
	}

	if len(userProfile.UserName) == 0 {
		return nil, ports.ErrUserNotFound
	}

	return userProfile, nil
}
