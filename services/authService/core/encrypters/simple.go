package encrypters

import (
	b64 "encoding/base64"
	"strings"

	"github.com/JuanGQCadavid/now-project/services/authService/core/core/domain"
	"github.com/JuanGQCadavid/now-project/services/authService/core/core/ports"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/golang-jwt/jwt/v5"
)

type SimpleEncrypt struct {
	jwtKey []byte
}

func NewSimpleEncrypt(jwtKey []byte) *SimpleEncrypt {
	return &SimpleEncrypt{
		jwtKey: jwtKey,
	}

}

func (srv *SimpleEncrypt) DecodeJWTToken(token string) (*domain.UserDetails, error) {

	claims := jwt.MapClaims{}
	tokenParsed, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return srv.jwtKey, nil
	})

	if err != nil {
		logs.Warning.Println("We fail to parse the token ", err.Error())
		return nil, ports.ErrBadFormatToken
	}

	if !tokenParsed.Valid {
		return nil, ports.ErrTokenNotValid
	}

	var userDetail *domain.UserDetails = &domain.UserDetails{}

	for key, val := range claims {
		switch key {
		case "userId":
			userDetail.UserID = val.(string)
		case "userPhone":
			userDetail.PhoneNumber = val.(string)
		case "userName":
			userDetail.Name = val.(string)
		case "session":
			userDetail.SessionHash = val.(string)
		}
	}

	return userDetail, nil
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
