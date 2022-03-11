package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func a(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "default")
}

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog")
}

func m(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(res, "tpl.gohtml", "Ido")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func main() {
	http.HandleFunc("/", a)
	http.HandleFunc("/dog", d)
	http.HandleFunc("/me", m)

	http.ListenAndServe(":8090", nil)
}
