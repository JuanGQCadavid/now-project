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

// swagger:model ProxymityResult
type ProxymityResult struct {
	// The result set containing the differents spots that were founded for the giving request params.
	Result domain.Locations `json:"result"`

	// Session details for the actual response, containing the header name and value for future calls.
	// if sent into the request then the backend will not send the same spots + the new ones that where found in a next call.
	// example: "session_details": {
	//             "session_id": "SessionIdUUID",
	//             "header_name": "X-Now-Search-Session-Id",
	//             "ttl":"timeStamp"
	//         }
	SearchSessionResponse SearchSessionResponse `json:"search_session,omitempty"`

	// Only present if a error happend.
	OnError OnError `json:"on_error,omitempty"`
}

type OnError struct {
	Error   ErrorType `json:"error,omitempty"`
	ErrorId string    `json:"error_id,omitempty"`
	TraceId string    `json:"traceId,omitempty"`
}
