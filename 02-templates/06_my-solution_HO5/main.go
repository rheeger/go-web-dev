package main

import (
	"log"
	"os"
	"text/template"
)

type food struct {
	Name     string
	Calories int
	Cost     float64
}

type meal struct {
	Name  string
	Items []food
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main() {
	menu := []meal{
		meal{
			Name: "Breakfast",
			Items: []food{
				food{"eggs benedict", 470, 12.65},
				food{"eggs benedict", 470, 12.65},
				food{"eggs benedict", 470, 12.65},
				food{"eggs benedict", 470, 12.65},
			},
		},
		meal{
			Name: "lunch",
			Items: []food{
				food{"eggs benedict", 470, 12.65},
				food{"eggs benedict", 470, 12.65},
				food{"eggs benedict", 470, 12.65},
				food{"eggs benedict", 470, 12.65},
			},
		},
		meal{
			Name: "Dinner",
			Items: []food{
				food{"eggs benedict", 470, 12.65},
				food{"eggs benedict", 470, 12.65},
				food{"eggs benedict", 470, 12.65},
				food{"eggs benedict", 470, 12.65},
			},
		},
	}

	err := tpl.Execute(os.Stdout, menu)
	if err != nil {
		log.Fatalln(err)
	}
}
