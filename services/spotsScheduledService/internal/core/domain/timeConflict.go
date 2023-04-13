package domain

type TimeConflict struct {
	SchedulePattern SchedulePattern   `json:"schedulePattern"`
	ConflictWith    []SchedulePattern `json:"conflictWiths"`
}
