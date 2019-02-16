package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/rheeger/go-web-dev/15-mvc-and-mongodb/03-hands-on/refactored/controllers"
	"github.com/rheeger/go-web-dev/15-mvc-and-mongodb/03-hands-on/refactored/sessions"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type session struct {
	un           string
	lastActivity time.Time
}

var tpl *template.Template
var dbUsers = map[string]user{}       // user ID, user
var dbSessions = map[string]session{} // session ID, session
var dbSessionsCleaned time.Time

const sessionLength int = 30

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", sessions.Index)
	http.HandleFunc("/bar", sessions.Bar)
	http.HandleFunc("/signup", sessions.Signup)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/logout", sessions.Logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
