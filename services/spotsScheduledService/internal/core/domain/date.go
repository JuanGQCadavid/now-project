package domain

type Date struct {
	MaximunCapacty                int64           `json:"maximunCapacty,omitempty"`
	DurationApproximatedInSeconds int64           `json:"durationApproximatedInSeconds,omitempty"`
	StartTime                     string          `json:"startTime,omitempty"`
	Confirmed                     bool            `json:"confirmed,omitempty"`
	Id                            string          `json:"id,omitempty"`
	DateStamp                     string          `json:"dateStamp,omitempty"`
	FromSchedulePattern           SchedulePattern `json:"fromSchedulePattern,omitempty"`
	State                         State           `json:"state,omitempty"`
	Host                          Host            `json:"host,omitempty"`
}
