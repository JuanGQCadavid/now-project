package tokens

import (
	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/domain"
	"github.com/golang-jwt/jwt/v5"
)

type JWTTokenGenerator struct {
	key []byte
}

func NewJWTTokenGenerator(key []byte) *JWTTokenGenerator {
	return &JWTTokenGenerator{
		key: key,
	}
}

func (repo *JWTTokenGenerator) GenerateJWTToken(user *domain.User) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId":    user.UserId,
			"userPhone": user.PhoneNumber,
			"userName":  user.Name,
			"session":   user.ValidatedHash,
		},
	)
	return token.SignedString(repo.key)
}
