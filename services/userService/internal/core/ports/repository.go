package ports

import (
	"errors"
	"time"

	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/domain"
)

var (
	ErrUserDoesNotExist    = errors.New("err the user does not exist in the registry")
	ErrRegistryUnavailable = errors.New("err registry is temporaly unavailable")
	ErrInvalidOTP          = errors.New("err Invalid OTP")
	ErrThereIsNotOTP       = errors.New("err there is not OTP alive for the accout")
	ErrMaxRetriesOnTOP     = errors.New("err max retries in OTP reached")
	ErrOTPTTLStillLive     = errors.New("err the OTP is still alive")
	ErrOTPNotAlive         = errors.New("err the OTP is not alive")
	ErrOnDynamoDB          = errors.New("err while fetching/marshaling the response fom dynamodb")
)

// phoneNumber string Key | name string | validated bool | UserId tring | phoneSiganuture string | otp string/[]int | otp_ttl string

type UserRepository interface {
	// Fetch user data from repository
	GetUser(phoneNumber string) (*domain.User, error)
	UpdateProfile(*domain.UserProfile) error

	// Fetch User profile from repository
	// Returns:
	//	- ErrUserDoesNotExist
	// 	- UserProfile
	GetUserProfile(userId string) (*domain.UserProfile, error)

	// Returns latest OTP generation timestap
	GetLastOTPGenerationTimestap(phoneNumber string) (*time.Time, error)

	// Create a user in the repository
	CreateUser(phoneNumber, userName string) (*domain.User, error)

	// Save OTP
	AddOTP(phoneNumber string, otp []int, ttl time.Duration) error

	// Validate OTP, this should punish on wrong attemp
	ValidateOTP(user *domain.User, otp []int) (*domain.User, error)
}

// tokenId string | userId  string | longLiveRefreshToken string | shortLiveToken string | shortLiveTokenTTL number
type TokensRepository interface {
	GenerateJWTToken(user *domain.User) (string, error)
}
