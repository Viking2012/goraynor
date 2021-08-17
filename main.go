package main

import (
	"github.com/Viking2012/goraynor/src/organizr"
	"github.com/Viking2012/goraynor/src/readr"
)

func main() {
	raw, err := readr.ParseCSV("./test/test_data.csv", 1, &readr.DefaultFieldMap)
	if err != nil {
		panic(err)
	}

	organizr.OrderedBy(organizr.ByProduct, organizr.ByCustomer, organizr.ByDate, organizr.ByDocumentNumber, organizr.ByDocumentLineNumber).Sort(raw)
}
