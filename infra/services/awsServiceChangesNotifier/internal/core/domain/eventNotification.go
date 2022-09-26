package domain

type EventDetail struct {
	SourceArn        string `json:"SourceArn"`
	Date             string `json:"Date"`
	Message          string `json:"Message"`
	SourceIdentifier string `json:"SourceIdentifier"`
	EventID          string `json:"EventID"`
}

type EventNotification struct {
	Id        string      `json:"id"`
	Account   string      `json:"account"`
	Time      string      `json:"time"`
	Region    string      `json:"region"`
	Resources []string    `json:"resources"`
	Detail    EventDetail `json:"detail"`
}
