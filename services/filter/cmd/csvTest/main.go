// hello.go

package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/models"
	fakedata "github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/fakeData"
)

// Application struct
type Application struct {
	App      string
	Company  string
	Category string
}

func main() {

	cp := models.LatLng{
		Lat: 6.2409826,
		Lng: -75.5862183,
	}

	gen := fakedata.NewDummyDataGenerator(2000, cp, 0.05)
	gen.GeneratePoints()

	csvFile, err := os.Create("./data.csv")

	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)

	spots := gen.GetAllData()

	for _, usance := range spots {
		var row []string
		row = append(row, usance.Id)
		row = append(row, fmt.Sprintf("%f", usance.LatLng.Lat))
		row = append(row, fmt.Sprintf("%f", usance.LatLng.Lng))
		writer.Write(row)
	}

	var row []string
	row = append(row, "CP")
	row = append(row, fmt.Sprintf("%f", cp.Lat))
	row = append(row, fmt.Sprintf("%f", cp.Lng))
	writer.Write(row)

	// remember to flush!
	writer.Flush()
}