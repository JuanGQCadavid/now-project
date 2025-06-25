package ports

import (
	"errors"

	authDomain "github.com/JuanGQCadavid/now-project/services/authService/core/core/domain"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/domain"
)

var (
	ErrUserNotFound                     = errors.New("err user not found in registry")
	ErrInternalError                    = errors.New("err internal")
	ErrOnSavingOTP                      = errors.New("err on saving OTP")
	ErrOnSendingOTP                     = errors.New("err on sending OTP")
	ErrUserIsAlreadyCreated             = errors.New("err the user already exist")
	ErrUserNameShouldContainOnlyLetters = errors.New("err user name should only contain letters")
	ErrUserNotLogged                    = errors.New("err user needs to be logged in to perform the opperation")
	ErrSameProfile                      = errors.New("err user is the same")
)

type UserService interface {
	// Fetch user info
	// If the user token is the same user then
	// Return all data, if not only public data
	GetUserInfo(user *authDomain.UserDetails, userId string) (*domain.UserProfile, error)
	UpdateProfile(user *authDomain.UserDetails, profile *domain.UserProfile) error

	// user/init/login

	// Look up in user registry, if user phone nomber is already stored then it a OTP
	// will be sent to the user with their predererred method, if user is not in registry
	// then a error ErrUserNotFound will be returned
	InitLogin(userInfo domain.Login) error

	// user/init/singup
	InitSingUp(userInfo domain.SingUp) error

	// user/validate/<userID>/otp
	// user/otp/<userID>/validate
	// This is crucial as it should validate the OTP
	// If it is correct then the user process (login or sing up) is completed
	// moreover, as the user is already in, a set of  refresh and short-term key token
	// in order to authenticate the user to make request
	ValidateProcess(validateProcess domain.ValidateProcess) (*domain.Tokens, error)

	// user/otp/<userId>/resent
	// user/validare/<userId>/otp/resent
	GenerateNewOTP(userInfo domain.Login) error
}
