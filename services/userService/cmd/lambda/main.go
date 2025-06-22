package main

import (
	"context"
	"os"

	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/services"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/handlers/httphdl"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/notificators/awssns"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/notificators/localnotificator"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/tokens"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/users"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

const (
	USER_TABLE_ENV_NAME   string = "usersTableName"
	USER_INDEX_ENV_NAME   string = "userIndexName"
	TOKENS_TABLE_ENV_NAME string = "tokensTableName"
	KEY_JWT_ENV_NAME      string = "jwtKey"
)

func init() {

	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	userTableName := getenv(USER_TABLE_ENV_NAME, "Users")
	tokensTableName := getenv(TOKENS_TABLE_ENV_NAME, "Tokens")
	userIndexName := getenv(USER_INDEX_ENV_NAME, "UserId-index")
	jwtKey := getenv(KEY_JWT_ENV_NAME, "DEFAULT")

	var userRepository ports.UserRepository = users.NewDynamoDBUserRepository(userTableName, userIndexName, session)
	var tokensRepository ports.TokensRepository = tokens.NewDynamoDBTokensRepository([]byte(jwtKey), tokensTableName, session)

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
	ginLambda = ginadapter.New(router)
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
