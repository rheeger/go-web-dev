package models

import (
	"encoding/json"
	"fmt"
	"os"
)

//User is a struct for submitting users via JSON over curl
type User struct {
	Id     string `json:"id" bson:"_id"`
	Name   string `json:"name" bson:"name"`
	Gender string `json:"gender" bson:"gender"`
	Age    int    `json:"age" bson:"age"`
}

//StoreUsers intializes the data for the application
func StoreUsers(m map[string]User) {
	f, err := os.Create("data")
	if err != nil {
		fmt.Println(err)

	}
	defer f.Close()

	json.NewEncoder(f).Encode(m)

}

//LoadUsers returns a map of all users.
func LoadUsers() map[string]User {
	m := make(map[string]User)

	f, err := os.Open("data")
	if err != nil {
		fmt.Println(err)
		return m
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(m)
	if err != nil {
		fmt.Println(err)
	}

	return m

}
