package domain

type TimeConflict struct {
	SchedulePattern SchedulePattern   `json:"schedulePattern,omitempty"`
	ConflictWith    []SchedulePattern `json:"conflictWiths,omitempty"`
}
