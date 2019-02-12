package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/dog.jpg", image)
	http.HandleFunc("/", foo)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func image(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./unknown.jpg")
}
