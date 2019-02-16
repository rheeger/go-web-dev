package sessions

import (
	"net/http"

	"github.com/rheeger/go-web-dev/15-mvc-and-mongodb/03-hands-on/refactored/controllers"
	"github.com/rheeger/go-web-dev/15-mvc-and-mongodb/03-hands-on/refactored/models"
)

func Index(w http.ResponseWriter, req *http.Request) {
	u := controllers.GetUser(w, req)
	controllers.ShowSessions() // for demonstration purposes
	models.Tpl.ExecuteTemplate(w, "index.gohtml", u)
}
