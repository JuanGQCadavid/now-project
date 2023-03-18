package domain

type SpotState struct {
	Confirmed bool       `json:"confirmed,omitempty"`
	Status    SpotStatus `json:"status,omitempty"`
	Since     int64      `json:"status_since,omitempty"`
}
