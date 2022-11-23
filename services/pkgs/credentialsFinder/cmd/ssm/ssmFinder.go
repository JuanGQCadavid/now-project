package ssm

import (
	"database/sql"
	"log"

	"github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/internal/core/service"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
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

// 1. Fetch all parametes using fetchSSMParams
// 2. validated nonempty
// 3. return driversFinder
func (cf *SSMCredentialsFinder) FindNeo4jCredentials(neo4jUserSSMParam string, neo4jPasswordSSMParam string, neo4jUriSSMParam string) (neo4j.Driver, error) {

	response := cf.fetchSSMParams(&neo4jUserSSMParam, &neo4jPasswordSSMParam, &neo4jUriSSMParam)

	neo4jCreds := domain.Neo4jCredentials{
		User:     response[neo4jUserSSMParam],
		Password: response[neo4jPasswordSSMParam],
		URI:      response[neo4jUriSSMParam],
	}

	log.Printf("%+v\n", neo4jCreds)

	return service.FindNeo4jDriver(neo4jCreds)
}

func (cf *SSMCredentialsFinder) FindDBCredentials(dbUserSSMParam string, dbPasswordSSMParam string, dbNameSSMParam string, dbUrlSSMParam string) (*sql.DB, error) {

	response := cf.fetchSSMParams(&dbUserSSMParam, &dbPasswordSSMParam, &dbNameSSMParam, &dbUrlSSMParam)

	neo4jCreds := domain.DBCredentials{
		User:     response[dbUserSSMParam],
		Password: response[dbPasswordSSMParam],
		Url:      response[dbUrlSSMParam],
		Name:     response[dbNameSSMParam],
	}

	log.Printf("%+v\n", neo4jCreds)

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
