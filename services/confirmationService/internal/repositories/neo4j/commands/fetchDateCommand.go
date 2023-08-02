package commands

import (
	_ "embed"

	"github.com/JuanGQCadavid/now-project/services/confirmationService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
)

type FetchDateCommand struct {
	dateId string
}

var (
	//go:embed queries/fetchDate.cypher
	fetchDate string
)

func NewFetchDateCommand(dateId string) *FetchDateCommand {
	return &FetchDateCommand{
		dateId: dateId,
	}
}

func (cmd *FetchDateCommand) Run(trx neo4j.Transaction) (interface{}, error) {

	queryParams := map[string]interface{}{
		"dateId": cmd.dateId,
	}

	result, err := trx.Run(fetchDate, queryParams)

	if err != nil {
		logs.Error.Println("Error on Fetch Date Command trx.Run", err.Error())
		return nil, err
	}

	if result.Next() {
		record := result.Record()
		return cmd.transformRecord(record)

	} else {
		logs.Warning.Println("No records were found")
		return nil, nil
	}

}

func (cmd *FetchDateCommand) transformRecord(record *db.Record) (*domain.Date, error) {

	dateId, isThere := record.Get("dateUUID")

	if !(isThere && dateId != nil && len(dateId.(string)) > 0) {
		logs.Warning.Println("There were a record but missing dateId")
		return nil, nil
	}

	spotId, isThere := record.Get("eventUUID")
	if !isThere || spotId == nil {
		logs.Warning.Println("Missing dateStartTime")
	}

	dateStartTime, isThere := record.Get("dateStartTime")
	if !isThere || dateStartTime == nil {
		logs.Warning.Println("Missing dateStartTime")
	}

	dateConfirmed, isThere := record.Get("dateConfirmed")
	if !isThere || dateConfirmed == nil {
		logs.Warning.Println("Missing dateConfirmed")
	}

	dateDate, isThere := record.Get("dateDate")
	if !isThere || dateDate == nil {
		logs.Warning.Println("Missing dateDate")
	}

	hostId, isThere := record.Get("hostId")
	if !isThere || hostId == nil {
		logs.Warning.Println("Missing hostId")
	}

	hostName, isThere := record.Get("hostName")
	if !isThere || hostName == nil {
		logs.Warning.Println("Missing hostName")
	}

	status, isThere := record.Get("status")
	if !isThere || status == nil {
		logs.Warning.Println("Missing status")
	}

	placeLat, isThere := record.Get("placeLat")
	if !isThere || placeLat == nil {
		logs.Warning.Println("Missing placeLat")
	}

	placeLon, isThere := record.Get("placeLon")
	if !isThere || placeLon == nil {
		logs.Warning.Println("Missing placeLon")
	}

	placeMapProviderId, isThere := record.Get("placeMapProviderId")
	if !isThere || placeMapProviderId == nil {
		logs.Warning.Println("Missing placeMapProviderId")
	}

	placeName, isThere := record.Get("placeName")
	if !isThere || placeName == nil {
		logs.Warning.Println("Missing placeName")
	}

	return &domain.Date{
		StartTime: dateStartTime.(string),
		Confirmed: dateConfirmed.(bool),
		Id:        dateId.(string),
		SpotId:    spotId.(string),
		OnPlace: domain.Place{
			Name:          placeName.(string),
			Lat:           placeLat.(float64),
			Lon:           placeLon.(float64),
			MapProviderId: placeMapProviderId.(string),
		},
		DateStamp: dateDate.(string),
		Host: domain.Host{
			HostId:   hostId.(string),
			HostName: hostName.(string),
		},
		Status: status.(string),
	}, nil

}
