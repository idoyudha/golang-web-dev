package main

import (
	"log"
	"os"
	"text/template"
)

type meal struct {
	Type string
	Item []string
}

type menus []meal

var tpl *template.Template

func init() {
	tpl = template.Must((template.ParseFiles("tpl.gohtml")))
}

func main() {
	data := menus{
		meal{
			Type: "Breakfast",
			Item: []string{
				"Meatball",
				"Indomie",
				"Bubur Ayam",
			},
		},
		meal{
			Type: "Lunch",
			Item: []string{
				"Penyetan",
				"Lemonilo",
				"Soto Ayam",
			},
		},
		meal{
			Type: "Dinner",
			Item: []string{
				"Nasi Goreng",
				"Tahu Telor",
				"Bakso",
			},
		},
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
