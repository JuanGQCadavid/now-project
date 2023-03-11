package neo4j

import (
	"log"
	"os"
	"sync"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Neo4jRepoDriver struct {
	driver neo4j.Driver
}

var neo4jRepoDriver *Neo4jRepoDriver
var doOnce = &sync.Once{}

func GetNeo4jRepoDriver() *Neo4jRepoDriver {

	doOnce.Do(func() {
		neo4jUri, isPresentURL := os.LookupEnv("neo4jUri")
		neo4jUser, isPresentUser := os.LookupEnv("neo4jUser")
		neo4jPassword, isPresentPass := os.LookupEnv("neo4jPassword")

		if !isPresentURL || !isPresentUser || !isPresentPass {
			log.Println("neo4jUri: ", neo4jUri)
			log.Println("neo4jUser: ", neo4jUser)
			log.Println("neo4jPassword: ", neo4jPassword)
			log.Fatalln("The ULR, Password or Username is not present in the env.")
		}

		auth := neo4j.BasicAuth(neo4jUser, neo4jPassword, "")
		// You typically have one driver instance for the entire application. The
		// driver maintains a pool of database connections to be used by the sessions.
		// The driver is thread safe.
		driver, err := neo4j.NewDriver(neo4jUri, auth)
		if err != nil {
			panic(err)
		}
		neo4jRepoDriver = &Neo4jRepoDriver{
			driver: driver,
		}
	})

	return neo4jRepoDriver
}
