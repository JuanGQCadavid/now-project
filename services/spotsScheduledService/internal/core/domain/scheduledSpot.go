package domain

type ScheduledSpot struct {
	SpotInfo SpotInfo          `json:"spotInfo,omitempty"`
	Patterns []SchedulePattern `json:"patterns,omitempty"`
}
