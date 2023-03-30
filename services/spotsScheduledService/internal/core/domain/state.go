package domain

type State struct {
	Status Status `json:"status"`
	Since  int64  `json:"since"`
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
	Conclude Status = "conclude"
)

type Status string
