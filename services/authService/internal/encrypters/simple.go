package encrypters

import (
	b64 "encoding/base64"
	"strings"

	"github.com/JuanGQCadavid/now-project/services/authService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/authService/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
)

type SimpleEncrypt struct {
}

func NewSimpleEncrypt() *SimpleEncrypt {
	return &SimpleEncrypt{}

}
func (srv *SimpleEncrypt) DecodeToken(token string) (domain.Token, error) {
	sDec, err := b64.StdEncoding.DecodeString(token)

	if err != nil {
		logs.Error.Println("We fail to decode the token, err", err.Error())
		return domain.Token{}, err
	}

	tokenDecoded := strings.Split(string(sDec), "+")

	if len(tokenDecoded) != 2 {
		logs.Error.Println("Bad format token: ", string(sDec))
		return domain.Token{}, ports.ErrBadFormatToken
	}

	return domain.Token{
		TokenID:    tokenDecoded[0],
		TokenValue: tokenDecoded[1],
	}, nil
}
