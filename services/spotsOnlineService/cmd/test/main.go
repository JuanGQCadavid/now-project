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

	out, err := service.Start("UID_A-1234", "1234556778899", 300, 10)

	if err != nil {
		log.Println("There were an error while fetching the data")
		log.Fatalln(err.Error())
	}

	log.Printf("RESULTAD := %+v", out)

}

// {
// 		SpotInfo:{
// 			SpotId:UID_A-1234
// 			OwnerId:1234556778899
// 		}
// 		DatesInfo:
// 			[
// 					{
// 						DateId:ce9c664e-cb34-47ad-a150-7f0c10466172
// 						DurationApproximatedInSeconds:300
// 						StartTime:2023-03-07 18:00:19.36997 -0500 -05 m=+1.361923168
// 						Date:2023-03-07 18:00:19.36997 -0500 -05 m=+1.361923168
// 						Confirmed:true
// 						MaximunCapacty:10
// 						HostInfo:{
// 							HostId: HostName:
// 						}
// 					}
// 				]
// 		PlaceInfo:{Name:I'M JUAN Lat:12.34 Lon:12.56 MapProviderId:PROV_ID_123}}
