package domain

type Notification struct {
	SpotId           string      `json:"spotId,omitempty"`
	UserId           string      `json:"userId,omitempty"`
	Aditionalpayload interface{} `json:"aditionalpayload,omitempty"`
}
