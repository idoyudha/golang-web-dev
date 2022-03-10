package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type TableBody struct {
	Date                   time.Time
	Open, High, Low, Close float64
	Volume                 int
}

type Data struct {
	Head []string
	Body []TableBody
}

func x(res http.ResponseWriter, req *http.Request) {
	// parse csv
	head, body := parseFile("table.csv")

	// parse template
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	data := Data{
		Head: head,
		Body: body,
	}

	// execute template
	err = tpl.Execute(res, data)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", x)
	http.ListenAndServe(":8090", nil)
}

func parseFile(dataFile string) ([]string, []TableBody) {
	// open file
	file, err := os.Open(dataFile)
	if err != nil {
		log.Fatalln(err)
	}

	// remember to close the file at the end of the program
	defer file.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	// convert head and body data to array of structs
	head := data[0]
	body := make([]TableBody, 0, len(data))

	for i, row := range data {
		if i == 0 {
			continue
		}
		data, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)
		high, _ := strconv.ParseFloat(row[2], 64)
		low, _ := strconv.ParseFloat(row[3], 64)
		close, _ := strconv.ParseFloat(row[4], 64)
		vol, _ := strconv.Atoi(row[5])

		body = append(body, TableBody{
			Date:   data,
			Open:   open,
			High:   high,
			Low:    low,
			Close:  close,
			Volume: vol,
		})
	}
	return head, body
}
