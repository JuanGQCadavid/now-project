package session

type SessionTypes string

const (
	SpotsReturned SessionTypes = "SPOTS_RETURNED"
	Empty         SessionTypes = "Empty"
)

type SessionConfig struct {
	TTL         int64
	SessionType SessionTypes
	SessionId   string
}

func NewSessionConfigWithData(ttl int64, sessionType SessionTypes, sessionId string) *SessionConfig {
	return &SessionConfig{
		TTL:         ttl,
		SessionType: sessionType,
		SessionId:   sessionId,
	}
}

func NewSessionConfig(sessionType SessionTypes) *SessionConfig {
	return &SessionConfig{
		TTL:         0,
		SessionType: sessionType,
		SessionId:   "",
	}
}
