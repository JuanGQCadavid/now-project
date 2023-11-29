package httphdl

import (
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
