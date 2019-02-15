package models

import (
	"html/template"
	"time"
)

type Session struct {
	Un           string
	LastActivity time.Time
}

var Tpl *template.Template
var DbSessions = map[string]Session{} // session ID, session
var DbSessionsCleaned time.Time

const SessionLength int = 30
