package commands

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Command interface {
	Run(neo4j.Transaction) (interface{}, error)
}
