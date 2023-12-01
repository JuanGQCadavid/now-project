package httphdl

import (
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
			context.AbortWithStatusJSON(400, ErrorMessage{
				Message:       "User does not exist",
				InternalError: err.Error(),
			})
			return
		}

		context.AbortWithStatusJSON(500, ErrorMessage{
			Message:       "Ups there is a internal problem",
			InternalError: err.Error(),
		})

		return
	}

	context.Status(204)

}

// user/init/singup
func (hdl *UserServiceHandler) InitSingUp(context *gin.Context) {

}

// user/otp/<userID>/validate
func (hdl *UserServiceHandler) ValidateProcess(context *gin.Context) {

}

// user/otp/<userId>/resent
// user/validare/<userId>/otp/resent
func (hdl *UserServiceHandler) GenerateNewOTP(context *gin.Context) {

}
