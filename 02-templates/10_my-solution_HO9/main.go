package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

type Record struct {
	Date  time.Time
	Close float64
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	records := prs("table.csv")

	tpl, err := template.ParseFiles("tpl.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(res, records)
	if err != nil {
		log.Fatalln(err)
	}
}

func prs(filePath string) []Record {

	src, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()

	rdr := csv.NewReader(src)
	rows, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	records := make([]Record, 0, len(rows))

	for i, row := range rows {

		if i == 0 {
			continue
		}
		date, _ := time.Parse("2006-01-02", row[0])
		close, _ := strconv.ParseFloat(row[6], 64)

		records = append(records, Record{
			Date:  date,
			Close: close,
		})

	}

	return records

}
