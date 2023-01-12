package neo4jRepository

import (
	"fmt"
	"log"

	"github.com/JuanGQCadavid/now-project/services/spotsCore/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsCore/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/spotsCore/internal/repositories/neo4jRepository/commands"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Neo4jSpotRepo struct {
	neo4jRepoDriver *Neo4jRepoDriver
	driver          neo4j.Driver
}

func NewNeo4jSpotRepo() *Neo4jSpotRepo {

	neo4jRepoDriver := GetNeo4jRepoDriver()
	return &Neo4jSpotRepo{
		driver: neo4jRepoDriver.driver,
	}
}

func NewNeo4jSpotRepoWithDriver(driver neo4j.Driver) *Neo4jSpotRepo {
	return &Neo4jSpotRepo{
		driver: driver,
	}
}

func (r Neo4jSpotRepo) Get(id string, format ports.OutputFormat) (domain.Spot, error) {
	println("Get id -> ", id)

	var command commands.Command

	switch format {
	case ports.FULL_FORMAT:
		command = commands.NewGetFullCommand(id)
	case ports.SMALL_FORMAT:
		command = commands.NewGetSmallCommand(id)
	default:
		command = commands.NewGetFullCommand(id)
	}

	session := r.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	records, err := session.ReadTransaction(func(tr neo4j.Transaction) (interface{}, error) {
		return command.Run(tr)
	})

	if err != nil {
		return domain.Spot{}, err
	}

	return *records.(*domain.Spot), nil
}

func (r Neo4jSpotRepo) GetSpots(spotIds []string, format ports.OutputFormat) (domain.MultipleSpots, error) {
	log.Println("Repository: GetSpots", fmt.Sprintf("%+v", spotIds))

	var command commands.Command

	switch format {
	case ports.FULL_FORMAT:
		command = commands.NewGetFullMultipleSpotsCommand(spotIds)
	case ports.SMALL_FORMAT:
		command = commands.NewGetSmallMultipleSpotsCommand(spotIds)
	default:
		command = commands.NewGetFullMultipleSpotsCommand(spotIds)
	}

	session := r.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	records, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return command.Run(tx)
	})

	if err != nil {
		return domain.MultipleSpots{}, err
	}

	return *records.(*domain.MultipleSpots), nil
}

func (r Neo4jSpotRepo) CreateSpot(spot domain.Spot) error {
	log.Printf("Repository: CreateSpot %+v \n", spot)

	var command commands.Command = commands.NewCreateSpotCommand(&spot)

	session := r.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return command.Run(tx)
	})

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r Neo4jSpotRepo) GetSpotByUserId(personId string) (domain.Spot, error) {
	return domain.Spot{}, nil
}
func (r Neo4jSpotRepo) EndSpot(spotId string) error {
	return nil
}

func (r Neo4jSpotRepo) CreateSpotTags(spotId string, principalTag domain.Optional, secondaryTags []string) error {
	log.Println("Repository: CreateSpotTags", "\nspotId: ", spotId, "\nprincipalTag: ", fmt.Sprintf("%+v", principalTag), "\nsecondaryTags: ", fmt.Sprintf("%+v", secondaryTags))

	if !principalTag.IsPresent() && (secondaryTags == nil || len(secondaryTags) == 0) {
		log.Println("Avoiding process as both Principal and secondary tags are empty.")
		return nil
	}

	session := r.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	var cmd commands.Command = commands.NewCreateSpotTagsCommand(spotId, principalTag, secondaryTags)

	output, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return cmd.Run(tx)
	})
	log.Println(fmt.Sprintf("%+v", output))
	return err
}
