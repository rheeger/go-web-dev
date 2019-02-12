package main

import (
	"html/template"
	"io"
	"log"
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
	http.HandleFunc("/hello", heeger)

	http.ListenAndServe(":8080", nil)
}

func heeger(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", "Robbie Heeger")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
