package domain

type Locations struct {
	Places []Spot `json:"places"`
}

type Spot struct {
	DateInfo   Date   `json:"dateInfo,omitempty"`
	EventInfo  Event  `json:"eventInfo,omitempty"`
	HostInfo   Person `json:"hostInfo,omitempty"`
	PlaceInfo  Place  `json:"placeInfo,omitempty"`
	TopicsInfo Topic  `json:"topicInfo,omitempty"`
}

type Date struct {
	DateTime                      string   `json:"dateTime,omitempty"`
	Type                          SpotType `json:"type,omitempty"`
	Id                            string   `json:"id,omitempty"`
	DurationApproximatedInSeconds int64    `json:"durationApproximated,omitempty"`
	StartTime                     string   `json:"startTime,omitempty"`
}
type Event struct {
	Name           string `json:"name,omitempty"`
	UUID           string `json:"id,omitempty"`          // Only fields with Cappital letter are exported / Visibles to the outside
	Description    string `json:"description,omitempty"` // So as Json use it, we need to make ut public.
	MaximunCapacty int64  `json:"maximunCapacty,omitempty"`
	Emoji          string `json:"emoji,omitempty"`
}

type Person struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Place struct {
	Name          string  `json:"name,omitempty"`
	Lat           float64 `json:"lat,omitempty"` // those nas are tags!
	Lon           float64 `json:"lon,omitempty"`
	MapProviderId string  `json:"mapProviderId,omitempty"`
}

type Topic struct {
	// deprecated
	Name            string   `json:"name,omitempty"`
	PrincipalTopic  string   `json:"principalTopic,omitempty"`
	SecondaryTopics []string `json:"secondaryTopics,omitempty"`
}
