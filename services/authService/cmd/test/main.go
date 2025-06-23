package main

import (
	"github.com/JuanGQCadavid/now-project/services/authService/core/core/service"
	"github.com/JuanGQCadavid/now-project/services/authService/core/encrypters"
	"github.com/JuanGQCadavid/now-project/services/authService/core/user"
	"github.com/JuanGQCadavid/now-project/services/authService/core/utils"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	serviceTest()
}

func reposTest() {
	// session := session.Must(session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// }))

	// encryptor := encrypters.NewSimpleEncrypt()

	// if resp, err := encryptor.DecodeToken("YjgxYjUyMjQtZGUyNS00NjIzLWExNDAtNTlkMjkyNjI3ZjZhKzhjZGQzN2M2LWNmNWUtNDc5Yi1iMzczLTY2MjJkOTQ3YjlkMg=="); err != nil {
	// 	logs.Error.Fatalln(err.Error())
	// } else {
	// 	logs.Info.Printf("Token: %+v", resp)
	// }

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

	// tokensTableName := utils.Getenv("tokensTableName", "Tokens")
	// tokens := tokens.NewDynamoDBTokensRepository(tokensTableName, session)

	// if err := tokens.IsTokenValid(domain.Token{
	// 	TokenValue: "8cdd37c6-cf5e-479b-b373-6622d947b9d2",
	// 	TokenID:    "ed956a1f-1731-4400-b76b-0f3480d0df36",
	// }); err != nil {
	// 	logs.Error.Println("Token err:= !", err.Error())
	// } else {
	// 	logs.Info.Println("We did it!")
	// }
}

func serviceTest() {
	var (
		sess          *session.Session
		userTableName string
		usersIndex    string
		repo          *user.DynamoDBUserRepository
		encryptor     *encrypters.SimpleEncrypt
	)

	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	encryptor = encrypters.NewSimpleEncrypt([]byte("DEFAULT"))

	userTableName = utils.Getenv("UsersTable", "Users-staging")
	usersIndex = utils.Getenv("UsersIndexTable", "UserId-index")
	repo = user.NewDynamoDBUserRepository(userTableName, usersIndex, sess)

	appService := service.NewAuthService(encryptor, repo)

	if resp, err := appService.GetUserDetailsFromToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjoiZTg5MjU0MTEtNzFlZS00OTgxLThlOTYtYTBmM2ViZWU5MzFkIiwidXNlcklkIjoiYjM0MzlmOTItNWQzZS00NThlLThlMDItZDIyMjAyNDVjNTI0IiwidXNlck5hbWUiOiJKdWFuR1FDYWRhdmlkIiwidXNlclBob25lIjoiKzM3MjUzOTU2NTgxIn0.OBNj2ihfLDasbY5534eqF_ccCMhiKTybymsvbBElSTw"); err != nil {
		logs.Error.Fatalln(err.Error())
	} else {
		logs.Info.Printf("%+v", resp)
	}

}
