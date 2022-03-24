package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", cookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func cookie(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("my-cookie")
	log.Printf("cookie %v, ", cookie)
	log.Printf("err %v, ", err)
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
			Path:  "/",
		}
	}
	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)
}

// Using cookies, track how many times a user has been to your website domain.
