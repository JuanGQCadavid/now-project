package fakedata

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/models"
	"github.com/google/uuid"
)

type DummyDataGenerator struct {
	size         int32
	centralPoint models.LatLng
	radious      float32
	spots        []models.Spot
}

func NewDummyDataGenerator(size int32, centralPoint models.LatLng, radious float32) *DummyDataGenerator {

	spots := make([]models.Spot, size)

	return &DummyDataGenerator{
		size:         size,
		centralPoint: centralPoint,
		radious:      radious,
		spots:        spots,
	}
}

func (gen *DummyDataGenerator) GetAllData() []models.Spot {
	return gen.spots
}

func (gen *DummyDataGenerator) GeneratePoints() {
	// Procedure
	// 1. Generate the UUIDS
	// 2. Go through a For, split the elements into three families ( just iterate byt with a counter check it 1/3 for each one)
	// 	1. Now
	// 	2. In a 3  hour window
	// 	3. Afther 3 Hour window
	// 3. With a random number assign a emoji from the emojis list and a type for the event.
	// 4. Btwn the central point and the cp + radious generates lat and lot for each spot
	// 5. Add the result to the spots array

	emojies := []string{
		":D", ":O", ":P", ":V", ":Q_",
	}

	randomCoords := [][2]float32{}

	for range gen.spots {
		innerarray := [2]float32{
			gen.radious * rand.Float32(),
			gen.radious * rand.Float32(),
		}
		randomCoords = append(randomCoords, innerarray)
	}

	for index, _ := range gen.spots {
		spot := models.Spot{}

		// 1. Generate the UUIDS
		spot.Id = gen.NewUUID()

		// 2. Go through a For, split the elements into three families ( just iterate byt with a counter check it 1/3 for each one)
		startsAt := time.Now()

		if index%2 == 0 {
			delta := rand.Int63n(3) // Max 3 hours
			println("Delta in hours ->", delta)

			extraTime := time.Duration(delta) * time.Hour
			println("Extra time added ->", extraTime)

			startsAt = startsAt.Add(3*time.Hour + extraTime)
			println("final time -> ", startsAt.String())

		}
		// TODO -> with time.Unix we could get it in integer format.
		spot.StartTime = startsAt.Format(time.RFC3339)

		// 3. With a random number assign a emoji from the emojis list and a type for the event.
		spot.Emoji = emojies[rand.Intn(len(emojies)-1)]

		// 4. Btwn the central point and the cp + radious generates lat and lot for each spot
		spot.LatLng = gen.splitData(index, len(gen.spots), randomCoords[index][0], randomCoords[index][1])

		fmt.Printf("%+v\n", spot)
		fmt.Println("")
		gen.spots[index] = spot
	}

	println("Finish! the result is")
	fmt.Printf("%+v\n", gen.spots)

}

func (gen *DummyDataGenerator) splitData(actualIndex int, maxIndex int, deltaX float32, deltaY float32) models.LatLng {

	if actualIndex < maxIndex/4 {
		return models.LatLng{
			Lat: gen.centralPoint.Lat + deltaX,
			Lng: gen.centralPoint.Lng + deltaY,
		}
	} else if actualIndex >= maxIndex/4 && actualIndex < ((maxIndex/4)*2) {
		return models.LatLng{
			Lat: gen.centralPoint.Lat - deltaX,
			Lng: gen.centralPoint.Lng + deltaY,
		}
	} else if actualIndex >= ((maxIndex/4)*2) && actualIndex < ((maxIndex/4)*3) {
		return models.LatLng{
			Lat: gen.centralPoint.Lat + deltaX,
			Lng: gen.centralPoint.Lng - deltaY,
		}
	} else {
		return models.LatLng{
			Lat: gen.centralPoint.Lat - deltaX,
			Lng: gen.centralPoint.Lng - deltaY,
		}
	}
}

func (gen *DummyDataGenerator) NewUUID() string {
	return uuid.NewString()
}
