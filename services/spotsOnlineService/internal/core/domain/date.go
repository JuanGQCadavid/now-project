package domain

type SpotDate struct {
	DateId                        string    `json:"dateId"`
	DurationApproximatedInSeconds int64     `json:"durationApproximated"`
	StartTime                     string    `json:"startTime,omitempty"`
	Date                          string    `json:"date,omitempty"`
	MaximunCapacty                int64     `json:"maximunCapacty,omitempty"`
	HostInfo                      HostInfo  `json:"hostInfo,omitempty"`
	State                         SpotState `json:"state,omitempty"`
}

func (spot *SpotDate) IsEmpty() bool {
	if len(spot.DateId) == 0 {
		return true
	}

	return false
}
