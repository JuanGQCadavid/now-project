package fakedata

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"
	"github.com/google/uuid"
)

type DummyDataGenerator struct {
	size         int32
	centralPoint domain.LatLng
	radious      float64
	spots        []domain.SimpleSpot
}

func NewDummyDataGenerator(size int32, centralPoint domain.LatLng, radious float64) *DummyDataGenerator {

	spots := make([]domain.SimpleSpot, size)

	return &DummyDataGenerator{
		size:         size,
		centralPoint: centralPoint,
		radious:      radious,
		spots:        spots,
	}
}

func (gen *DummyDataGenerator) GetAllData() []domain.SimpleSpot {
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

	randomCoords := [][2]float64{}

	for range gen.spots {
		innerarray := [2]float64{
			gen.radious * rand.Float64(),
			gen.radious * rand.Float64(),
		}
		randomCoords = append(randomCoords, innerarray)
	}

	for index, _ := range gen.spots {
		spot := domain.SimpleSpot{}

		// 1. Generate the UUIDS
		spot.Id = gen.NewUUID()

		// 2. Go through a For, split the elements into three families ( just iterate byt with a counter check it 1/3 for each one)
		startsAt := time.Now()

		if index%2 == 0 {
			delta := rand.Int63n(6) // Max 6 hours
			println("Delta in hours ->", delta)

			extraTime := time.Duration(delta) * time.Hour
			println("Extra time added ->", extraTime)

			startsAt = startsAt.Add(3*time.Hour + extraTime)
			println("final time -> ", startsAt.String())

		}
		// TODO -> with time.Unix we could get it in integer format.
		spot.StartTime = startsAt.Format(time.TimeOnly)

		// 3. With a random number assign a emoji from the emojis list and a type for the event.
		spot.Emoji = emojies[rand.Intn(len(emojies)-1)]

		// 4. Btwn the central point and the cp + radious generates lat and lot for each spot
		spot.LatLng = gen.splitData(index, len(gen.spots), randomCoords[index][0], randomCoords[index][1])

		log.Println(fmt.Sprintf("%+v\n", spot))
		gen.spots[index] = spot
	}

	println("Finish! the result is")
	log.Println(fmt.Sprintf("%+v\n", gen.spots))

}

func (gen *DummyDataGenerator) splitData(actualIndex int, maxIndex int, deltaX float64, deltaY float64) domain.LatLng {

	if actualIndex < maxIndex/4 {
		return domain.LatLng{
			Lat: gen.centralPoint.Lat + deltaX,
			Lng: gen.centralPoint.Lng + deltaY,
		}
	} else if actualIndex >= maxIndex/4 && actualIndex < ((maxIndex/4)*2) {
		return domain.LatLng{
			Lat: gen.centralPoint.Lat - deltaX,
			Lng: gen.centralPoint.Lng + deltaY,
		}
	} else if actualIndex >= ((maxIndex/4)*2) && actualIndex < ((maxIndex/4)*3) {
		return domain.LatLng{
			Lat: gen.centralPoint.Lat + deltaX,
			Lng: gen.centralPoint.Lng - deltaY,
		}
	} else {
		return domain.LatLng{
			Lat: gen.centralPoint.Lat - deltaX,
			Lng: gen.centralPoint.Lng - deltaY,
		}
	}
}

func (gen *DummyDataGenerator) NewUUID() string {
	return uuid.NewString()
}
