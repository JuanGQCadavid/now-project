package tokens

import "github.com/JuanGQCadavid/now-project/services/userService/internal/core/domain"

// tokenId string | userId  string | longLiveRefreshToken string | shortLiveToken string | shortLiveTokenTTL number
type DynamoDBTokensRepository struct {
}

func (repo *DynamoDBTokensRepository) GeneratePairOfTokens(userID string) *domain.Tokens {
	return nil
}
