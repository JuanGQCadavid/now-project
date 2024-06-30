package main

import (
	"context"
	"encoding/json"
	"fmt"

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

func Handler(ctx context.Context, event events.APIGatewayV2CustomAuthorizerV2Request) (events.APIGatewayCustomAuthorizerResponse, error) {
	var (
		userDetails *domain.UserDetails = nil

		err error = nil
	)

	token := event.Headers[domain.APP_TOKEN]
	fmt.Println("token:", token)

	if token == domain.ANONYMOUS_KEY {
		return generatePolicy(domain.AnonymousUser.Name, "Allow", event.RouteArn, domain.AnonymousUser), nil
	}

	userDetails, err = appService.GetUserDetailsFromToken(token)

	if err != nil {
		logs.Error.Println("Auth service fail: err: ", err.Error())
	}

	if userDetails == nil || len(userDetails.UserID) == 0 {
		userDetails = domain.AnonymousUser
	}

	return generatePolicy(userDetails.Name, "Allow", event.RouteArn, userDetails), nil
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
		// Struct to map
		var inInterface map[string]interface{}
		inrec, _ := json.Marshal(userDetails)
		json.Unmarshal(inrec, &inInterface)

		authResponse.Context = inInterface
	}

	return authResponse
}

func init() {
	var (
		sess            *session.Session
		tokensTableName string
		userTableName   string
		usersIndex      string
		token           *tokens.DynamoDBTokensRepository
		repo            *user.DynamoDBUserRepository
		encryptor       *encrypters.SimpleEncrypt
	)

	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	encryptor = encrypters.NewSimpleEncrypt()

	tokensTableName = utils.Getenv("TokensTable", "Tokens-staging")
	token = tokens.NewDynamoDBTokensRepository(tokensTableName, sess, encryptor)

	userTableName = utils.Getenv("UsersTable", "Users-staging")
	usersIndex = utils.Getenv("UsersIndexTable", "UserID-index")
	repo = user.NewDynamoDBUserRepository(userTableName, usersIndex, sess)

	appService = service.NewAuthService(token, encryptor, repo)
}

func main() {
	lambda.Start(Handler)
}
