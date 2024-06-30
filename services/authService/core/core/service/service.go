package service

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/authService/core/core/domain"
	"github.com/JuanGQCadavid/now-project/services/authService/core/core/ports"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
)

var (
	ErrWhileDecryptingToken = errors.New("err while decoding the token, maybe bad format")
	ErrWhileFetchingUser    = errors.New("err unknown user")
)

type AuthService struct {
	tokensRepo  ports.TokensRepository
	userRepo    ports.UserRepository
	encryptRepo ports.Encrypt
}

func NewAuthService(tokensRepo ports.TokensRepository, encryptRepo ports.Encrypt, userRepo ports.UserRepository) *AuthService {
	return &AuthService{
		tokensRepo:  tokensRepo,
		encryptRepo: encryptRepo,
		userRepo:    userRepo,
	}
}

func (svc *AuthService) GetUserDetailsFromToken(tokenEncoded string) (*domain.UserDetails, error) {
	var (
		token domain.Token
		user  *domain.User
	)

	if len(tokenEncoded) == 0 {
		logs.Info.Println("Anonymous user")
		return &domain.UserDetails{}, nil
	}

	token, err := svc.encryptRepo.DecodeToken(tokenEncoded)

	if err != nil {
		return nil, ErrWhileDecryptingToken
	}

	tokenData, err := svc.tokensRepo.GetTokenData(token)

	if err != nil {
		return nil, err
	}

	user, err = svc.userRepo.GetUserData(tokenData)

	if err != nil {
		return nil, err
	}

	return &domain.UserDetails{
		UserID:      user.UserId,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
	}, nil
}
