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
	ErrTokenHashNotValid    = errors.New("err token hash is not valid")
)

type AuthService struct {
	userRepo    ports.UserRepository
	encryptRepo ports.Encrypt
}

func NewAuthService(encryptRepo ports.Encrypt, userRepo ports.UserRepository) *AuthService {
	return &AuthService{
		encryptRepo: encryptRepo,
		userRepo:    userRepo,
	}
}

func (svc *AuthService) GetUserDetailsFromToken(tokenEncoded string) (*domain.UserDetails, error) {
	var (
		useDetail *domain.UserDetails
		user      *domain.User
	)

	if len(tokenEncoded) == 0 {
		logs.Info.Println("Anonymous user")
		return nil, nil
	}

	useDetail, err := svc.encryptRepo.DecodeJWTToken(tokenEncoded)

	if err != nil {
		return nil, ErrWhileDecryptingToken
	}

	user, err = svc.userRepo.GetUserData(useDetail.UserID)

	if err != nil {
		return nil, err
	}

	if useDetail.SessionHash != user.ValidatedHash {
		return nil, ErrTokenHashNotValid
	}

	return &domain.UserDetails{
		UserID:      user.UserId,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
	}, nil
}
