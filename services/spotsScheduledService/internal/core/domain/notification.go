package domain

type Notification struct {
	ScheduleId       string      `json:"scheduleId,omitempty"`
	SpotId           string      `json:"spotId,omitempty"`
	UserId           string      `json:"userId,omitempty"`
	Aditionalpayload interface{} `json:"aditionalpayload,omitempty"`
}
