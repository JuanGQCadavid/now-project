package domain

type Event struct {
	Name           string `json:"name"`
	UUID           string `json:"id"`          // Only fields with Cappital letter are exported / Visibles to the outside
	Description    string `json:"description"` // So as Json use it, we need to make ut public.
	MaximunCapacty int64  `json:"maximunCapacty"`
	EventType      string `json:"eventType"`
	Emoji          string `json:"emoji"`
}

const (
	OnlineEventType   = "online"
	ScheduleEventType = "scheduled"
)
