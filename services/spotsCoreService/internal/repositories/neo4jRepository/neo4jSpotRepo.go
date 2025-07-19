package neo4jRepository

import (
	"context"
	"fmt"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/repositories/neo4jRepository/commands"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/zerolog/log"
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

func (r Neo4jSpotRepo) GetUserEventRole(ctx context.Context, userId, eventId string) (*domain.Access, error) {
	var (
		logger                   = log.Ctx(ctx)
		cmd     commands.Command = commands.NewGetUserEventRoleCommand(ctx, userId, eventId)
		session                  = r.driver.NewSession(neo4j.SessionConfig{})
	)

	defer session.Close()

	record, err := session.ReadTransaction(func(tr neo4j.Transaction) (interface{}, error) {
		return cmd.Run(tr)
	})

	if err != nil {
		logger.Err(err).Msg("Repository crash!")
		return nil, err
	}

	return record.(*domain.Access), nil
}

func (r Neo4jSpotRepo) GetDateAttendantsWithRole(ctx context.Context, eventId, dateId string) ([]*domain.Access, error) {
	var (
		logger                   = log.Ctx(ctx)
		cmd     commands.Command = commands.NewGetDateAttendantsWithRoleCommand(ctx, dateId, eventId)
		session                  = r.driver.NewSession(neo4j.SessionConfig{})
	)

	defer session.Close()

	record, err := session.ReadTransaction(func(tr neo4j.Transaction) (interface{}, error) {
		return cmd.Run(tr)
	})

	if err != nil {
		logger.Err(err).Msg("Repository crash!")
		return nil, err
	}

	return record.([]*domain.Access), nil
}

func (r Neo4jSpotRepo) Get(id string, format ports.OutputFormat) (domain.Spot, error) {
	logs.Info.Println("Get id -> ", id)

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

func (r Neo4jSpotRepo) GetSpotsByDatesId(datesIds []string, format ports.OutputFormat) (domain.MultipleSpots, error) {
	logs.Info.Println("Repository: GetSpots", fmt.Sprintf("%+v", datesIds))

	var command commands.Command

	switch format {
	case ports.FULL_FORMAT:
		command = commands.NewGetFullMultipleSpotsCommand(datesIds)
	case ports.SMALL_FORMAT:
		command = commands.NewGetSmallMultipleSpotsCommand(datesIds)
	default:
		command = commands.NewGetFullMultipleSpotsCommand(datesIds)
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
	logs.Info.Printf("Repository: CreateSpot %+v \n", spot)

	var command commands.Command = commands.NewCreateSpotCommand(&spot)

	session := r.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return command.Run(tx)
	})

	if err != nil {
		logs.Error.Println(err)
		return err
	}

	return nil
}

func (r Neo4jSpotRepo) GetSpotByUserId(personId string) (domain.Spot, error) {
	return domain.Spot{}, nil
}

func (r Neo4jSpotRepo) DeleteSpot(spotId string) error {
	logs.Info.Println("Repository: DeleteSpot:", spotId)

	session := r.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	var cmd commands.Command = commands.NewAddSelfRelationship(spotId, domain.Deleted)

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return cmd.Run(tx)
	})

	return err
}

func (r Neo4jSpotRepo) CreateSpotTags(spotId string, principalTag domain.Optional, secondaryTags []string) error {
	logs.Info.Println("Repository: CreateSpotTags", "\nspotId: ", spotId, "\nprincipalTag: ", fmt.Sprintf("%+v", principalTag), "\nsecondaryTags: ", fmt.Sprintf("%+v", secondaryTags))

	if !principalTag.IsPresent() && (secondaryTags == nil || len(secondaryTags) == 0) {
		logs.Info.Println("Avoiding process as both Principal and secondary tags are empty.")
		return nil
	}

	session := r.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	var cmd commands.Command = commands.NewCreateSpotTagsCommand(spotId, principalTag, secondaryTags)

	output, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return cmd.Run(tx)
	})
	logs.Info.Println(fmt.Sprintf("%+v", output))
	return err
}

func (r *Neo4jSpotRepo) UpdateSpotEvent(spotEvent domain.Event, spotId string) error {
	logs.Info.Printf("Repository: UpdateSpotEvent %+v \n", spotEvent)

	var command commands.Command = commands.NewUpdateSpotEventCommand(&spotEvent, spotId)

	session := r.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	eventInterface, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return command.Run(tx)
	})

	if err != nil {
		logs.Error.Println(err)
		return err
	}

	eventUpdated := eventInterface.(*domain.Event)
	fmt.Printf("\n THE EVENT Updated was -> %+v \n", eventUpdated)

	if !spotEvent.IsEquals(eventUpdated) {
		return ports.ErrSpotUpdatedFail
	}

	return nil
}
func (r *Neo4jSpotRepo) UpdateSpotPlace(spotEvent domain.Place, spotId string) error { return nil }
func (r *Neo4jSpotRepo) UpdateSpotTopic(spotEvent domain.Topic, spotId string) error { return nil }
