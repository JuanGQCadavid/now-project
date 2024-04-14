package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/authService/internal/core/domain"
)

var (
	ErrUserDoesNotExist = errors.New("err the user does not exist in the registry")
	ErrTokenNotFound    = errors.New("err invalidToken")
)

type TokensRepository interface {
	IsTokenValid(domain.Token) error
	GetTokenData(token domain.Token) (*domain.Tokens, error)
}

type UserRepository interface {
	GetUserData(token *domain.Tokens) (*domain.User, error)
}
