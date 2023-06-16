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
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/notifiers/dummy"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/notifiers/topics"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/repository/localrepository"
	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/repository/neo4jrepo"
)

type Result struct {
	Result []domain.Spot `json:"result,omitempty"`
}

const (
	TopicArnEnvName = "snsArn"
)

func main() {

	//manualCheck()
	//paralleCheck()
	// serviceManualTest()

	// testRepo()
	// generateDaysString(104)

	// testScheduleAdded2()
	testDeleteSchedulePattern()
	// geneateDatesFromRepoWithDrivers()

	// logs.Info.Println(domain.Monday | domain.Wednesday)
}
func generateDaysString(dayInt int) {
	day := domain.Day(dayInt)
	fmt.Println(dayInt)
	if domain.IsMonday(day) {
		fmt.Println("IsMonday")
	}
	if domain.IsTuesday(day) {
		fmt.Println("IsTuesday")
	}
	if domain.IsWednesday(day) {
		fmt.Println("IsWednesday")
	}
	if domain.IsThursday(day) {
		fmt.Println("IsThursday")
	}
	if domain.IsFriday(day) {
		fmt.Println("IsFriday")
	}
	if domain.IsSaturday(day) {
		fmt.Println("IsSaturday")
	}
	if domain.IsSunday(day) {
		fmt.Println("IsSunday")
	}
}

func getDrivers() (*neo4jrepo.Neo4jRepository, *queue.SQSConfirmation, *topics.Notifier) {
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

	notifier, err := topics.NewNotifierFromEnv(TopicArnEnvName)

	return repo, queueConfirmation, notifier
}

func testScheduleAdded() {
	repo, confirmation, notifier := getDrivers()
	srv := service.NewCheckerService(repo, confirmation, notifier, 1)
	spots := make([]domain.Spot, 2)

	sp := make([]domain.SchedulePattern, 1)
	sp[0] = domain.SchedulePattern{
		Id:        "25ea0c92-0df4-4028-b9de-6a65d0321996",
		HostId:    "33ddab28-006b-4790-bf42-1832f90dc8d4",
		StartTime: "13:00:00",
		Day:       domain.Day(104),
		EndTime:   "16:00:00",
		FromDate:  "2023-05-10",
		ToDate:    "2023-07-01",
	}

	spot := domain.Spot{
		SpotId:           "2248c84c-bb69-4aeb-85bc-5f6414260c6e",
		SchedulePatterns: sp,
	}

	spots[0] = spot

	sp2 := make([]domain.SchedulePattern, 1)
	sp2[0] = domain.SchedulePattern{
		Id:        "3ac21e4e-14d2-416e-9597-ce1cbb29d9fe",
		HostId:    "33ddab28-006b-4790-bf42-1832f90dc8d4",
		StartTime: "13:00:00",
		Day:       domain.Day(104),
		EndTime:   "16:00:00",
		FromDate:  "2023-05-10",
		ToDate:    "2023-07-01",
	}

	spot2 := domain.Spot{
		SpotId:           "93b3cd3c-1694-4a3f-aa92-ecb4effd79e7",
		SchedulePatterns: sp2,
	}

	spots[1] = spot2
	srv.OnSchedulePatternAppended(spots, 604800)

}

func testScheduleAdded2() {
	repo, confirmation, notifier := getDrivers()
	srv := service.NewCheckerService(repo, confirmation, notifier, 1)
	spots := make([]domain.Spot, 1)

	sp := make([]domain.SchedulePattern, 2)
	sp[0] = domain.SchedulePattern{
		Id:        "15ceaa61-a7f0-461b-b629-d4c88ba9b469",
		HostId:    "33ddab28-006b-4790-bf42-1832f90dc8d4",
		StartTime: "15:00:00",
		Day:       domain.Day(5),
		EndTime:   "18:00:00",
		FromDate:  "2023-05-10",
		ToDate:    "2023-07-01",
	}
	sp[1] = domain.SchedulePattern{
		Id:        "3ac21e4e-14d2-416e-9597-ce1cbb29d9fe",
		HostId:    "33ddab28-006b-4790-bf42-1832f90dc8d4",
		StartTime: "13:00:00",
		Day:       domain.Day(104),
		EndTime:   "16:00:00",
		FromDate:  "2023-05-10",
		ToDate:    "2023-07-01",
	}

	spot := domain.Spot{
		SpotId:           "93b3cd3c-1694-4a3f-aa92-ecb4effd79e7",
		SchedulePatterns: sp,
	}

	spots[0] = spot
	srv.OnSchedulePatternAppended(spots, 604800)
}

func testDeleteSchedulePattern() {
	repo, confirmation, notifier := getDrivers()
	srv := service.NewCheckerService(repo, confirmation, notifier, 1)

	sp := make([]string, 2)
	sp[0] = "3ac21e4e-14d2-416e-9597-ce1cbb29d9fe"
	sp[1] = "15ceaa61-a7f0-461b-b629-d4c88ba9b469"

	err := srv.DeleteScheduleDatesFromSchedulePattern(sp)

	if err != nil {
		logs.Error.Println("Errr ->", err.Error())
	}
}

func testRepo() {
	repo, _, _ := getDrivers()

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
	snsLocal := &dummy.Notifier{}

	if err != nil {
		logs.Error.Fatalln("error while creatin repo", err.Error())
	}

	// cores := runtime.NumCPU()
	srv := service.NewCheckerService(localRepo, queueConfirmation, snsLocal, 1)

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

func geneateDatesFromRepoWithDrivers() {
	repo, confirmation, notifier := getDrivers()
	// cores := runtime.NumCPU()
	srv := service.NewCheckerService(repo, confirmation, notifier, 1)

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
	notifier := &dummy.Notifier{}
	cores := runtime.NumCPU()
	srv := service.NewCheckerService(localRepo, localconfirmation, notifier, cores)

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
