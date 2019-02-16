package sessions

import (
	"net/http"
	"time"

	"github.com/rheeger/go-web-dev/15-mvc-and-mongodb/03-hands-on/refactored/controllers"
	"github.com/rheeger/go-web-dev/15-mvc-and-mongodb/03-hands-on/refactored/models"
)

func Logout(w http.ResponseWriter, req *http.Request) {
	if !controllers.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	// delete the session
	delete(models.DbSessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	// clean up dbSessions
	if time.Now().Sub(models.DbSessionsCleaned) > (time.Second * 30) {
		go controllers.CleanSessions()
	}

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
