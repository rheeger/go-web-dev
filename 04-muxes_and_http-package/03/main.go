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
	http.Handle("/me", http.HandlerFunc(me))
	http.Handle("/dog", http.HandlerFunc(dog))
	http.Handle("/hello", http.HandlerFunc(heeger))

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
