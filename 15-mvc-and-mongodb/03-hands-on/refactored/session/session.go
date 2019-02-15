package main

import (
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, req *http.Request) User {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	// if the User exists already, get User
	var u User
	if s, ok := DbSessions[c.Value]; ok {
		s.lastActivity = time.Now()
		DbSessions[c.Value] = s
		u = DbUser[s.un]
	}
	return u
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := DbSessions[c.Value]
	if ok {
		s.lastActivity = time.Now()
		DbSessions[c.Value] = s
	}
	_, ok = DbUser[s.un]
	// refresh session
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return ok
}

func cleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	showSessions()              // for demonstration purposes
	for k, v := range DbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(DbSessions, k)
		}
	}
	DbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	showSessions()             // for demonstration purposes
}

// for demonstration purposes
func showSessions() {
	fmt.Println("********")
	for k, v := range DbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}
