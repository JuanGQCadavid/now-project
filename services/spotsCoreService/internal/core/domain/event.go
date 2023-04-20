package domain

import (
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
)

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

func (e *Event) IsEquals(ee *Event) bool {
	if &e == &ee {
		return true
	}

	if e.Description == ee.Description &&
		e.Emoji == ee.Emoji &&
		e.MaximunCapacty == ee.MaximunCapacty &&
		e.Name == ee.Name &&
		e.UUID == ee.UUID {
		return true
	}

	logs.Info.Println(e.Description == ee.Description)
	logs.Info.Println(e.Emoji == ee.Emoji)
	logs.Info.Println(e.Name == ee.Name)
	logs.Info.Println(e.UUID == ee.UUID)

	logs.Info.Println(e.UUID, ee.UUID)

	return false
}
