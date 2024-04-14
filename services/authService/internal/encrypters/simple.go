package encrypters

import "github.com/JuanGQCadavid/now-project/services/authService/internal/core/domain"

type SimpleEncrypt struct {
}

func NewSimpleEncrypt() *SimpleEncrypt {
	return &SimpleEncrypt{}

}
func (srv *SimpleEncrypt) DecodeToken(token string) (domain.Token, error) {
	return domain.Token{}, nil
}
