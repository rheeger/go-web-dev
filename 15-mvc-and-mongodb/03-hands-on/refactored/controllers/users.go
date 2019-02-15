package controllers

import (
	"net/http"
	"time"

	"github.com/rheeger/go-web-dev/15-mvc-and-mongodb/03-hands-on/refactored/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, req *http.Request) {
	var u models.User
	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")
		// username taken?
		if _, ok := models.DbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		//create session
		CreateSession(w, req)
		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u = models.User{un, bs, f, l, r}
		models.DbUsers[un] = u
		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
}

func Login(w http.ResponseWriter, req *http.Request) {
	if AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u models.User
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		// is there a username?
		u, ok := models.DbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		CreateSession(w, req)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "login.gohtml", u)
}

func AlreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := models.DbSessions[c.Value]
	if ok {
		s.lastActivity = time.Now()
		models.DbSessions[c.Value] = s
	}
	_, ok = models.DbUsers[s.un]
	// refresh session
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return ok
}
