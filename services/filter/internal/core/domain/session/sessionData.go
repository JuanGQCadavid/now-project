package session

type SessionData struct {
	SessionConfiguration SessionConfig
}

type SearchSessionData struct {
	SessionData
	// Spots -> [("TimeStamp", ["SpotId1", "SpotId2" ...]), ("TimeStamp", ["SpotId1", "SpotId2" ...])]
	Spots map[string][]string
}
