package handler

type ErrorMessage struct {
	Id            string `json:"id,omitempty"`
	Message       string `json:"message,omitempty"`
	InternalError string `json:"internalError,omitempty"`
}
