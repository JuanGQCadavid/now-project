package tokens

import (
	"github.com/JuanGQCadavid/now-project/services/authService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/authService/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/authService/internal/encrypters"
	"github.com/JuanGQCadavid/now-project/services/authService/internal/utils"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBTokensRepository struct {
	tableName string
	keyName   string
	svc       *dynamodb.DynamoDB
	decrypter *encrypters.SimpleEncrypt
}

const (
	defaultKeyName string = "TokenId"
)

func NewDynamoDBTokensRepository(tableName string, session *session.Session, decrypter *encrypters.SimpleEncrypt) *DynamoDBTokensRepository {
	return &DynamoDBTokensRepository{
		tableName: tableName,
		keyName:   defaultKeyName,
		svc:       dynamodb.New(session),
		decrypter: decrypter,
	}
}

func (repo *DynamoDBTokensRepository) GetTokenData(token domain.Token) (*domain.Tokens, error) {
	var (
		tokens *domain.Tokens
	)

	tokens = &domain.Tokens{}
	if err := utils.DynamoGetAndMapTo(repo.keyName, token.TokenID, repo.tableName, tokens, repo.svc); err != nil {
		return nil, err
	}

	if len(tokens.UserID) == 0 {
		return nil, ports.ErrTokenNotFound
	}

	shortDecryted, _ := repo.decrypter.DecodeToken(tokens.ShortLiveToken)

	if shortDecryted.TokenValue != token.TokenValue {
		return nil, ports.ErrTokenNotFound
	}

	return tokens, nil
}

func (repo *DynamoDBTokensRepository) IsTokenValid(token domain.Token) error {
	var (
		tokens *domain.Tokens
	)

	tokens = &domain.Tokens{}
	if err := utils.DynamoGetAndMapTo(repo.keyName, token.TokenID, repo.tableName, tokens, repo.svc); err != nil {
		return err
	}

	if len(tokens.UserID) == 0 || tokens.ShortLiveToken != token.TokenValue {
		return ports.ErrTokenNotFound
	}

	return nil
}
