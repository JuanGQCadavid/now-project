package main

import (
	"time"

	"github.com/JuanGQCadavid/now-project/services/userService/internal/users"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	repo := users.NewDynamoDBUserRepository("Users", session)

	// repo.CreateUser("+573013475995", "JuanGo")
	// repo.CreateUser("+573235237844", "Sofilongas")

	otp := []int{1, 2, 3, 4, 5}
	repo.AddOTP("+573013475995", otp, time.Duration(time.Hour*3))
}
