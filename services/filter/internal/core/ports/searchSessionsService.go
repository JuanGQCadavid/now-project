package ports

import "github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain/session"

type SearchSessionService interface {
	CreateSession(session.SessionTypes) (*session.SessionConfig, error)
	GetSessionData(string, session.SessionTypes) (session.SessionData, error)
}
