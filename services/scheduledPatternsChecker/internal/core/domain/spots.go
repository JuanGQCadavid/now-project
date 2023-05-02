package domain

type Spot struct {
	SpotId           string            `json:"spotId,omitempty"`
	SchedulePatterns []SchedulePattern `json:"schedulePatterns,omitempty"`
	Dates            Dates             `json:"dates,omitempty"`
}

type SchedulePattern struct {
	Id          string `json:"patternId,omitempty"`
	HostId      string `json:"hostId,omitempty"`
	Day         Day    `json:"day,omitempty"`
	FromDate    string `json:"fromDate,omitempty"`
	ToDate      string `json:"toDate,omitempty,omitempty"`
	StartTime   string `json:"StartTime,omitempty"`
	EndTime     string `json:"endTime,omitempty"`
	CheckedUpTo int64  `json:"checkedUntil,omitempty"`
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
}
