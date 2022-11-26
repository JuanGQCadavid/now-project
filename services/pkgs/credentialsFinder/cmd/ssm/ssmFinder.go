package ssm

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/internal/core/service"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

const (
	standardNeo4jUserEnvName     = "neo4jUser"
	standardNeo4jPasswordEnvName = "neo4jPassword"
	standardNeo4jUriEnvName      = "neo4jUri"
	standardDbNameEnvName        = "dbName"
	standardDbPasswordEnvName    = "dbPassword"
	standardDbUrlEnvName         = "dbUrl"
	standardDbUserEnvName        = "dbUser"
)

var (
	ErrMissingEnvs = errors.New("missing envs for fetching the credentials")
)

type SSMCredentialsFinder struct {
	svc *ssm.SSM
}

func NewSSMCredentialsFinder() *SSMCredentialsFinder {
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := ssm.New(session)

	return &SSMCredentialsFinder{
		svc: svc,
	}
}

func (cf *SSMCredentialsFinder) FindNeo4jCredentialsFromDefaultEnv() (neo4j.Driver, error) {
	neo4jUri, isPresentURL := os.LookupEnv(standardNeo4jUriEnvName)
	neo4jUser, isPresentUser := os.LookupEnv(standardNeo4jUserEnvName)
	neo4jPassword, isPresentPass := os.LookupEnv(standardNeo4jPasswordEnvName)

	if !isPresentURL || !isPresentUser || !isPresentPass {
		log.Println("neo4jUri: ", neo4jUri)
		log.Println("neo4jUser: ", neo4jUser)
		log.Println("neo4jPassword: ", neo4jPassword)
		log.Println("The ULR, Password or Username is not present in the env.")

		return nil, ErrMissingEnvs
	}

	return cf.FindNeo4jCredentials(neo4jUser, neo4jPassword, neo4jUri)

}

func (cf *SSMCredentialsFinder) FindDBCredentialsFromDefaultEnv() (*sql.DB, error) {
	dbName, isPresentName := os.LookupEnv(standardDbNameEnvName)
	dbPassword, isPresentPass := os.LookupEnv(standardDbPasswordEnvName)
	dbUrl, isPresentUrl := os.LookupEnv(standardDbUrlEnvName)
	dbUser, isPresentUser := os.LookupEnv(standardDbUserEnvName)

	if !isPresentName || !isPresentPass || !isPresentUrl || !isPresentUser {
		log.Println("dbName: ", dbName)
		log.Println("dbPassword: ", dbPassword)
		log.Println("dbUrl: ", dbUrl)
		log.Println("dbUser: ", dbUser)
		log.Println("The ULR, Password or Username is not present in the env.")

		return nil, ErrMissingEnvs
	}

	return cf.FindDBCredentials(dbUser, dbPassword, dbName, dbUrl)

}

// 1. Fetch all parametes using fetchSSMParams
// 2. validated nonempty
// 3. return driversFinder
func (cf *SSMCredentialsFinder) FindNeo4jCredentials(neo4jUserSSMParam string, neo4jPasswordSSMParam string, neo4jUriSSMParam string) (neo4j.Driver, error) {
	log.Printf("FindNeo4jCredentials -> neo4jUserSSMParam: %s, neo4jPasswordSSMParam: %s, neo4jUriSSMParam: %s", neo4jUserSSMParam, neo4jPasswordSSMParam, neo4jUriSSMParam)
	response := cf.fetchSSMParams(&neo4jUserSSMParam, &neo4jPasswordSSMParam, &neo4jUriSSMParam)

	neo4jCreds := domain.Neo4jCredentials{
		User:     response[neo4jUserSSMParam],
		Password: response[neo4jPasswordSSMParam],
		URI:      response[neo4jUriSSMParam],
	}

	//log.Printf("%+v\n", neo4jCreds)

	return service.FindNeo4jDriver(neo4jCreds)
}

func (cf *SSMCredentialsFinder) FindDBCredentials(dbUserSSMParam string, dbPasswordSSMParam string, dbNameSSMParam string, dbUrlSSMParam string) (*sql.DB, error) {
	log.Printf("FindNeo4jCredentials -> dbUserSSMParam: %s, dbPasswordSSMParam: %s, dbNameSSMParam: %s, dbUrlSSMParam: %s", dbUserSSMParam, dbPasswordSSMParam, dbNameSSMParam, dbUrlSSMParam)
	response := cf.fetchSSMParams(&dbUserSSMParam, &dbPasswordSSMParam, &dbNameSSMParam, &dbUrlSSMParam)

	neo4jCreds := domain.DBCredentials{
		User:     response[dbUserSSMParam],
		Password: response[dbPasswordSSMParam],
		Url:      response[dbUrlSSMParam],
		Name:     response[dbNameSSMParam],
	}

	//log.Printf("%+v\n", neo4jCreds)

	return service.FindDBDriver(&neo4jCreds)
}

func (cf *SSMCredentialsFinder) fetchSSMParams(params ...*string) map[string]string {

	response := make(map[string]string)

	input := &ssm.GetParametersInput{
		Names: params,
	}

	output, err := cf.svc.GetParameters(input)

	if err != nil {
		log.Println("ERROR -> ssm getParameters -> \n", err.Error())
	}

	for _, parameter := range output.Parameters {
		response[*parameter.Name] = *parameter.Value
	}
	return response
}
