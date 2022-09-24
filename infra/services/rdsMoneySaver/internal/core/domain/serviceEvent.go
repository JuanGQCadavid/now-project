package domain

type Actions string

const (
	STOP  Actions = "STOP"
	START Actions = "START"
)

type ServiceEvent struct {
	Action     Actions  `json:"action"`
	DBInstance []string `json:"instances"`
}
