package domain

type Spot struct {
	EventInfo  Event  `json:"eventInfo"`
	HostInfo   Person `json:"hostInfo"`
	PlaceInfo  Place  `json:"placeInfo"`
	TopicsInfo Topic  `json:"topicInfo"`
}
