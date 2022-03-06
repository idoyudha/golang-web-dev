package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type TableRecord struct {
	Date                   time.Time
	Open, High, Low, Close float64
	Volume                 int
}

func main() {
	// open file
	file, err := os.Open("table.csv")
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

	fmt.Println(data)

	// convert records to array of structs

}
