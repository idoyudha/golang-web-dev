package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type status struct {
	Code     int
	Response string
}

type statuses []status

func main() {
	var data statuses
	rcvd := `[
        {
            "code": 200,
            "response": "StatusOK"
        },
        {
            "code": 301,
            "response": "StatusMovedPermanently"
        },
        {
            "code": 302,
            "response": "StatusFound"
        },
        {
            "code": 303,
            "response": "StatusSeeOther"
        },
        {
            "code": 307,
            "response": "StatusTemporaryRedirect"
        },
        {
            "code": 400,
            "response": "StatusBadRequest"
        },
        {
            "code": 401,
            "response": "StatusUnauthorized"
        },
        {
            "code": 402,
            "response": "StatusPaymentRequired"
        },
        {
            "code": 403,
            "response": "StatusForbidden"
        },
        {
            "code": 404,
            "response": "StatusNotFound"
        },
        {
            "code": 405,
            "response": "StatusMethodNotAllowed"
        },
        {
            "code": 418,
            "response": "StatusTeapot"
        },
        {
            "code": 500,
            "response": "StatusInternalServerError"
        }
    ]`
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}
	for _, val := range data {
		fmt.Printf("code = %d & response = %s \n", val.Code, val.Response)
	}
}
