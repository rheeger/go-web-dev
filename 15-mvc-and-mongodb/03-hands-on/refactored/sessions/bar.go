package sessions

import (
	"net/http"

	"github.com/rheeger/go-web-dev/15-mvc-and-mongodb/03-hands-on/refactored/controllers"
	"github.com/rheeger/go-web-dev/15-mvc-and-mongodb/03-hands-on/refactored/models"
)

func Bar(w http.ResponseWriter, req *http.Request) {
	u := controllers.GetUser(w, req)
	if !controllers.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	controllers.ShowSessions() // for demonstration purposes
	models.Tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
