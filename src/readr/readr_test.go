package readr

import (
	"encoding/csv"
	"strings"
	"testing"

	"github.com/Viking2012/goraynor/src/structs"
	"github.com/Viking2012/goraynor/src/utils"
)

// Reused testing variables
var (
	rawCsvPath string = "../../test/test_data.csv"
	rawLines   string = `productId,customerId,purchaseDate,totalPrice
bed_bath_table:8,15df0,2017-02-28,101.14
bed_bath_table:8,f4c13,2017-02-28,104.7
bed_bath_table:9,0dc4b,2017-03-01,101.14
bed_bath_table:8,d98e2,2017-03-02,104.7
bed_bath_table:8,2ed85,2017-03-04,101.14
`
	rawFieldMap *FieldIndexMap = &FieldIndexMap{
		UUID:               -1,
		ProductID:          0,
		CustomerID:         1,
		PurchaseDate:       2,
		DocumentNumber:     -1,
		DocumentLineNumber: -1,
		Price:              3,
	}

	priceRecords []structs.PriceRecord = []structs.PriceRecord{
		{Uuid: 0, ProductID: "bed_bath_table:8", CustomerID: "15df0", PurchaseDate: utils.QuickParse("2017-02-28"), DocumentNumber: 0, DocumentLineNumber: 0, Price: 101.14},
		{Uuid: 1, ProductID: "bed_bath_table:8", CustomerID: "f4c13", PurchaseDate: utils.QuickParse("2017-02-28"), DocumentNumber: 1, DocumentLineNumber: 1, Price: 104.70},
		{Uuid: 2, ProductID: "bed_bath_table:9", CustomerID: "0dc4b", PurchaseDate: utils.QuickParse("2017-03-01"), DocumentNumber: 2, DocumentLineNumber: 2, Price: 101.14},
		{Uuid: 3, ProductID: "bed_bath_table:8", CustomerID: "d98e2", PurchaseDate: utils.QuickParse("2017-03-02"), DocumentNumber: 3, DocumentLineNumber: 3, Price: 104.70},
		{Uuid: 4, ProductID: "bed_bath_table:8", CustomerID: "2ed85", PurchaseDate: utils.QuickParse("2017-03-04"), DocumentNumber: 4, DocumentLineNumber: 4, Price: 101.14},
	}

	badLine [][]string = [][]string{
		{"product_ok", "customer_ok", "2001-02-03", "BadFloat"},
	}
)

// ParseCsv Testing

func TestParseCsvSkipsLines(t *testing.T) {
	parsed, err := ParseCSV(rawCsvPath, 1, &DefaultFieldMap)

	if err != nil {
		t.Errorf("ParseCsv should have skipped a header row which would have caused an error. Instead it returned: \n%s", err)
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

func TestParseCsvUsesDefaultFieldMap(t *testing.T) {
	parsed, err := ParseCSV(rawCsvPath, 1, nil)

	if err != nil {
		t.Errorf("ParseCsv should have skipped a header row which would have caused an error. Instead it returned: \n%s", err)
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

func TestParseCsvCatchesFieldMapWithoutPriceIndex(t *testing.T) {
	_, err := ParseCSV(rawCsvPath, 1, &FieldIndexMap{Price: -1})
	if err == nil {
		t.Errorf("ParseCsv should have errored on a FieldIndexMap without a price index, but didn't")
	}
}

func TestParseCsvErrorsOnNonExistentFile(t *testing.T) {
	_, err := ParseCSV("DoesNotExist", 1, nil)
	if err == nil {
		t.Errorf("ParseCsv should have errored when receiving a non-existent filepath, but didn't")
	}
}

func TestParseCsvErrorsOnBadRecord(t *testing.T) {
	// here, we use 0 for skipLines since we want to "incorrectly" convert the header into a PriceRecord
	_, err := ParseCSV(rawCsvPath, 0, nil)
	if err == nil {
		t.Errorf("ParseCsv should have errored when trying to convert a bad record, but didn't")
	}
}

// parseLines Testing
func TestParseLines(t *testing.T) {
	r := csv.NewReader(strings.NewReader(rawLines))
	lines, err := r.ReadAll()
	parsed, err := parseLines(lines[1:], rawFieldMap)

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
	_, err := parseLines(badLine, rawFieldMap)

	if err == nil {
		t.Error("parseLines should have raised a bad float conversion error, but didn't")
	}
}

func TestParseLinesErrorsOnBadUuid(t *testing.T) {
	var badRecord [][]string = [][]string{{"BadUuid", "product_ok", "customer_ok", "2001-02-03", "1", "1", "100.00"}}
	_, err := parseLines(badRecord, &DefaultFieldMap)
	if err == nil {
		t.Error("parseLines should have raised a bad UUID conversion error, but didn't")
	}
}
func TestParseLinesErrorsOnBadPrice(t *testing.T) {
	var badRecord [][]string = [][]string{{"1", "product_ok", "customer_ok", "2001-02-03", "1", "1", "BadPrice"}}
	_, err := parseLines(badRecord, &DefaultFieldMap)
	if err == nil {
		t.Error("parseLines should have raised a bad UUID conversion error, but didn't")
	}
}
func TestParseLinesErrorsOnBadDoc(t *testing.T) {
	var badRecord [][]string = [][]string{{"1", "product_ok", "customer_ok", "2001-02-03", "BadDoc", "1", "100.00"}}
	_, err := parseLines(badRecord, &DefaultFieldMap)
	if err == nil {
		t.Error("parseLines should have raised a bad UUID conversion error, but didn't")
	}
}
func TestParseLinesErrorsOnBadDocLine(t *testing.T) {
	var badRecord [][]string = [][]string{{"1", "product_ok", "customer_ok", "2001-02-03", "1", "BadLine", "100.00"}}
	_, err := parseLines(badRecord, &DefaultFieldMap)
	if err == nil {
		t.Error("parseLines should have raised a bad UUID conversion error, but didn't")
	}
}

// Extraction & Conversion Testing
func TestExtractIntFieldHandlesNegativeIndex(t *testing.T) {
	want := int(10)
	got, err := extractIntField([]string{"99"}, -1, 10)

	if err != nil {
		t.Error(err)
	}
	if want != got {
		t.Errorf("extractIntField should have returned %d, but got %d instead", want, got)
	}
}

func TestExtractFloatFieldHandlesNegativeIndex(t *testing.T) {
	want := float64(10)
	got, err := extractFloatField([]string{"99.9"}, -1, 10)

	if err != nil {
		t.Error(err)
	}
	if want != got {
		t.Errorf("extractFloatField should have returned %.2f, but got %.2f instead", want, got)
	}
}
