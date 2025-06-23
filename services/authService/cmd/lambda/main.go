package main

import (
	"context"
	"encoding/json"

	"github.com/JuanGQCadavid/now-project/services/authService/core/core/domain"
	"github.com/JuanGQCadavid/now-project/services/authService/core/core/service"
	"github.com/JuanGQCadavid/now-project/services/authService/core/encrypters"
	"github.com/JuanGQCadavid/now-project/services/authService/core/tokens"
	"github.com/JuanGQCadavid/now-project/services/authService/core/user"
	"github.com/JuanGQCadavid/now-project/services/authService/core/utils"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	appService *service.AuthService
)

const (
	POLICY_ALLOW              = "Allow"
	TOKENS_TABLE_ENV_NAME     = "TokensTable"
	USER_TABLE_ENV_NAME       = "UsersTable"
	USER_TABLE_INDEX_ENV_NAME = "UsersIndexTable"
	JWT_KEY_ENV_NAME          = "jwtKey"
)

func Handler(ctx context.Context, event events.APIGatewayV2CustomAuthorizerV2Request) (events.APIGatewayCustomAuthorizerResponse, error) {
	var (
		userDetails *domain.UserDetails = nil
		err         error               = nil
	)

	token := event.Headers[domain.APP_TOKEN]
	if token == domain.ANONYMOUS_KEY {
		return generatePolicy(domain.AnonymousUser.Name, POLICY_ALLOW, event.RouteArn, domain.AnonymousUser), nil
	}

	userDetails, err = appService.GetUserDetailsFromToken(token)

	if err != nil {
		logs.Error.Println("Auth service fail: err: ", err.Error())
	}

	if userDetails == nil || len(userDetails.UserID) == 0 {
		userDetails = domain.AnonymousUser
	}

	return generatePolicy(userDetails.Name, POLICY_ALLOW, event.RouteArn, userDetails), nil
}

func generatePolicy(principalId, effect, resource string, userDetails *domain.UserDetails) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalId}

	if effect != "" && resource != "" {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
	}

	if userDetails != nil {
		var inInterface map[string]interface{}
		inrec, _ := json.Marshal(userDetails)
		json.Unmarshal(inrec, &inInterface)
		authResponse.Context = inInterface
	}

	return authResponse
}

func init() {
	var (
		sess = session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
		tokensTableName = utils.Getenv(TOKENS_TABLE_ENV_NAME, "Tokens-staging")
		userTableName   = utils.Getenv(USER_TABLE_ENV_NAME, "Users-staging")
		usersIndex      = utils.Getenv(USER_TABLE_INDEX_ENV_NAME, "UserID-index")
		jwtKey          = utils.Getenv(JWT_KEY_ENV_NAME, "DEFAULT")

		token     *tokens.DynamoDBTokensRepository
		repo      *user.DynamoDBUserRepository
		encryptor *encrypters.SimpleEncrypt
	)

	encryptor = encrypters.NewSimpleEncrypt([]byte(jwtKey))

	token = tokens.NewDynamoDBTokensRepository(tokensTableName, sess, encryptor)
	repo = user.NewDynamoDBUserRepository(userTableName, usersIndex, sess)
	appService = service.NewAuthService(token, encryptor, repo)
}

func main() {
	lambda.Start(Handler)
}
