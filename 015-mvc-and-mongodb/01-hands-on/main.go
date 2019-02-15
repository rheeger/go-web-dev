package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rheeger/go-web-dev/015-mvc-and-mongodb/01-handson/controllers"
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
