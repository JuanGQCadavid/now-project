package domain

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
