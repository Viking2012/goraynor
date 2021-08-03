package readr

import (
	"encoding/csv"
	"strings"
	"testing"

	"github.com/Viking2012/goraynor/src/records"
)

var (
	rawLines string = `productId,customerId,purchaseDate,totalPrice
bed_bath_table:8,15df0,2017-02-28,101.14
bed_bath_table:8,f4c13,2017-02-28,104.7
bed_bath_table:9,0dc4b,2017-03-01,101.14
bed_bath_table:8,d98e2,2017-03-02,104.7
bed_bath_table:8,2ed85,2017-03-04,101.14
`
	priceRecords []records.PriceRecord = []records.PriceRecord{
		{Uuid: 0, ProductID: "bed_bath_table:8", CustomerID: "15df0", PurchaseDate: records.QuickParse("2017-02-28"), DocumentNumber: 0, DocumentLineNumber: 1, Price: 101.14},
		{Uuid: 1, ProductID: "bed_bath_table:8", CustomerID: "f4c13", PurchaseDate: records.QuickParse("2017-02-28"), DocumentNumber: 1, DocumentLineNumber: 1, Price: 104.70},
		{Uuid: 2, ProductID: "bed_bath_table:9", CustomerID: "0dc4b", PurchaseDate: records.QuickParse("2017-03-01"), DocumentNumber: 2, DocumentLineNumber: 1, Price: 101.14},
		{Uuid: 3, ProductID: "bed_bath_table:8", CustomerID: "d98e2", PurchaseDate: records.QuickParse("2017-03-02"), DocumentNumber: 3, DocumentLineNumber: 1, Price: 104.70},
		{Uuid: 4, ProductID: "bed_bath_table:8", CustomerID: "2ed85", PurchaseDate: records.QuickParse("2017-03-04"), DocumentNumber: 4, DocumentLineNumber: 1, Price: 101.14},
	}

	badLine [][]string = [][]string{
		{"product_ok", "customer_ok", "2001-02-03", "BadFloat"},
	}
)

func TestParseLines(t *testing.T) {
	r := csv.NewReader(strings.NewReader(rawLines))
	lines, err := r.ReadAll()
	parsed, err := parseLines(lines[1:])

	if err != nil {
		t.Fatal(err)
	}

	for i, got := range parsed {
		want := priceRecords[i]
		if want.Uuid != got.Uuid {
			t.Errorf("For CSV Line %d, wanted               UUID: %d, but got: %d", i, want.Uuid, got.Uuid)
		}
		if want.ProductID != got.ProductID {
			t.Errorf("For CSV Line %d, wanted          ProductID: %s, but got: %s", i, want.ProductID, got.ProductID)
		}
		if want.CustomerID != got.CustomerID {
			t.Errorf("For CSV Line %d, wanted         CustomerID: %s, but got: %s", i, want.CustomerID, got.CustomerID)
		}
		if !want.PurchaseDate.Equal(got.PurchaseDate) {
			t.Errorf("For CSV Line %d, wanted       PurchaseDate: %s, but got: %s", i, want.PurchaseDate.Format("2006-01-02"), got.PurchaseDate.Format("2006-01-02"))
		}
		if want.DocumentNumber != got.DocumentNumber {
			t.Errorf("For CSV Line %d, wanted     DocumentNumber: %d, but got: %d", i, want.DocumentNumber, got.DocumentNumber)
		}
		if want.DocumentLineNumber != got.DocumentLineNumber {
			t.Errorf("For CSV Line %d, wanted DocumentLineNumber: %d, but got: %d", i, want.DocumentLineNumber, got.DocumentLineNumber)
		}
		if want.Price != got.Price {
			t.Errorf("For CSV Line %d, wanted              Price: %.2F, but got: %.2F", i, want.Price, got.Price)
		}
	}
}

func TestParseLinesShouldRaiseFloatError(t *testing.T) {
	_, err := parseLines(badLine)

	if err == nil {
		t.Error("parseLines should have raised a bad float conversion error, but didn't")
	}
}

func TestParseCsvSkipsLines(t *testing.T) {
	parsed, err := ParseCSV("../../test/test_data.csv", 1)

	if err != nil {
		t.Error(err)
	}

	for i, got := range parsed {
		want := priceRecords[i]
		if want.Uuid != got.Uuid {
			t.Errorf("For CSV Line %d, wanted               UUID: %d, but got: %d", i, want.Uuid, got.Uuid)
		}
		if want.ProductID != got.ProductID {
			t.Errorf("For CSV Line %d, wanted          ProductID: %s, but got: %s", i, want.ProductID, got.ProductID)
		}
		if want.CustomerID != got.CustomerID {
			t.Errorf("For CSV Line %d, wanted         CustomerID: %s, but got: %s", i, want.CustomerID, got.CustomerID)
		}
		if !want.PurchaseDate.Equal(got.PurchaseDate) {
			t.Errorf("For CSV Line %d, wanted       PurchaseDate: %s, but got: %s", i, want.PurchaseDate.Format("2006-01-02"), got.PurchaseDate.Format("2006-01-02"))
		}
		if want.DocumentNumber != got.DocumentNumber {
			t.Errorf("For CSV Line %d, wanted     DocumentNumber: %d, but got: %d", i, want.DocumentNumber, got.DocumentNumber)
		}
		if want.DocumentLineNumber != got.DocumentLineNumber {
			t.Errorf("For CSV Line %d, wanted DocumentLineNumber: %d, but got: %d", i, want.DocumentLineNumber, got.DocumentLineNumber)
		}
		if want.Price != got.Price {
			t.Errorf("For CSV Line %d, wanted              Price: %.2F, but got: %.2F", i, want.Price, got.Price)
		}
	}
}
