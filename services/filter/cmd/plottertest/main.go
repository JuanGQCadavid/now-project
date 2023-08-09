package main

import (
	"log"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"
	fakedata "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/fakeData"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func main() {
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
}
