package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
}

type region struct {
	Region string
	Hotels []hotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main() {
	hotels := []region{
		region{
			Region: "Southern",
			Hotels: []hotel{
				hotel{"Ritz", "located here", "palo alto", "94303"},
				hotel{"Nobu", "located here", "palo alto", "94303"},
				hotel{"Westin", "located here", "palo alto", "94303"},
				hotel{"Crowne Plaza", "located here", "palo alto", "94303"},
			},
		},
		region{
			Region: "Northern",
			Hotels: []hotel{
				hotel{"Ritz", "located here", "palo alto", "94303"},
				hotel{"Nobu", "located here", "palo alto", "94303"},
				hotel{"Westin", "located here", "palo alto", "94303"},
				hotel{"Crowne Plaza", "located here", "palo alto", "94303"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
