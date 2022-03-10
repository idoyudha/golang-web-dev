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
	Zip     int
	Region  string
}

type hotels []hotel

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	data := hotels{
		hotel{
			Name:    "South California",
			Address: "California 1",
			City:    "Cal1",
			Zip:     11011,
			Region:  "Southern",
		},
		hotel{
			Name:    "Central California",
			Address: "California 2",
			City:    "Cal2",
			Zip:     11012,
			Region:  "Central",
		},
		hotel{
			Name:    "North California",
			Address: "California 3",
			City:    "Cal3",
			Zip:     11013,
			Region:  "Northern",
		},
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
