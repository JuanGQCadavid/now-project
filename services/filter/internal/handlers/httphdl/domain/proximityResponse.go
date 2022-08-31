package domain

import "github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"

// {
//     "result": {
//     },
//     "search_session" : {
//         "session_details": {
//             "session_id": "SessionIdUUID",
//             "header_name": "X-Now-Search-Session-Id",
//             "ttl":"timeStamp"
//         }
//     }
// }
type ProxymityResult struct {
	Result                domain.Locations      `json:"result"`
	SearchSessionResponse SearchSessionResponse `json:"search_session,omitempty"`
	OnError               OnError               `json:"on_error,omitempty"`
}

type OnError struct {
	Error   ErrorType `json:"error,omitempty"`
	ErrorId string    `json:"error_id,omitempty"`
	TraceId string    `json:"traceId,omitempty"`
}
