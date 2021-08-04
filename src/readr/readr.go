package readr

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"

	"github.com/Viking2012/goraynor/src/records"
)

// FieldIndexMap provides a structure which enables readr to find data fields in the csv provided.
// Seven fields are required to build records, but not all might be present in the data sent to this command.
// Use a negative value for readr to "build" the field using the row index.
type FieldIndexMap struct {
	UUID               int
	ProductID          int
	CustomerID         int
	PurchaseDate       int
	DocumentNumber     int
	DocumentLineNumber int
	Price              int
}

var DefaultFieldMap FieldIndexMap = FieldIndexMap{
	UUID:               0,
	ProductID:          1,
	CustomerID:         2,
	PurchaseDate:       3,
	DocumentNumber:     4,
	DocumentLineNumber: 5,
	Price:              6,
}

// ParseCSV reads a csv at the given filepath and converts it to pricing records we can examine
func ParseCSV(filepath string, headerRows int, fieldMap *FieldIndexMap) ([]records.PriceRecord, error) {
	// if fieldMap was a nil pointer, then use the default map
	if fieldMap == nil {
		fieldMap = &DefaultFieldMap
	}

	// Price MUST be provided - without it, what are we running this for?
	if fieldMap.Price < 0 {
		return []records.PriceRecord{}, errors.New("Index for price must be provided")
	}

	csvFile, err := os.Open(filepath)

	if err != nil {
		return []records.PriceRecord{}, err
	}

	r := csv.NewReader(csvFile)

	lines, err := r.ReadAll()

	var clean []records.PriceRecord

	clean, err = parseLines(lines[headerRows:], fieldMap)

	if err != nil {
		return []records.PriceRecord{}, err
	}

	return clean, nil

}

func extractIntField(record []string, givenIndex, i int) (int, error) {
	if givenIndex < 0 {
		return i, nil
	}

	r, err := strconv.Atoi(record[givenIndex])
	return r, err
}

func extractFloatField(record []string, givenIndex, i int) (float64, error) {
	if givenIndex < 0 {
		return float64(i), nil
	}

	r, err := strconv.ParseFloat(record[givenIndex], 64)
	return r, err
}

// parseLines allows for easier testing (can take an array of an array of strings from any source location)
func parseLines(lines [][]string, fieldMap *FieldIndexMap) ([]records.PriceRecord, error) {
	numLines := len(lines)
	priceRecords := make([]records.PriceRecord, numLines)

	for i := 0; i < numLines; i++ {
		thisUuid, err := extractIntField(lines[i], fieldMap.UUID, i)
		if err != nil {
			return []records.PriceRecord{}, err
		}

		thisPrice, err := extractFloatField(lines[i], fieldMap.Price, i)
		if err != nil {
			return []records.PriceRecord{}, err
		}

		thisDoc, err := extractIntField(lines[i], fieldMap.DocumentNumber, i)
		if err != nil {
			return []records.PriceRecord{}, err
		}

		thisDocNum, err := extractIntField(lines[i], fieldMap.DocumentLineNumber, i)
		if err != nil {
			return []records.PriceRecord{}, err
		}

		priceRecords[i] = records.PriceRecord{
			Uuid:               int64(thisUuid),
			ProductID:          lines[i][fieldMap.ProductID],
			CustomerID:         lines[i][fieldMap.CustomerID],
			PurchaseDate:       records.QuickParse(lines[i][fieldMap.PurchaseDate]),
			Price:              thisPrice,
			DocumentNumber:     int64(thisDoc),
			DocumentLineNumber: int64(thisDocNum),
		}
	}

	return priceRecords, nil
}
