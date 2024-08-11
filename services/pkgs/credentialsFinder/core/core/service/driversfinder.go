package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/core/core/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var (
	ErrDBSessionCreation = errors.New("unable to create the db session")
	ErrDBPing            = errors.New("unable to ping the db")
)

func FindNeo4jDriver(credentials domain.Neo4jCredentials) (neo4j.Driver, error) {
	log.Println("On FindNeo4jDriver")
	auth := neo4j.BasicAuth(credentials.User, credentials.Password, "")
	// You typically have one driver instance for the entire application. The
	// driver maintains a pool of database connections to be used by the sessions.
	// The driver is thread safe.
	driver, err := neo4j.NewDriver(credentials.URI, auth)
	if err != nil {
		log.Println("[ERROR] an error happen while creating Neo4j driver.")
		log.Println("[ERROR] ", err.Error())
		return nil, err
	}

	log.Println("Driver creation goes well")

	err = driver.VerifyConnectivity()

	if err != nil {
		log.Println("[ERROR] Neo4j verify connection fail.")
		log.Println("[ERROR] ", err.Error())
		return nil, err
	}
	log.Println("Driver tested sucessfully")
	return driver, nil

}

func FindDBDriver(creds *domain.DBCredentials) (*sql.DB, error) {
	log.Println("On FindDBDriver")
	var dataSourceConnection string = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", creds.User, creds.Password, creds.Url, creds.Name)
	sqlDb, err := sql.Open("mysql", dataSourceConnection)

	if err != nil {
		log.Println("[ERROR] We face a problem while creating the db session")
		log.Println("[ERROR] ", err.Error())
		return nil, ErrDBSessionCreation
	}
	log.Println("Connection goes well")
	err = sqlDb.Ping()

	if err != nil {
		log.Println("Ping error!")
		log.Println("[ERROR] ", err.Error())
		return nil, ErrDBPing
	}
	log.Println("Ping operation goes well")
	return sqlDb, nil
}
