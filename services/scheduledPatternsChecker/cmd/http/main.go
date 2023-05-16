package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/cmd/ssm"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/confirmation/localconfirmation"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/confirmation/queue"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/service"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/repository/localrepository"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/repository/neo4jrepo"
)

type Result struct {
	Result []domain.Spot `json:"result,omitempty"`
}

func main() {

	//manualCheck()
	//paralleCheck()
	// serviceManualTest()

	// testRepo()
	serviceManualTest2()
}

func testRepo() {
	credsFinder := ssm.NewSSMCredentialsFinder()

	neo4jDriver, err := credsFinder.FindNeo4jCredentialsFromDefaultEnv()

	if err != nil {
		logs.Error.Println("There were an error while attempting to create drivers")
		logs.Error.Fatalln(err.Error())
	}
	repo := neo4jrepo.NewNeo4jRepoWithDriver(neo4jDriver)

	reponse, err := repo.FetchActiveSchedulePatterns()

	if err != nil {
		logs.Error.Fatalln(err.Error())
	}

	for _, spot := range reponse {
		logs.Info.Printf("%+v \n", spot)
	}
}

func serviceManualTest() {
	localRepo := localrepository.NewLocalRepository(10, 10)
	queueConfirmation, err := queue.NewSQSConfirmationFromEnv("sqsConfirmationArn")

	if err != nil {
		logs.Error.Fatalln("error while creatin repo", err.Error())
	}

	// cores := runtime.NumCPU()
	srv := service.NewCheckerService(localRepo, queueConfirmation, 1)

	result, err := srv.GenerateDatesFromRepository(604800)

	if err != nil {
		fmt.Println("There is a error! ", err.Error())
		os.Exit(1)
	}

	f, _ := os.Create("/Users/personal/Pululapp/now-project/services/scheduledPatternsChecker/cmd/http/ouput.json")

	resultOutput := Result{
		Result: result,
	}

	defer f.Close()

	json, _ := json.MarshalIndent(resultOutput, "", "    ")
	fmt.Println(string(json))
	f.Write(json)
}

func serviceManualTest2() {
	credsFinder := ssm.NewSSMCredentialsFinder()

	neo4jDriver, err := credsFinder.FindNeo4jCredentialsFromDefaultEnv()

	if err != nil {
		logs.Error.Println("There were an error while attempting to create drivers")
		logs.Error.Fatalln(err.Error())
	}
	repo := neo4jrepo.NewNeo4jRepoWithDriver(neo4jDriver)
	queueConfirmation, err := queue.NewSQSConfirmationFromEnv("sqsConfirmationArn")

	if err != nil {
		logs.Error.Fatalln("error while creatin repo", err.Error())
	}

	// cores := runtime.NumCPU()
	srv := service.NewCheckerService(repo, queueConfirmation, 1)

	result, err := srv.GenerateDatesFromRepository(604800)

	if err != nil {
		fmt.Println("There is a error! ", err.Error())
		os.Exit(1)
	}

	f, _ := os.Create("/Users/personal/Pululapp/now-project/services/scheduledPatternsChecker/cmd/http/ouput.json")

	resultOutput := Result{
		Result: result,
	}

	defer f.Close()

	json, _ := json.MarshalIndent(resultOutput, "", "    ")
	fmt.Println(string(json))
	f.Write(json)
}

func paralleCheck() {
	localRepo := localrepository.NewLocalRepository(10, 10)
	localconfirmation := localconfirmation.NewLocalConfirmation()
	cores := runtime.NumCPU()
	srv := service.NewCheckerService(localRepo, localconfirmation, cores)

	result, err := srv.GenerateDatesFromRepository(604800)

	if err != nil {
		fmt.Println("There is a error! ", err.Error())
		os.Exit(1)
	}

	f, _ := os.Create("/Users/personal/Pululapp/now-project/services/scheduledPatternsChecker/cmd/http/ouput.json")

	resultOutput := Result{
		Result: result,
	}

	defer f.Close()

	json, _ := json.MarshalIndent(resultOutput, "", "    ")
	fmt.Println(string(json))
	f.Write(json)
}

// func manualCheck() {
// 	localRepo := localrepository.NewLocalRepository(10, 10)
// 	localconfirmation := localconfirmation.NewLocalConfirmation()

// 	srv := service.NewCheckerService(localRepo, localconfirmation)

// 	patterns, _ := localRepo.FetchActiveSchedulePatterns()

// 	// logs.Info.Println("BEFORE")
// 	// fmt.Print("[")
// 	// for _, spot := range patterns {
// 	// 	fmt.Print(len(spot.SchedulePatterns), ", ")
// 	// }
// 	// fmt.Println("]")

// 	result := srv.GetSortedSpotsPatternsByDeep(patterns)

// 	logs.Info.Println("After")
// 	fmt.Print("[")
// 	for _, spot := range result {
// 		fmt.Print(len(spot.Spot.SchedulePatterns), ", ")
// 	}
// 	fmt.Println("]")
// 	cores := 4
// 	fmt.Println("cores -> ", cores)
// 	result2 := srv.SplitDatesPerCore(result, cores)

// 	fmt.Println("---")
// 	for _, spots := range result2 {
// 		counter := 0

// 		for _, spot := range spots {
// 			counter += len(spot.SchedulePatterns)
// 		}
// 		fmt.Print(counter, ", ")
// 	}
// 	fmt.Println()

// 	manualPatterns := make([]domain.Spot, 1)

// 	firstSP := make([]domain.SchedulePattern, 2)
// 	firstSP[0] = domain.SchedulePattern{
// 		Id:        "1",
// 		HostId:    "JUAN123",
// 		Day:       domain.Saturday,
// 		FromDate:  "2023-03-01",
// 		ToDate:    "2023-07-01",
// 		StartTime: "13:00:00",
// 		EndTime:   "16:00:00",
// 	}

// 	firstSP[1] = domain.SchedulePattern{
// 		Id:        "2",
// 		HostId:    "JUAN123",
// 		Day:       domain.Sunday | domain.Monday,
// 		FromDate:  "2023-03-01",
// 		ToDate:    "2023-07-01",
// 		StartTime: "13:00:00",
// 		EndTime:   "16:00:00",
// 	}

// 	// secondSP = make([]domain.Spot, 2)

// 	manualPatterns[0] = domain.Spot{
// 		SpotId:           "SpotId_1",
// 		SchedulePatterns: firstSP,
// 	}

// 	result3, _ := srv.GenerateDates(manualPatterns, 604800)
// 	// result3, _ := srv.GenerateDatesParallel()
// 	f, _ := os.Create("/Users/personal/Pululapp/now-project/services/scheduledPatternsChecker/cmd/http/ouput.json")

// 	resultOutput := Result{
// 		Result: result3,
// 	}

// 	defer f.Close()

// 	json, _ := json.MarshalIndent(resultOutput, "", "    ")
// 	fmt.Println(string(json))
// 	f.Write(json)

// }

func printLent(patterns []domain.Spot) {
	logs.Info.Print("[")
	for _, spot := range patterns {
		logs.Info.Print(len(spot.SchedulePatterns), ", ")
	}
	logs.Info.Print("]")
}
