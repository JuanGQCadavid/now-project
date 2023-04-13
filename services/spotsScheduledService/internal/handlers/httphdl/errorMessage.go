package httphdl

import "github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/domain"

type ErrorMessage struct {
	Id            string `json:"id,omitempty"`
	Message       string `json:"message,omitempty"`
	InternalError string `json:"internalError,omitempty"`
}

type TimeErrorsConflictsMessage struct {
	Id            string                `json:"id,omitempty"`
	Message       string                `json:"message,omitempty"`
	TimeConflicts []domain.TimeConflict `json:"timeConflicts,omitempty"`
}
