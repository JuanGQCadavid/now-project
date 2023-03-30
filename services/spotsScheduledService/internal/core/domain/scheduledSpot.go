package domain

type ScheduledSpot struct {
	SpotInfo SpotInfo          `json:"spotInfo"`
	Patterns []SchedulePattern `json:"patterns"`
}
