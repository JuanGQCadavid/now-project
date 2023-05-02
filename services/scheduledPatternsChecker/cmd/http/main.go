package main

import (
	"fmt"
	"runtime"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/confirmation/localconfirmation"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/service"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/repository/localrepository"
)

func main() {
	localRepo := localrepository.NewLocalRepository(15, 30)
	localconfirmation := localconfirmation.NewLocalConfirmation()

	srv := service.NewCheckerService(localRepo, localconfirmation)

	patterns, _ := localRepo.FetchActiveSchedulePatterns()

	logs.Info.Println("BEFORE")
	fmt.Print("[")
	for _, spot := range patterns {
		fmt.Print(len(spot.SchedulePatterns), ", ")
	}
	fmt.Println("]")

	result := srv.GetSortedSpotsPatternsByDeep(patterns)

	logs.Info.Println("After")
	fmt.Print("[")
	for _, spot := range result {
		fmt.Print(len(spot.Spot.SchedulePatterns), ", ")
	}
	fmt.Println("]")
	cores := runtime.NumCPU()
	fmt.Println("cores -> ", cores)
	result2 := srv.SplitDatesPerCore(result, cores)

	fmt.Println("---")
	for _, spots := range result2 {
		counter := 0

		for _, spot := range spots {
			counter += len(spot.SchedulePatterns)
		}
		fmt.Print(counter, ", ")
	}
	fmt.Println()
}

func printLent(patterns []domain.Spot) {
	logs.Info.Print("[")
	for _, spot := range patterns {
		logs.Info.Print(len(spot.SchedulePatterns), ", ")
	}
	logs.Info.Print("]")
}
