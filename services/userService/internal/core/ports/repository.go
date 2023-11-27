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
	ErrMaxRetriesOnTOP     = errors.New("err max retries in OTP reached")
	ErrOTPTTLStillLive     = errors.New("err the OTP is still alive")
)

// phoneNumber string Key | name string | validated bool | UserId tring | phoneSiganuture string | otp string/[]int | otp_ttl string

type UserRepository interface {
	// Fetch user data from repository
	GetUser(phoneNumber string) (*domain.User, error)

	// Returns latest OTP generation timestap
	GetLastOTPGenerationTimestap(phoneNumber string) *time.Time

	// Create a user in the repository
	CreateUser(phoneNumber, userName string) (*domain.User, error)

	// Save OTP
	AddOTP(phoneNumber string, otp []int, ttl time.Duration) error

	// Validate OTP, this should punish on wrong attemp
	ValidateOTP(phoneNumber string, otp []int) error
}

// tokenId string | userId  string | longLiveRefreshToken string | shortLiveToken string | shortLiveTokenTTL number
type TokensRepository interface {
	GeneratePairOfTokens(userID string) *domain.Tokens
}
