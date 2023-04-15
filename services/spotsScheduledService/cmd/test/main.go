package main

import (
	"fmt"
	"log"

	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/logs"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/services"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/repositories/local"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/repositories/neo4j/commands"
)

func main() {
	// testService()
	//testAppendCommand()
	log.Println(domain.Monday | domain.Wednesday | domain.Friday)
	log.Println(domain.Thursday | domain.Saturday | domain.Sunday)

	number := domain.Day(21)
	log.Println("IsMonday ", domain.IsMonday(number))
	log.Println("IsThursday ", domain.IsThursday(number))
	log.Println("IsWednesday ", domain.IsWednesday(number))
	log.Println("IsTuesday ", domain.IsTuesday(number))
	log.Println("IsFriday ", domain.IsFriday(number))
	log.Println("IsSaturday ", domain.IsSaturday(number))
	log.Println("IsSunday ", domain.IsSunday(number))

}

func testAppendCommand() {
	sp := make([]domain.SchedulePattern, 2)

	sp[0] = domain.SchedulePattern{
		Id: "1_UUID",
		State: domain.State{
			Status: domain.ACTIVATE,
			Since:  1004,
		},
		Day:       domain.Monday | domain.Wednesday | domain.Friday,
		FromDate:  "2007-02-01",
		ToDate:    "2007-06-01",
		StartTime: "14:00:00",
		EndTime:   "16:00:00",
	}

	sp[1] = domain.SchedulePattern{
		Id: "2_UUID",
		State: domain.State{
			Status: domain.ACTIVATE,
			Since:  1005,
		},
		Day:       domain.Thursday | domain.Saturday | domain.Sunday,
		FromDate:  "2007-03-01",
		ToDate:    "2007-07-01",
		StartTime: "13:00:00",
		EndTime:   "16:00:00",
	}

	scheduleSpot := domain.ScheduledSpot{
		SpotInfo: domain.SpotInfo{
			SpotId:  "1_SPOT_ID",
			OwnerId: "1_OWNER_ID",
		},
		Patterns: sp,
	}

	cmd := commands.NewAppendScheduleCommand(scheduleSpot)
	cmd.Run(nil)
}

func testService() {
	repo := &local.LocalRepository{}
	service := services.NewScheduledService(repo)
	patterns := []domain.SchedulePattern{
		{
			Host: domain.Host{
				HostId:   "JUAN",
				HostName: "JUAN",
			},
			Day:       domain.Monday | domain.Friday | domain.Saturday, // domain.Thursday, //
			FromDate:  "2007-03-01",
			ToDate:    "2007-07-01",
			StartTime: "13:00:00",
			EndTime:   "16:00:00",
		},
	}
	sp, conflicts, err := service.AppendSchedule("123", "JUAN", &patterns)

	if err != nil {
		logs.Error.Println(err.Error())
	}

	if conflicts != nil && len(*conflicts) > 0 {
		logs.Warning.Println("There are conflicts associaated to the request")
		logs.Warning.Printf("%+v \n", conflicts)
	}

	logs.Info.Printf("Sp: %+v \n", sp)
}

func testOverlaps() {

	a := domain.SchedulePattern{
		Day:       domain.Monday | domain.Wednesday | domain.Friday,
		FromDate:  "2007-02-01",
		ToDate:    "2007-06-01",
		StartTime: "14:00:00",
		EndTime:   "16:00:00",
	}

	b := domain.SchedulePattern{
		Day:       domain.Thursday | domain.Saturday | domain.Sunday,
		FromDate:  "2007-03-01",
		ToDate:    "2007-07-01",
		StartTime: "13:00:00",
		EndTime:   "16:00:00",
	}

	overlaps, err := a.Overlaps(b)

	if err != nil {
		fmt.Println("There were an error", err.Error())
	} else {
		fmt.Println("Overlaps: ", overlaps)
	}

}
