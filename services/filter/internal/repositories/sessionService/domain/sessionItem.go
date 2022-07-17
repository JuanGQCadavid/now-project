package domain

type SessionItem struct {
	SessionId string
	State     string
	TTL       int64 `dynamodbav:",omitempty"`
}
