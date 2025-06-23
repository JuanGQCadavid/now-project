package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/authService/core/core/domain"
)

var (
	ErrBadFormatToken = errors.New("err the user token is bad formatted")
	ErrTokenNotValid  = errors.New("err the user token is not valid")
)

type Encrypt interface {
	DecodeToken(token string) (domain.Token, error)
	DecodeJWTToken(token string) (*domain.UserDetails, error)
}
