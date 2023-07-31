package utils

type BatchRequest struct {
	Operation  string `json:"Operation"`
	TimeWindow int64  `json:"TimeWindow"`
}
