package sessions

import (
	"net/http"

	"github.com/rheeger/go-web-dev/15-mvc-and-mongodb/03-hands-on/refactored/controllers"
	"github.com/rheeger/go-web-dev/15-mvc-and-mongodb/03-hands-on/refactored/models"
)

func Signup(w http.ResponseWriter, req *http.Request) {
	if controllers.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		u := controllers.CreateUser(w, req)
		controllers.ShowSessions() // for demonstration purposes
		models.Tpl.ExecuteTemplate(w, "signup.gohtml", u)
	}
	http.Redirect(w, req, "/", http.StatusSeeOther)

}
