package ports

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/authService/internal/core/domain"
)

var (
	ErrBadFormatToken = errors.New("err the user token is bad formatted")
)

type Encrypt interface {
	DecodeToken(token string) (domain.Token, error)
}
