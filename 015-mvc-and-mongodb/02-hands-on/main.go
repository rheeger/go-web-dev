package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rheeger/go-web-dev/015-mvc-and-mongodb/02-hands-on/controllers"
	"github.com/rheeger/go-web-dev/015-mvc-and-mongodb/02-hands-on/models"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() map[string]models.User {
	return models.LoadUsers()
}
