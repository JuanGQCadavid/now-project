package main

import (
	"os"

	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/services"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/handlers/httphdl"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/notificators/awssns"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/notificators/localnotificator"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/tokens"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/users"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
)

const (
	USER_TABLE_ENV_NAME string = "usersTableName"
	USER_INDEX_ENV_NAME string = "userIndexName"
	KEY_JWT_ENV_NAME    string = "jwtKey"
)

func main() {

	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	userTableName := getenv(USER_TABLE_ENV_NAME, "Users")
	userIndexName := getenv(USER_INDEX_ENV_NAME, "UserID-index")
	jwtKey := getenv(KEY_JWT_ENV_NAME, "DEFAULT")

	var userRepository ports.UserRepository = users.NewDynamoDBUserRepository(userTableName, userIndexName, session)
	var tokensRepository ports.TokensRepository = tokens.NewJWTTokenGenerator([]byte(jwtKey))

	var defaultNotificator ports.Notificator = localnotificator.LocalNotificator{}
	var snsNotificator ports.Notificator = awssns.NewSNSNotificator(session)

	var notificators map[domain.NotificatorType]ports.Notificator = map[domain.NotificatorType]ports.Notificator{
		domain.WHATSAPP: defaultNotificator,
		domain.DEFAULT:  snsNotificator,
		domain.SMS:      snsNotificator,
	}

	var service ports.UserService = services.NewService(userRepository, notificators, tokensRepository)

	userService := httphdl.NewUserServiceHandler(service)

	router := gin.Default()
	userService.ConfigureRouter(router)

	router.Run("0.0.0.0:8000")
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
