package main

import (
	"log"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"
	fakedata "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/fakeData"
	locationrepositories "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/locationRepositories"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gorm.io/gorm"
)

func main() {
	driver, err := locationrepositories.NewConector("admin", "admin", "pululapp", "localhost")

	if err != nil {
		logs.Error.Fatal("Ups error on Conector, err :", err.Error())
	}

	locationrepositories.Migrate(driver)

	// populateDummyData(driver)

	repo, err := locationrepositories.NewLocationRepoWithDriver(driver)

	if err != nil {
		logs.Error.Fatal("Ups error on repo, err :", err.Error())
	}

	pA := domain.LatLng{
		Lat: 6.19213630297576,
		Lng: -75.58816055664691,
	}

	pB := domain.LatLng{
		Lat: 6.289973745877467,
		Lng: -75.5995698004149,
	}

	result, err := repo.FetchSpotsIdsByArea(pA, pB)

	if err != nil {
		logs.Error.Println(err)
	}

	logs.Info.Println(len(result.Places))
	logs.Info.Println(result.Places[len(result.Places)-1])
}

func populateDummyData(driver *gorm.DB) {
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

		result := driver.Create(&locationrepositories.DatesLocation{
			DateID: fp.Id,
			Lat:    fp.LatLng.Lat,
			Lon:    fp.LatLng.Lng,
			Type: locationrepositories.Types{
				TypeID: "online", //locationrepositories.DateType(fp.Type),
			},
			State: locationrepositories.States{
				StateID: "online",
			},
		})

		if result.Error != nil {
			logs.Error.Fatal("Error in iserting", result.Error.Error())
		}

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

}
