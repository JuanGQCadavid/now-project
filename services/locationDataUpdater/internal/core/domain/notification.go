package domain

type Notification struct {
	DateId           string                 `json:"dateId,omitempty"`
	UserId           string                 `json:"userId,omitempty"`
	SpotId           string                 `json:"spotId,omitempty"`
	Aditionalpayload map[string]interface{} `json:"aditionalpayload,omitempty"`
}
