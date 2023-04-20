package main

import (
	"fmt"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/services/spotsrv"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/repositories/neo4jRepository"
	spotactivityservices "github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/repositories/spotActivityServices"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/pkg/uuidgen"
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
		logs.Error.Fatal(err)
	}

	logs.Info.Println(fmt.Sprintf("%+v", domainSpotResult))
}
