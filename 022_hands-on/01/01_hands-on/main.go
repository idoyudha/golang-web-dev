package main

import (
	"io"
	"net/http"
)

func a(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "default")
}

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog")
}

func m(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Ido")
}

func main() {
	http.HandleFunc("/", a)
	http.HandleFunc("/dog", d)
	http.HandleFunc("/me", m)

	http.ListenAndServe(":8090", nil)
}
