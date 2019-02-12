package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "woof woof")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Robbie Heeger")

}

func main() {
	http.HandleFunc("/me", me)
	http.HandleFunc("/dog", dog)

	http.ListenAndServe(":8080", nil)
}
