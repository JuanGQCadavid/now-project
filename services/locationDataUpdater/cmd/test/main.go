package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatesLocation struct {
	DateID string  `gorm:"primaryKey"`
	Lat    float64 `gorm:"index"`
	Lon    float64 `gorm:"index"`

	TypeID string `gorm:"size:256"`
	Type   Types

	StateID string `gorm:"size:256"`
	State   States

	// GORM Variables
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type States struct {
	StateID     string `gorm:"primaryKey"`
	Description string

	// GORM Variables
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Types struct {
	TypeID      string `gorm:"primaryKey"`
	Description string

	// GORM Variables
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func main() {
	dsn := "admin:admin@tcp(localhost:3306)/pululapp?charset=utf8"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(&DatesLocation{}, &States{}, &Types{})

	date := DatesLocation{
		DateID: "TEST_123",
		Lat:    123.456,
		Lon:    789.123,
		Type: Types{
			TypeID:      "ONLINE",
			Description: "It is online",
		},
		State: States{
			StateID:     "ACTIVE",
			Description: "Active date",
		},
	}

	result := db.Create(&date) // pass pointer of data to Create

	fmt.Println(date)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)

	// //db.AutoMigrate(&DatesLocation{}, &State{}, &Type{})
}
