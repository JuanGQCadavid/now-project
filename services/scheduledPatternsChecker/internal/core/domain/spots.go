package domain

type Spot struct {
	SpotId           string            `json:"spotId,omitempty"`
	SchedulePatterns []SchedulePattern `json:"schedulePatterns,omitempty"`
}

type SchedulePattern struct {
	Id          string  `json:"patternId,omitempty"`
	HostId      string  `json:"hostId,omitempty"`
	Day         Day     `json:"day,omitempty"`
	FromDate    string  `json:"fromDate,omitempty"`
	ToDate      string  `json:"toDate,omitempty,omitempty"`
	StartTime   string  `json:"StartTime,omitempty"`
	EndTime     string  `json:"endTime,omitempty"`
	CheckedUpTo int64   `json:"checkedUntil,omitempty"`
	Dates       []Dates `json:"dates,omitempty"`
}

const (
	Monday Day = 1 << iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

var Days []Day = []Day{
	Monday,
	Tuesday,
	Wednesday,
	Thursday,
	Friday,
	Saturday,
	Sunday,
}

type Day uint

type SpotPatternsDeep struct {
	Spot Spot
	Deep int
}

type Dates struct {
	DateId                        string `json:"dateId"`
	DurationApproximatedInSeconds int64  `json:"durationApproximated"`
	StartTime                     string `json:"startTime,omitempty"`
	Date                          string `json:"date,omitempty"`
	MaximunCapacty                int64  `json:"maximunCapacty,omitempty"`
	HostId                        string `json:"hostId,omitempty"`
	CreatedFromSchedulePatternId  string `json:"createdFromSchedulePatternId,omitempty"`
}

func IsMonday(day Day) bool {
	return (day & Monday) == Monday
}

func IsTuesday(day Day) bool {
	return (day & Tuesday) == Tuesday
}

func IsWednesday(day Day) bool {
	return (day & Wednesday) == Wednesday
}

func IsThursday(day Day) bool {
	return (day & Thursday) == Thursday
}

func IsFriday(day Day) bool {
	return (day & Friday) == Friday
}

func IsSaturday(day Day) bool {
	return (day & Saturday) == Saturday
}

func IsSunday(day Day) bool {
	return (day & Sunday) == Sunday
}
