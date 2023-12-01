package httphdl

import (
	"net/http"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/ports"
	"github.com/gin-gonic/gin"
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
	router.POST("/user/init/login", hdl.InitLoging)
	router.POST("/user/init/singup", hdl.InitSingUp)

	router.POST("/user/validate/:userID/otp", hdl.ValidateProcess)
	router.POST("/user/validate/:userID/otp/resent", hdl.GenerateNewOTP)
}

// user/init/login
func (hdl *UserServiceHandler) InitLoging(context *gin.Context) {
	loginRequest := domain.Login{}
	context.BindJSON(&loginRequest)

	logs.Info.Printf("\nHandler: InitLoging \n\tRequest: %+v", loginRequest)

	if len(loginRequest.PhoneNumber) == 0 {
		context.AbortWithStatusJSON(400, ErrorMessage{
			Message: "phoneNumber is required",
		})
		return
	}

	if err := hdl.userService.InitLogin(loginRequest); err != nil {

		if err == ports.ErrUserNotFound {
			context.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{
				Message:       "User does not exist",
				InternalError: err.Error(),
			})
			return
		}

		context.AbortWithStatusJSON(http.StatusInternalServerError, ErrorMessage{
			Message:       "Ups there is a internal problem",
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
			Message: "phoneNumber and userName are required",
		})
		return
	}

	if err := hdl.userService.InitSingUp(singUpRequest); err != nil {

		if err == ports.ErrUserIsAlreadyCreated {
			context.AbortWithStatusJSON(http.StatusConflict, ErrorMessage{
				Message:       "Phonenomber is already registered",
				InternalError: err.Error(),
			})
			return
		}

		if err == ports.ErrOTPTTLStillLive {
			context.AbortWithStatusJSON(http.StatusNotAcceptable, ErrorMessage{
				Message:       "OPT is still valid, you should wait",
				InternalError: err.Error(),
			})
			return
		}

		context.AbortWithStatusJSON(http.StatusInternalServerError, ErrorMessage{
			Message:       "Ups there is a internal problem",
			InternalError: err.Error(),
		})

		return
	}

	context.Status(http.StatusAccepted)
}

// user/otp/<userID>/validate
func (hdl *UserServiceHandler) ValidateProcess(context *gin.Context) {

}

// user/otp/<userId>/resent
// user/validare/<userId>/otp/resent
func (hdl *UserServiceHandler) GenerateNewOTP(context *gin.Context) {

}
