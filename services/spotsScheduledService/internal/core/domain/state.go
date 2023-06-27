package domain

type State struct {
	Status Status `json:"status,omitempty"`
	Since  int64  `json:"since,omitempty"`
}

type ScheduleStateFlags uint

const (
	ActivateFlag ScheduleStateFlags = 1 << iota
	FreezeFlag
	ConcludeFlag
)

const (
	ACTIVATE Status = "activate"
	FREEZE   Status = "freeze"
	CONCLUDE Status = "conclude"

	// Dates Status
	SCHEDULED Status = "SCHEDULED"
)

type Status string
