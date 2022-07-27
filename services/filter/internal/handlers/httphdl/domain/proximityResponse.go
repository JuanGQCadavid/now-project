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
}
