package main

import (
	"os"

	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/services"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/handlers/httphdl"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/notificators/awssns"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/notificators/localnotificator"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/profile"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/tokens"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/users"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
)

const (
	USER_TABLE_ENV_NAME    string = "usersTableName"
	PROFILE_TABLE_ENV_NAME string = "userProfileTableName"
	USER_INDEX_ENV_NAME    string = "userIndexName"
	KEY_JWT_ENV_NAME       string = "jwtKey"
)

func main() {

	var (
		session = session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
		userTableName        = getenv(USER_TABLE_ENV_NAME, "Users")
		userProfileTableName = getenv(PROFILE_TABLE_ENV_NAME, "UserProfile-staging")
		userIndexName        = getenv(USER_INDEX_ENV_NAME, "UserID-index")
		jwtKey               = getenv(KEY_JWT_ENV_NAME, "DEFAULT")

		// Repositories
		userRepository     ports.UserRepository                         = users.NewDynamoDBUserRepository(userTableName, userIndexName, session)
		profileRepository  ports.ProfileRepository                      = profile.NewProfileRepositoryDynamoDB(userProfileTableName, session)
		tokensRepository   ports.TokensRepository                       = tokens.NewJWTTokenGenerator([]byte(jwtKey))
		defaultNotificator ports.Notificator                            = localnotificator.LocalNotificator{}
		snsNotificator     ports.Notificator                            = awssns.NewSNSNotificator(session)
		notificators       map[domain.NotificatorType]ports.Notificator = map[domain.NotificatorType]ports.Notificator{
			domain.WHATSAPP: defaultNotificator,
			domain.DEFAULT:  snsNotificator,
			domain.SMS:      snsNotificator,
		}

		// Service
		service ports.UserService = services.NewService(userRepository, notificators, tokensRepository, profileRepository)
	)

	userService := httphdl.NewUserServiceHandler(service)

	router := gin.Default()
	userService.ConfigureRouter(router)

	router.Run("0.0.0.0:8002")
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
