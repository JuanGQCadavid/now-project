package httphdl

import (
	"net/http"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/ports"
	"github.com/gin-gonic/gin"

	authDomain "github.com/JuanGQCadavid/now-project/services/authService/core/core/domain"
	authUtils "github.com/JuanGQCadavid/now-project/services/authService/core/utils"
)

type UserServiceHandler struct {
	userService ports.UserService
}

func NewUserServiceHandler(userService ports.UserService) *UserServiceHandler {
	return &UserServiceHandler{
		userService: userService,
	}
}

func (hdl *UserServiceHandler) ConfigureRouter(router *gin.Engine) {
	router.GET("/profile/:userId", hdl.GetUserProfile)

	router.POST("/user/init/login", hdl.InitLoging)
	router.POST("/user/init/singup", hdl.InitSingUp)

	router.POST("/user/validate/otp", hdl.ValidateProcess)
	router.POST("/user/validate/otp/resent", hdl.GenerateNewOTP)
}

// /user/:userId
func (hdl *UserServiceHandler) GetUserProfile(context *gin.Context) {
	id := context.Param("userId")
	if len(id) == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{
			Message: "Missing userId path param",
		})
		return
	}

	var userDetails authDomain.UserDetails = *authUtils.GetHeaders(context.Request.Header)

	profile, err := hdl.userService.GetUserInfo(&userDetails, id)

	if err != nil {
		if err == ports.ErrUserNotFound {
			context.AbortWithStatus(http.StatusNotFound)
			return
		}
		context.AbortWithStatusJSON(http.StatusInternalServerError, ErrorMessage{
			Message:       "We face an error while fetching the profile",
			InternalError: err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, profile)
}

// user/init/login
func (hdl *UserServiceHandler) InitLoging(context *gin.Context) {
	loginRequest := domain.Login{}
	context.BindJSON(&loginRequest)

	logs.Info.Printf("\nHandler: InitLoging \n\tRequest: %+v", loginRequest)

	if len(loginRequest.PhoneNumber) == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{
			Message: "phoneNumber is required",
			Id:      BadBodyRequest,
		})
		return
	}

	if err := hdl.userService.InitLogin(loginRequest); err != nil {

		if err == ports.ErrUserNotFound {
			context.AbortWithStatusJSON(http.StatusNotFound, ErrorMessage{
				Message:       "User does not exist",
				Id:            UserNotFound,
				InternalError: err.Error(),
			})
			return
		}

		context.AbortWithStatusJSON(http.StatusInternalServerError, ErrorMessage{
			Message:       "Ups there is a internal problem",
			Id:            Internal,
			InternalError: err.Error(),
		})

		return
	}

	context.Status(http.StatusNoContent)

}

// user/init/singup
func (hdl *UserServiceHandler) InitSingUp(context *gin.Context) {
	singUpRequest := domain.SingUp{}
	context.BindJSON(&singUpRequest)

	logs.Info.Printf("\nHandler: InitSingUp \n\tRequest: %+v", singUpRequest)

	if len(singUpRequest.PhoneNumber) == 0 || len(singUpRequest.UserName) == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{
			Id:      BadBodyRequest,
			Message: "phoneNumber and userName are required",
		})
		return
	}

	if err := hdl.userService.InitSingUp(singUpRequest); err != nil {

		if err == ports.ErrUserIsAlreadyCreated {
			context.AbortWithStatusJSON(http.StatusConflict, ErrorMessage{
				Message:       "Phonenomber is already registered",
				Id:            PhoneNumberAlreadyTaken,
				InternalError: err.Error(),
			})
			return
		}

		if err == ports.ErrOTPTTLStillLive {
			context.AbortWithStatusJSON(http.StatusNotAcceptable, ErrorMessage{
				Message:       "OPT is still valid, you should wait",
				Id:            OTPAlive,
				InternalError: err.Error(),
			})
			return
		}

		context.AbortWithStatusJSON(http.StatusInternalServerError, ErrorMessage{
			Message:       "Ups there is a internal problem",
			Id:            Internal,
			InternalError: err.Error(),
		})

		return
	}

	context.Status(http.StatusAccepted)
}

// user/otp/<userID>/validate
func (hdl *UserServiceHandler) ValidateProcess(context *gin.Context) {

	validateProcess := domain.ValidateProcess{}
	context.BindJSON(&validateProcess)

	logs.Info.Printf("\nHandler: ValidateProcess \n\tRequest: %+v", validateProcess)

	if len(validateProcess.PhoneNumber) == 0 || len(validateProcess.Code) == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{
			Id:      BadBodyRequest,
			Message: "phoneNumber and Code are required",
		})
		return
	}

	tokens, err := hdl.userService.ValidateProcess(validateProcess)

	if err != nil {

		if err == ports.ErrUserNotFound {
			context.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{
				Message:       "User does not exist",
				Id:            UserNotFound,
				InternalError: err.Error(),
			})
			return
		}

		if err == ports.ErrInvalidOTP {
			context.AbortWithStatusJSON(http.StatusUnauthorized, ErrorMessage{
				Message:       "OPT is not valid",
				Id:            WrongOTP,
				InternalError: err.Error(),
			})
			return
		}

		if err == ports.ErrThereIsNotOTP {
			context.AbortWithStatusJSON(http.StatusPreconditionFailed, ErrorMessage{
				Message:       "There is not pending OTP, request a OTP first",
				Id:            NoPendingOTP,
				InternalError: err.Error(),
			})
			return
		}

		if err == ports.ErrMaxRetriesOnTOP {
			context.AbortWithStatusJSON(http.StatusGone, ErrorMessage{
				Message:       "OPT attemps reach the limit, request a new one",
				Id:            OTPMaxTriesReached,
				InternalError: err.Error(),
			})
			return
		}

		context.AbortWithStatusJSON(http.StatusInternalServerError, ErrorMessage{
			Id:            Internal,
			Message:       "Ups there is a internal problem",
			InternalError: err.Error(),
		})

		return
	}

	context.JSON(200, tokens)

}

// user/otp/<userId>/resent
// user/validare/<userId>/otp/resent
func (hdl *UserServiceHandler) GenerateNewOTP(context *gin.Context) {

	loginRequest := domain.Login{}
	context.BindJSON(&loginRequest)

	logs.Info.Printf("\nHandler: GenerateNewOTP \n\tRequest: %+v", loginRequest)

	if len(loginRequest.PhoneNumber) == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{
			Id:      BadBodyRequest,
			Message: "phoneNumber and Code are required",
		})
		return
	}

	if err := hdl.userService.GenerateNewOTP(loginRequest); err != nil {

		if err == ports.ErrUserNotFound {
			context.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{
				Message:       "User does not exist",
				Id:            UserNotFound,
				InternalError: err.Error(),
			})
			return
		}

		if err == ports.ErrOTPTTLStillLive {
			context.AbortWithStatusJSON(http.StatusLocked, ErrorMessage{
				Message:       "OPT is still alive",
				Id:            OTPAlive,
				InternalError: err.Error(),
			})
			return
		}

		context.AbortWithStatusJSON(http.StatusInternalServerError, ErrorMessage{
			Id:            Internal,
			Message:       "Ups there is a internal problem",
			InternalError: err.Error(),
		})

		return
	}

	context.Status(http.StatusNoContent)

}
