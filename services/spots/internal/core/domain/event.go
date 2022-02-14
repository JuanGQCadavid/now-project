package domain

type Event struct {
	Name           string `json:"name,omitempty"`
	UUID           string `json:"id,omitempty"`          // Only fields with Cappital letter are exported / Visibles to the outside
	Description    string `json:"description,omitempty"` // So as Json use it, we need to make ut public.
	MaximunCapacty int64  `json:"maximunCapacty,omitempty"`
	EventType      string `json:"eventType,omitempty"`
	Emoji          string `json:"emoji,omitempty"`
}

const (
	OnlineEventType   = "online"
	ScheduleEventType = "scheduled"
)
