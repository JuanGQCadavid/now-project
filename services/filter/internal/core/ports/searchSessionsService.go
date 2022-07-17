package ports

import (
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain/session"
)

type SearchSessionService interface {
	/*
		Create a session only specifying the type of it,
		the SessionId and the Session TTL will be created internally
		and returned in the sessionConfig result.
	*/
	CreateSession(sessionType session.SessionTypes) (*session.SessionConfig, error)

	/*
		Add a list of spots ids to the session state as:
			Param/key name: Timestamp
			Value: ["spotId1", .. , "SpotId2"]
	*/
	AddSpotsToSession(sessionId string, sessionType session.SessionTypes, spotsIds []string) error

	/*
		Giving the sessionId and the session Type it will fecth the sessionf data
		SessionConfig (Sessionid, TTL, SessionType) + SessionData
		[("TimeStamp", ["SpotId1", "SpotId2" ...]), ("TimeStamp", ["SpotId1", "SpotId2" ...])]
	*/
	GetSessionData(sessionId string, sessionType session.SessionTypes) (session.SessionData, error)
}
