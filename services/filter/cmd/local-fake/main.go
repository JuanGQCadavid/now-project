package main

import (
	"fmt"
	"log"
	"time"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/services/filtersrv"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/handlers/httphdl"
	dbspotservicelambda "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/dbSpotServiceLambda"
	fakedata "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/fakeData"
	locationrepositories "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/locationRepositories"
	sessionservice "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/sessionService"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/gin-gonic/gin"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gorm.io/gorm"
)

var (
	dbDriver *gorm.DB
)

func main() {

	dbDriver, err := locationrepositories.NewConectorFromEnv()

	if err != nil {
		logs.Error.Println("There were an error while attempting to create drivers")
		logs.Error.Fatalln(err.Error())
	}

	locationrepositories.Migrate(dbDriver)
	dbDriver.AutoMigrate(&dbspotservicelambda.SpotsDB{})

	// TODO -> How can we return an error from an init method ?
	locationRepo, err := locationrepositories.NewLocationRepoWithDriver(dbDriver)

	if err != nil {
		logs.Error.Println("There were an error while attempting to create the repository")
		logs.Error.Fatalln(err.Error())
	}

	spotSrv, err := dbspotservicelambda.NewDBSpotServiceLambdaWithDriver(dbDriver)

	if err != nil {
		panic(err.Error())
	}

	filterSrv := filtersrv.New(locationRepo, spotSrv)
	sessionHdl := sessionservice.NewSearchSessionDynamoDbService()
	filterHandler := httphdl.NewHTTPHandler(filterSrv, sessionHdl)

	router := gin.Default()
	router.GET("/filter/proximity", filterHandler.FilterSpots)
	router.POST("/filter/populate", Populate)
	router.Static("/filter/swagger", "../../swagger")

	router.Run("0.0.0.0:8000")
}

func Populate(context *gin.Context) {

	dbDriver, err := locationrepositories.NewConectorFromEnv()

	// Generating data
	cp := domain.LatLng{
		Lat: 6.2409826,
		Lng: -75.5862183,
	}

	gen := fakedata.NewDummyDataGenerator(100, cp, 0.05)

	gen.GeneratePoints()

	fakeSpots := gen.GetAllData()

	xys := make([]plotter.XY, len(fakeSpots))

	for i, fp := range fakeSpots {
		xys[i] = plotter.XY{
			X: fp.LatLng.Lat,
			Y: fp.LatLng.Lng,
		}
	}

	p := plot.New()
	var pts plotter.XYs = xys
	scatter, err := plotter.NewScatter(pts)
	if err != nil {
		log.Panic(err)
	}
	p.Add(scatter)

	err = p.Save(500, 500, "plotLogo.png")
	if err != nil {
		log.Panic(err)
	}

	locationsToCreate := make([]locationrepositories.DatesLocation, len(fakeSpots))
	spotsToCreate := make([]dbspotservicelambda.SpotsDB, len(fakeSpots))

	counter := 0
	for i, point := range fakeSpots {

		if counter > 10 {
			counter = 0
		}
		counter++

		location := locationrepositories.DatesLocation{
			DateID: fmt.Sprintf("%s_date", point.Id),
			Lat:    point.LatLng.Lat,
			Lon:    point.LatLng.Lng,
			Type: locationrepositories.Types{
				TypeID: "online",
			},
			State: locationrepositories.States{
				StateID: "online",
			},
		}

		locationsToCreate[i] = location
		// result_location := dbDriver.Create(&location)

		// if result_location.Error != nil {
		// 	logs.Error.Fatalln("LOCATION -> An error ocoured!: ", result_location.Error)
		// }

		spot := dbspotservicelambda.SpotsDB{
			EventId:                           fmt.Sprintf("%s_event", point.Id),
			DateId:                            fmt.Sprintf("%s_date", point.Id),
			DateDateTime:                      time.Now().Format(time.DateOnly),
			DateDurationApproximatedInSeconds: 7200,
			DateStartTime:                     point.StartTime,
			EventName:                         fmt.Sprintf("%d - event", i),
			EventDescription:                  fmt.Sprintf("%d - Description", i),
			EventMaximunCapacty:               10,
			EventEmoji:                        point.Emoji,
			PersonId:                          fmt.Sprintf("%d-person-id", counter),
			PersonName:                        fmt.Sprintf("%d-person-name", counter),
			PlaceName:                         fmt.Sprintf("%d - place", i),
			PlaceLat:                          point.LatLng.Lat,
			PlaceLon:                          point.LatLng.Lng,
			PlaceMapProviderId:                fmt.Sprintf("%d", i),
			TopicPrincipalTopic:               fmt.Sprintf("%d-topic", counter),
		}
		spotsToCreate[i] = spot

		// result := dbDriver.Create(&spot)

		// if result.Error != nil {
		// 	logs.Error.Fatalln("An error ocoured!: ", result.Error)
		// }

		// logs.Info.Printf("%+v", result)
	}

	result := dbDriver.CreateInBatches(locationsToCreate, 10)

	if result.Error != nil {
		logs.Error.Fatalln("An error ocoured!: ", result.Error)
	}

	result = dbDriver.CreateInBatches(spotsToCreate, 10)

	if result.Error != nil {
		logs.Error.Fatalln("An error ocoured!: ", result.Error)
	}
}
