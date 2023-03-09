package main

import (
	"fmt"
	"log"
	"time"

	"github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/cmd/ssm"
	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/core/services"
	"github.com/JuanGQCadavid/now-project/services/spotsOnlineService/internal/repositories/neo4j"
)

func main() {
	fiveMinutesDuration := time.Minute * 5
	fmt.Println(fiveMinutesDuration.Seconds())

	credsFinder := ssm.NewSSMCredentialsFinder()

	neo4jDriver, err := credsFinder.FindNeo4jCredentialsFromDefaultEnv()

	if err != nil {
		log.Println("There were an error while attempting to create drivers")
		log.Fatalln(err.Error())
	}

	neoRepo := neo4j.NewNeo4jRepoWithDriver(neo4jDriver)
	service := services.NewService(neoRepo)

	out, err := service.Start("UID_C-1234", "1234556778899", 300, 10)

	if err != nil {
		log.Println("There were an error while fetching the data")
		log.Fatalln(err.Error())
	}

	log.Printf("RESULTAD := %+v", out)

}

// {
// 	SpotInfo:
// 		{
// 			SpotId:UID_A-1234
// 			OwnerId:1234556778899
// 		}
// 	DatesInfo:
// 			[
// 				{
// 					DateId:d92f0f05-fe61-4c4e-a360-f0ce9c975ca1
// 					DurationApproximatedInSeconds:300
// 					StartTime:2023-03-09 08:15:10.970165 -0500 -05 m=+2.073805959
// 					Date:2023-03-09 08:15:10.970165 -0500 -05 m=+2.073805959
// 					Confirmed:true
// 					MaximunCapacty:10
// 					HostInfo:
// 						{
// 							HostId:1234556778899
// 							HostName:
// 						}
// 				}
// 			]
// 	PlaceInfo:
// 		{
// 			Name:I'M JUAN
// 			Lat:12.34
// 			Lon:12.56
// 			MapProviderId:PROV_ID_123}
// 		}
