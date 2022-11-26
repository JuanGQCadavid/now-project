package main

import (
	"github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/cmd/ssm"
)

func main() {
	ssmFinder := ssm.NewSSMCredentialsFinder()
	ssmFinder.FindNeo4jCredentials("neo4jUser", "neo4jPassword", "neo4jUri")
	ssmFinder.FindDBCredentials("dbUser", "dbPassword", "dbName", "dbUrl")
}
