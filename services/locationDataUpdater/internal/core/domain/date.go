package domain

type DateStatus string

const (
	OnlineDateStatus  DateStatus = "online"
	StoppedDateStatus DateStatus = "stopped"
)

type DateType string

const (
	Online    DateType = "online"
	Scheduled DateType = "schedule"
)

type Date struct {
	DateId string
	Lat    float64
	Lon    float64
	Status DateStatus
	Type   DateType
}
