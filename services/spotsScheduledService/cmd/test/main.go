package main

import (
	"fmt"

	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/logs"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/services"
	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/repositories/local"
)

func main() {
	testService()
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
			Day:       domain.Thursday, // domain.Monday | domain.Friday | domain.Saturday,
			FromDate:  "2007-03-01",
			ToDate:    "2007-07-01",
			StartTIme: "13:00:00",
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
		StartTIme: "14:00:00",
		EndTime:   "16:00:00",
	}

	b := domain.SchedulePattern{
		Day:       domain.Thursday | domain.Saturday | domain.Sunday,
		FromDate:  "2007-03-01",
		ToDate:    "2007-07-01",
		StartTIme: "13:00:00",
		EndTime:   "16:00:00",
	}

	overlaps, err := a.Overlaps(b)

	if err != nil {
		fmt.Println("There were an error", err.Error())
	} else {
		fmt.Println("Overlaps: ", overlaps)
	}

}
