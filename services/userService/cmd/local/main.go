package main

import (
	"os"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/core/services"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/notificators/awssns"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/notificators/localnotificator"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/tokens"
	"github.com/JuanGQCadavid/now-project/services/userService/internal/users"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	service ports.UserService
)

func init() {
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	userTableName := getenv("usersTableName", "Users")
	tokensTableName := getenv("tokensTableName", "Tokens")
	userIndexName := getenv("userIndexName", "UserId-index")
	jwtKey := getenv("jwtKey", "DEFAULT")

	var userRepository ports.UserRepository = users.NewDynamoDBUserRepository(userTableName, userIndexName, session)
	var tokensRepository ports.TokensRepository = tokens.NewDynamoDBTokensRepository([]byte(jwtKey), tokensTableName, session)
	var defaultNotificator ports.Notificator = localnotificator.LocalNotificator{}
	var snsNotificator ports.Notificator = awssns.NewSNSNotificator(session)

	var notificators map[domain.NotificatorType]ports.Notificator = map[domain.NotificatorType]ports.Notificator{
		domain.WHATSAPP: defaultNotificator,
		domain.DEFAULT:  snsNotificator,
		domain.SMS:      snsNotificator,
	}

	service = services.NewService(userRepository, notificators, tokensRepository)
}

func main() {
	// Singup
	// singUp()

	// Login
	login()

	// genNewToken()
	// validate()
}

func genNewToken() {
	err := service.GenerateNewOTP(domain.Login{
		PhoneNumber: "+573235237844-2",
		MethodVerificator: domain.MethodVerifictor{
			Language: "en",
			WhatsApp: true,
		},
	})

	if err != nil {
		logs.Error.Println(err.Error())
	}
}

func validate() {
	tokens, err := service.ValidateProcess(domain.ValidateProcess{
		PhoneNumber: "+573235237844-2",
		Code:        []int{3, 6, 9, 4},
	})

	if err != nil {
		logs.Error.Println(err.Error())
	} else {
		logs.Info.Printf("%+v", tokens)
	}
}

func singUp() {

	// err := service.InitSingUp(domain.SingUp{
	// 	PhoneNumber: "+573235237844-3",
	// 	UserName:    "Sofia3",
	// 	MethodVerificator: domain.MethodVerifictor{
	// 		Language: "en",
	// 		WhatsApp: true,
	// 	},
	// })

	// if err != nil {
	// 	logs.Error.Println(err.Error())
	// }

	tokens, err := service.ValidateProcess(domain.ValidateProcess{
		PhoneNumber: "+573235237844-2",
		Code:        []int{5, 8, 9, 0},
	})

	if err != nil {
		logs.Error.Println(err.Error())
	} else {
		logs.Info.Printf("%+v", tokens)
	}

}

func login() {

	err := service.InitLogin(domain.Login{
		PhoneNumber: "+573137590102",
		MethodVerificator: domain.MethodVerifictor{
			Language: "en",
			SMS:      true,
		},
	})

	if err != nil {
		logs.Error.Println(err.Error())
	}

	//Validate process

	// tokens, err := service.ValidateProcess(domain.ValidateProcess{
	// 	PhoneNumber: "+573235237844",
	// 	Code:        []int{2, 9, 3, 7},
	// })

	// if err != nil {
	// 	logs.Error.Println(err.Error())
	// } else {
	// 	logs.Info.Printf("%+v", tokens)
	// }
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func repoTest() {
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
	tokenRepo := tokens.NewDynamoDBTokensRepository([]byte("DEFAULT"), "Tokens", session)
	token, _ := tokenRepo.GeneratePairOfTokens("JuanGo")
	logs.Info.Printf("%+v\n", token)
}
