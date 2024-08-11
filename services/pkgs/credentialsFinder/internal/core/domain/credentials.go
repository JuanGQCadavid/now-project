package domain

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

type Neo4jCredentials struct {
	User     string
	Password string
	URI      string
}

type DBCredentials struct {
	User     string
	Password string
	Name     string
	Url      string
}

type CredentialsFinder interface {
	FindNeo4jCredentialsFromDefaultEnv() (neo4j.Driver, error)
	GetDBCredentialsFromDefaultEnv() (*DBCredentials, error)
}
