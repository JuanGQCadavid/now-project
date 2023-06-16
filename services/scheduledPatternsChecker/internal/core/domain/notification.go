package domain

type Notification struct {
	DateId           string      `json:"scheduleId,omitempty"`
	UserId           string      `json:"userId,omitempty"`
	Aditionalpayload interface{} `json:"aditionalpayload,omitempty"`
}
