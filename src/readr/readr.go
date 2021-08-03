package readr

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/Viking2012/goraynor/src/records"
)

// PraseCSV reads a csv at the given filepath and converts it to pricing records we can examine
func ParseCSV(filepath string, headerRows int) ([]records.PriceRecord, error) {
	csvFile, err := os.Open(filepath)

	if err != nil {
		return []records.PriceRecord{}, err
	}

	r := csv.NewReader(csvFile)

	lines, err := r.ReadAll()

	clean, err := parseLines(lines[headerRows:])

	if err != nil {
		return []records.PriceRecord{}, err
	}

	return clean, nil

}

// parseLines allows for easier testing (can take an array of an array of strings from any source location)
func parseLines(lines [][]string) ([]records.PriceRecord, error) {
	numLines := len(lines)

	priceRecords := make([]records.PriceRecord, numLines)

	for i := 0; i < numLines; i++ {
		thisPrice, err := strconv.ParseFloat(lines[i][3], 64)
		if err != nil {
			return []records.PriceRecord{}, err
		}
		priceRecords[i] = records.PriceRecord{
			Uuid:               int64(i),
			ProductID:          lines[i][0],
			CustomerID:         lines[i][1],
			PurchaseDate:       records.QuickParse(lines[i][2]),
			Price:              thisPrice,
			DocumentNumber:     int64(i),
			DocumentLineNumber: int64(1),
		}
	}

	return priceRecords, nil
}
