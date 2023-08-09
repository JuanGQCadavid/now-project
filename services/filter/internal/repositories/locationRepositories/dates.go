package locationrepositories

import (
	"time"

	"gorm.io/gorm"
)

type DateState string

const (
	OnlineDateStatus  DateState = "online"
	StoppedDateStatus DateState = "stopped"
)

type DateType string

const (
	Online    DateType = "online"
	Scheduled DateType = "schedule"
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
	StateID     DateState `gorm:"primaryKey"`
	Description string

	// GORM Variables
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Types struct {
	TypeID      DateType `gorm:"primaryKey"`
	Description string

	// GORM Variables
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
