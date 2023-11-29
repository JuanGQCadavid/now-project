package main

import (
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/tokens"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	// repo := users.NewDynamoDBUserRepository("Users", session)

	// repo.CreateUser("+573013475995", "JuanGo")
	// repo.CreateUser("+573235237844", "Sofilongas")

	// otp := []int{2, 2, 2, 2, 5}
	// repo.AddOTP("+573013475995", otp, time.Duration(time.Hour*3))
	// fakeOTP := []int{1, 2, 3, 4, 5}
	// repo.ValidateOTP("+573013475995", fakeOTP)

	// ttl, _ := repo.GetLastOTPGenerationTimestap("+573013475995")

	// if ttl != nil {
	// 	logs.Info.Println(time.Now().Sub(*ttl))
	// 	logs.Info.Println(ttl)
	// } else {
	// 	logs.Warning.Println("Empty TTL")
	// }

	// user, err := repo.GetUser("+573013475995")

	// if err != nil {
	// 	logs.Error.Fatal(err.Error())
	// }
	// logs.Info.Println(user)

	// Tokens
	tokenRepo := tokens.NewDynamoDBTokensRepository("Tokens", session)
	token, _ := tokenRepo.GeneratePairOfTokens("JuanGo")
	logs.Info.Printf("%+v\n", token)

}
