package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rheeger/go-web-dev/015-mvc-and-mongodb/02-hands-on/models"
	uuid "github.com/satori/go.uuid"

	"github.com/julienschmidt/httprouter"
)

//UserController offers methods to the entire porgram to use with User-based actions
type UserController struct {
	session map[string]models.User
}

//NewUserController creates a new UserController for calling user-based methods.
func NewUserController(s map[string]models.User) *UserController {
	return &UserController{s}
}

//GetUser pulls information about a provided user and returns that user's JSON to the client.
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	u := uc.session[id]

	// Marshal provided interface into JSON structure
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

//CreateUser instantiates information about a provided user and returns that user's JSON to the client.
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id, _ = uuid.NewV4().String

	uc.session[u.Id] = u
	models.StoreUsers(uc.session)

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

//DeleteUser removes information about a provided user and confirms successful deletion to the client.
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	delete(uc.session, id)
	models.StoreUsers(uc.session)

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", id, "\n")
}
