package main

import (
	"github.com/JuanGQCadavid/now-project/services/authService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/authService/internal/tokens"
	"github.com/JuanGQCadavid/now-project/services/authService/internal/utils"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// userTableName := utils.Getenv("usersTableName", "Users")
	// // tokensTableName := utils.getenv("tokensTableName", "Tokens")
	// repo := user.NewDynamoDBUserRepository(userTableName, "UserId-index", session)

	// if resp, err := repo.GetUserData(domain.Token{
	// 	UserID: "b1b67749-483b-4b7a-b38d-fbec88120f92",
	// }); err != nil {
	// 	logs.Error.Println("We crash!", err.Error())
	// } else {
	// 	logs.Info.Printf("Resp: %+v\n", resp)
	// }

	tokensTableName := utils.Getenv("tokensTableName", "Tokens")
	tokens := tokens.NewDynamoDBTokensRepository(tokensTableName, session)

	if err := tokens.IsTokenValid(domain.Token{
		UserID:  "bcc1fc6e-aec9-4234-9e78-2c360219df70",
		TokenID: "ed956a1f-1731-4400-b76b-0f3480d0df36",
	}); err != nil {
		logs.Error.Println("Token err:= !", err.Error())
	} else {
		logs.Info.Println("We did it!")
	}
}
