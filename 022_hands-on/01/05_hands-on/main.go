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
	http.Handle("/", http.HandlerFunc(a))
	http.Handle("/dog", http.HandlerFunc(d))
	http.Handle("/me", http.HandlerFunc(m))

	http.ListenAndServe(":8090", nil)
}
