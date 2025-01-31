package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/ports"

	authDomain "github.com/JuanGQCadavid/now-project/services/authService/core/core/domain"
)

const (
	OTPEnglish string = "Hi, your code is %s. Best regards, Pululapp."
	OTPSpanish string = "Hola, tu codigo es %s. Pululap."
)

type OTPConfig struct {
	DefaultLength int
	TTL           time.Duration
}

type Service struct {
	userRepository   ports.UserRepository
	tokensRepository ports.TokensRepository
	otpConifg        *OTPConfig
	notificators     map[domain.NotificatorType]ports.Notificator
}

func NewService(userRepository ports.UserRepository, notificators map[domain.NotificatorType]ports.Notificator, tokensRepository ports.TokensRepository) *Service {
	return &Service{
		userRepository: userRepository,
		otpConifg: &OTPConfig{
			DefaultLength: 4,
			TTL:           time.Minute * 5, // 5 minutes
		},
		notificators:     notificators,
		tokensRepository: tokensRepository,
	}
}

func (svc *Service) GetUserInfo(user *authDomain.UserDetails, userId string) (*domain.UserProfile, error) {
	userProfile, err := svc.userRepository.GetUserProfile(userId)

	if err != nil {
		if err != ports.ErrUserDoesNotExist {
			return nil, ports.ErrInternalError
		}
		logs.Error.Println("We found an error while fetching the user profile for id [", userId, "] err:", err.Error())
		return nil, err
	}

	logs.Info.Println(userProfile)

	if user.UserID == authDomain.AnonymousUser.UserID || userId != user.UserID {
		userProfile.CleanSensitiveData()
		return userProfile, nil
	}

	return userProfile, nil
}

/*
 1. Find user by phone
    If user exist
    send OTP

    If user does not exist
    Return error

2. Generate OTP

3. Save OTP

4. Send OTP
*/
func (svc *Service) InitLogin(userInfo domain.Login) error {

	user, err := svc.getUser(userInfo.PhoneNumber)

	if err != nil {
		return err
	}

	return svc.initProcessValidation(user, userInfo.MethodVerificator)
}

// 1. Validate if the user is already in the registry
// 	TRUE - abort with error

// 2. save the user data in the registry

// 3. Send OTP to validate user

func (svc *Service) InitSingUp(userInfo domain.SingUp) error {

	userFetched, err := svc.getUser(userInfo.PhoneNumber)

	// Does the user already exist and it is validated?
	if err == nil && userFetched.Validated {
		logs.Info.Println("User is already created and validated")
		return ports.ErrUserIsAlreadyCreated
	}

	if err == nil && !userFetched.Validated {
		logs.Info.Println("User was attempted to be created but it is still not validated")

		isOTPAlive, err := svc.isOTPStillAlive(userFetched)

		if err != nil {
			logs.Error.Println("a error occour, err: ", err.Error())
			return err
		}

		if isOTPAlive {
			logs.Error.Println("OTP is still alive")
			return ports.ErrOTPTTLStillLive
		}

	}

	if err != nil && err != ports.ErrUserNotFound {
		logs.Error.Println("The service fail and it is not due to user not found")
		return err
	}

	userCreated, err := svc.userRepository.CreateUser(userInfo.PhoneNumber, userInfo.UserName)

	if err != nil {
		return err
	}

	return svc.initProcessValidation(userCreated, userInfo.MethodVerificator)
}

// This is crucial as it should validate the OTP
// If it is correct then the user process (login or sing up) is completed
// moreover, as the user is already in, a set of  refresh and short-term key token
// in order to authenticate the user to make request
// 1. Get user
// 2. validate OTP
// 3. Send new tokens
func (svc *Service) ValidateProcess(validateProcess domain.ValidateProcess) (*domain.Tokens, error) {
	userFetched, err := svc.getUser(validateProcess.PhoneNumber)

	if err != nil {
		return nil, err
	}

	isOTPAlive, err := svc.isOTPStillAlive(userFetched)

	if err != nil {
		logs.Error.Println("An error occur while checking OTP ", err.Error())
	}

	if !isOTPAlive {
		logs.Error.Println("OTP is not alive")
		return nil, ports.ErrOTPNotAlive
	}

	err = svc.userRepository.ValidateOTP(validateProcess.PhoneNumber, validateProcess.Code)

	if err != nil {
		return nil, err
	}

	tokens, err := svc.tokensRepository.GeneratePairOfTokens(userFetched.UserId)

	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (svc *Service) GenerateNewOTP(userInfo domain.Login) error {
	userFetched, err := svc.getUser(userInfo.PhoneNumber)

	if err != nil {
		return err
	}

	return svc.initProcessValidation(userFetched, userInfo.MethodVerificator)

}

func (svc *Service) getUser(phoneNumber string) (*domain.User, error) {
	user, err := svc.userRepository.GetUser(phoneNumber)

	if err != nil {
		switch err {
		case ports.ErrUserDoesNotExist:
			return nil, ports.ErrUserNotFound
		default:
			return nil, ports.ErrInternalError
		}
	}

	return user, err
}

func (svc *Service) isOTPStillAlive(user *domain.User) (bool, error) {
	latestOTPGenerationTime, err := svc.userRepository.GetLastOTPGenerationTimestap(user.PhoneNumber)

	if err != nil {
		return false, ports.ErrInternalError
	}

	if latestOTPGenerationTime != nil {
		logs.Info.Println(time.Since(*latestOTPGenerationTime))
		if time.Since(*latestOTPGenerationTime) < time.Second {
			logs.Info.Println("It is less tan a second")
		}
	} else {
		logs.Info.Println("NO OTP")
	}

	if latestOTPGenerationTime != nil && time.Since(*latestOTPGenerationTime) < time.Second {
		return true, nil
	}

	return false, nil
}

func (svc *Service) initProcessValidation(user *domain.User, methodVerificator domain.MethodVerifictor) error {
	isOTPAlive, err := svc.isOTPStillAlive(user)

	if err != nil {
		logs.Error.Println("a error occour, err: ", err.Error())
		return err
	}

	if isOTPAlive {
		logs.Error.Println("OTP is still alive")
		return ports.ErrOTPTTLStillLive
	}

	var otp []int = svc.generateOTP(svc.otpConifg.DefaultLength)

	if err := svc.userRepository.AddOTP(user.PhoneNumber, otp, svc.otpConifg.TTL); err != nil {
		return ports.ErrOnSavingOTP
	}

	var notificator ports.Notificator = svc.getNotificator(methodVerificator)

	if err := notificator.SendNotification(svc.genOTPMessage(methodVerificator, otp), user.PhoneNumber); err != nil {
		return ports.ErrOnSendingOTP
	}

	return nil
}

func (svc *Service) getNotificator(notType domain.MethodVerifictor) ports.Notificator {
	if notType.SMS {
		return svc.notificators[domain.SMS]
	}

	if notType.WhatsApp {
		return svc.notificators[domain.WHATSAPP]
	}

	return svc.notificators[domain.DEFAULT]
}

func (svc *Service) genOTPMessage(notType domain.MethodVerifictor, otp []int) string {

	codestr := ""

	for _, code := range otp {
		codestr = codestr + fmt.Sprint(code)
	}

	switch notType.Language {
	case "en":
		return fmt.Sprintf(OTPEnglish, codestr)
	case "es":
		return fmt.Sprintf(OTPSpanish, codestr)
	}

	return fmt.Sprintf(OTPEnglish, codestr)
}

func (svc *Service) generateOTP(length int) []int {
	result := make([]int, length)

	for i := 0; i < length; i++ {
		result[i] = rand.Intn(10)
	}
	return result
}
