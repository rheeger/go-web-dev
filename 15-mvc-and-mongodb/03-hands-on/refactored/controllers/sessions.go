package controllers

import (
	"net/http"
	"time"

	"github.com/rheeger/go-web-dev/15-mvc-and-mongodb/03-hands-on/refactored/models"
	uuid "github.com/satori/go.uuid"
)

func CreateSession(w http.ResponseWriter, r *http.Request) {
	// create session
	sID, _ := uuid.NewV4()
	c := &http.Cookie{
		Name:  "session",
		Value: sID.String(),
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	models.DbSessions[c.Value] = models.Session{un, time.Now()}
}
