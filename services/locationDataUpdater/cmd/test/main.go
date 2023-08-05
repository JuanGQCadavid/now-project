package main

import (
	"github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/core/domain"
	locationrepositories "github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/repositories/locationRepositories"
)

// type DatesLocation struct {
// 	DateID string  `gorm:"primaryKey"`
// 	Lat    float64 `gorm:"index"`
// 	Lon    float64 `gorm:"index"`

// 	TypeID string `gorm:"size:256"`
// 	Type   Types

// 	StateID string `gorm:"size:256"`
// 	State   States

// 	// GORM Variables
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }

// type States struct {
// 	StateID     string `gorm:"primaryKey"`
// 	Description string

// 	// GORM Variables
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }

// type Types struct {
// 	TypeID      string `gorm:"primaryKey"`
// 	Description string

// 	// GORM Variables
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }

func main() {
	// dsn := "admin:admin@tcp(localhost:3306)/pululapp?charset=utf8"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	panic(err)
	// }

	// date := domain.DatesLocation{
	// 	DateID: "OTHER_TEST_123",
	// 	Lat:    123.456,
	// 	Lon:    789.123,
	// 	Type: domain.Types{
	// 		TypeID:      domain.Online,
	// 		Description: "It is online, babe",
	// 	},
	// 	State: domain.States{
	// 		StateID:     domain.OnlineDateStatus,
	// 		Description: "Active date",
	// 	},
	// }

	// result := db.Create(&date) // pass pointer of data to Create

	// fmt.Println(date)
	// fmt.Println(result.Error)
	// fmt.Println(result.RowsAffected)

	// //db.AutoMigrate(&DatesLocation{}, &State{}, &Type{})

	connector, err := locationrepositories.NewConector("admin", "admin", "pululapp", "localhost")
	connector.Migrate()

	if err != nil {
		panic(err)
	}

	connector.Migrate()

	repo, err := locationrepositories.NewLocationRepo(connector)

	if err != nil {
		panic(err)
	}

	// repo.CrateLocation(date)
	repo.UpdateLocationStatus("OTHER_TEST_123", domain.StoppedDateStatus)

}
