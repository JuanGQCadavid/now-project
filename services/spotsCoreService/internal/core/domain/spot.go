package domain

type Spot struct {
	EventInfo  Event  `json:"eventInfo,omitempty"`
	HostInfo   Person `json:"hostInfo,omitempty"`
	PlaceInfo  Place  `json:"placeInfo,omitempty"`
	TopicsInfo Topic  `json:"topicInfo,omitempty"`
	DateInfo   Date   `json:"dateInfo,omitempty"`
}

type Date struct {
	DateTime                      string `json:"dateTime,omitempty"`
	Id                            string `json:"id,omitempty"`
	DurationApproximatedInSeconds int64  `json:"durationApproximated,omitempty"`
	StartTime                     string `json:"startTime,omitempty"`
}
