package controllers

import (
	"fmt"
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
	c.MaxAge = models.SessionLength
	http.SetCookie(w, c)
	models.DbSessions[c.Value] = models.Session{"Un", time.Now()}
}

func ShowSessions() {
	fmt.Println("********")
	for k, v := range models.DbSessions {
		fmt.Println(k, v.Un)
	}
	fmt.Println("")
}

func CleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	ShowSessions()              // for demonstration purposes
	for k, v := range models.DbSessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			delete(models.DbSessions, k)
		}
	}
	models.DbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	ShowSessions()             // for demonstration purposes
}
