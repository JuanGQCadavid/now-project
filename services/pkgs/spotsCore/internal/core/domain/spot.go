package domain

type Spot struct {
	EventInfo  Event  `json:"eventInfo,omitempty"`
	HostInfo   Person `json:"hostInfo,omitempty"`
	PlaceInfo  Place  `json:"placeInfo,omitempty"`
	TopicsInfo Topic  `json:"topicInfo,omitempty"`
}
