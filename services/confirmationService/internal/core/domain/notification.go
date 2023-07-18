package domain

type Notification struct {
	DateId           string      `json:"dateId,omitempty"`
	SpotId           string      `json:"spotId,omitempty"`
	UserId           string      `json:"userId,omitempty"`
	Aditionalpayload interface{} `json:"aditionalpayload,omitempty"`
}
