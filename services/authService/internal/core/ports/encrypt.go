package ports

import "github.com/JuanGQCadavid/now-project/services/authService/internal/core/domain"

type Encrypt interface {
	DecodeToken(token string) (domain.Token, error)
}
