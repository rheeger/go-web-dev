package models

import "time"

type Session struct {
	un           string
	lastActivity time.Time
}

var DbSessions = map[string]Session{} // session ID, session
var DbSessionsCleaned time.Time
