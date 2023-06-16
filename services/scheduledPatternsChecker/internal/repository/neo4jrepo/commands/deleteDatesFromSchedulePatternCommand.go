package commands

import (
	_ "embed"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/ports"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type DeleteDatesFromSchedulePatternCommand struct {
	schedulePatternsIds []string
	dateStatus          string
}

func NewDeleteDatesFromSchedulePatternCommand(schedulePatternsIds []string, dateStatus string) *DeleteDatesFromSchedulePatternCommand {
	return &DeleteDatesFromSchedulePatternCommand{
		schedulePatternsIds: schedulePatternsIds,
		dateStatus:          dateStatus,
	}
}

var (
	//go:embed queries/deleteDatesFromSchedulePattern.cypher
	deleteDatesFromSchedulePattern string
)

func (cmd *DeleteDatesFromSchedulePatternCommand) Run(tr neo4j.Transaction) (interface{}, error) {

	queryParams := make(map[string]interface{})
	queryParams["SchedulePatterns"] = cmd.schedulePatternsIds
	queryParams["DateStatus"] = cmd.dateStatus

	result, err := tr.Run(deleteDatesFromSchedulePattern, queryParams)

	if err != nil {
		logs.Error.Println("Run command fail, error -> ", err.Error())
		return nil, ports.ErrFetchingData
	}

	datesId := make([]string, 0, 10)

	for result.Next() {
		dateId, err := cmd.processRecord(result.Record())

		if err != nil {
			logs.Error.Printf("We fail to process a record %+v \n", result.Record())
			logs.Error.Println("error -> ", err.Error())
			continue
		}

		datesId = append(datesId, dateId...)
	}

	return datesId, nil
}

func (cmd *DeleteDatesFromSchedulePatternCommand) processRecord(record *neo4j.Record) ([]string, error) {
	logs.Info.Printf("Recods: %+v\n", record)

	datesId, _ := record.Get("datesId")
	datesIdsArr := datesId.([]interface{})
	result := make([]string, len(datesIdsArr))

	for index, date := range datesIdsArr {
		result[index] = date.(string)
	}

	if datesId != nil {
		return result, nil
	}

	logs.Warning.Println("There were no dates deleted")

	return nil, nil
}
