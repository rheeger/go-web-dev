package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.Handle("/pics/", http.StripPrefix("/pics/", http.FileServer(http.Dir("./starting-files/public/pics"))))
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func index(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("./starting-files/templates/index.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
