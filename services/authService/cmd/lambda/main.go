package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JuanGQCadavid/now-project/services/authService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/authService/internal/core/service"
	"github.com/JuanGQCadavid/now-project/services/authService/internal/encrypters"
	"github.com/JuanGQCadavid/now-project/services/authService/internal/tokens"
	"github.com/JuanGQCadavid/now-project/services/authService/internal/user"
	"github.com/JuanGQCadavid/now-project/services/authService/internal/utils"
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
		err         error               = nil
	)
	fmt.Printf("The event: %+v\n", event)
	token := event.Headers["x-auth"]

	fmt.Println("token:", token)
	fmt.Println("Headers:", event.Headers)

	if len(token) > 0 {
		userDetails, err = appService.GetUserDetailsFromToken(token)
		if err != nil {
			logs.Error.Println("Auth service fail: err: ", err.Error())
			userDetails = nil
		}
	}

	return generatePolicy("user", "Allow", event.RouteArn, userDetails), nil

	// switch strings.ToLower(token) {
	// case "allow":
	// 	return generatePolicy("user", "Allow", event.RouteArn), nil
	// case "deny":
	// 	return generatePolicy("user", "Deny", event.RouteArn), nil
	// case "unauthorized":
	// 	return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized") // Return a 401 Unauthorized response
	// default:
	// 	return events.APIGatewayCustomAuthorizerResponse{}, errors.New("err: Invalid token")
	// }
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
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	tokensTableName := utils.Getenv("tokensTableName", "Tokens")
	tokens := tokens.NewDynamoDBTokensRepository(tokensTableName, session)

	userTableName := utils.Getenv("usersTableName", "Users")
	repo := user.NewDynamoDBUserRepository(userTableName, "UserId-index", session)

	encryptor := encrypters.NewSimpleEncrypt()

	appService = service.NewAuthService(tokens, encryptor, repo)
}

func main() {
	lambda.Start(Handler)
}
