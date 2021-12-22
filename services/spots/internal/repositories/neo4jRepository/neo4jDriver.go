package neo4jRepository

import (
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
		// Aura requires you to use "neo4j+s" protocol
		// (You need to replace your connection details, username and password)
		uri := "neo4j+s://c36227ce.databases.neo4j.io"
		auth := neo4j.BasicAuth("neo4j", "eKsL1TO0UVU2iblhGTi5fe5JYd6JVHNgDgmsADlZeb4", "")
		// You typically have one driver instance for the entire application. The
		// driver maintains a pool of database connections to be used by the sessions.
		// The driver is thread safe.
		driver, err := neo4j.NewDriver(uri, auth)
		if err != nil {
			panic(err)
		}
		neo4jRepoDriver = &Neo4jRepoDriver{
			driver: driver,
		}
	})

	return neo4jRepoDriver
}
