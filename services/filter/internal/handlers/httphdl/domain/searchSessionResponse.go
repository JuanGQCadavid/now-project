package domain

// "search_session" : {
// 	"session_details": {
// 		"session_id": "SessionIdUUID",
// 		"header_name": "X-Now-Search-Session-Id",
// 		"ttl":"timeStamp"
// 	}
// }
const (
	DEFAULT_HEADER_NAME string = "X-Now-Search-Session-Id"
)

type SearchSessionResponse struct {
	SessionDetail SessionDetails `json:"session_details,omitempty"`
	OnError       OnError        `json:"on_error,omitempty"`
}

type ErrorType string

const (
	SessionNotFounded ErrorType = "Session not founded"
	SessionNotCreated ErrorType = "Session not created"
)

type SessionDetails struct {
	SessionId  string `json:"session_id,omitempty"`
	HeaderName string `json:"header_name,omitempty"`
	TTL        int64  `json:"ttl,omitempty"`
}

func NewSessionDetails(sessionId string, ttl int64) *SessionDetails {
	return &SessionDetails{
		SessionId:  sessionId,
		TTL:        ttl,
		HeaderName: DEFAULT_HEADER_NAME,
	}
}
