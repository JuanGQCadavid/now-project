package main

import (
	"fmt"
	"log"

	"github.com/JuanGQCadavid/now-project/services/spotsCore/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsCore/internal/core/services/spotsrv"
	"github.com/JuanGQCadavid/now-project/services/spotsCore/internal/repositories/neo4jRepository"
	spotactivityservices "github.com/JuanGQCadavid/now-project/services/spotsCore/internal/repositories/spotActivityServices"
	"github.com/JuanGQCadavid/now-project/services/spotsCore/pkg/uuidgen"
)

var (
	repoSpot     *neo4jRepository.Neo4jSpotRepo
	repoLocation *spotactivityservices.AWSSpotActivityTopic
)

func init() {

	repoSpot = neo4jRepository.NewNeo4jSpotRepo() //menRepository.New()
	repoLocation = spotactivityservices.NewAWSSpotActivityTopic()
}

func main() {
	uuid := uuidgen.New()
	service := spotsrv.New(repoSpot, repoLocation, uuid)

	var spotToCreate domain.Spot = domain.Spot{
		EventInfo: domain.Event{},
	}

	domainSpotResult, err := service.CreateSpot(spotToCreate)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(fmt.Sprintf("%+v", domainSpotResult))
}
