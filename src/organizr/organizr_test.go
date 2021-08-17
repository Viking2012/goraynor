package organizr

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Viking2012/goraynor/src/structs"
	"github.com/Viking2012/goraynor/src/utils"
)

var RawRecords = []structs.PriceRecord{
	{Uuid: 0, ProductID: "bed_bath_table:8", CustomerID: "15df0", PurchaseDate: utils.QuickParse("2017-02-28"), DocumentNumber: 100000000, DocumentLineNumber: 1, Price: 101.14},
	{Uuid: 1, ProductID: "bed_bath_table:8", CustomerID: "f4c13", PurchaseDate: utils.QuickParse("2017-02-28"), DocumentNumber: 100000100, DocumentLineNumber: 1, Price: 104.70},
	{Uuid: 2, ProductID: "bed_bath_table:9", CustomerID: "0dc4b", PurchaseDate: utils.QuickParse("2017-03-01"), DocumentNumber: 100000200, DocumentLineNumber: 1, Price: 101.14},
	{Uuid: 3, ProductID: "bed_bath_table:8", CustomerID: "d98e2", PurchaseDate: utils.QuickParse("2017-03-02"), DocumentNumber: 100000300, DocumentLineNumber: 1, Price: 104.70},
	{Uuid: 4, ProductID: "bed_bath_table:8", CustomerID: "2ed85", PurchaseDate: utils.QuickParse("2017-03-04"), DocumentNumber: 100000400, DocumentLineNumber: 1, Price: 101.14},
	{Uuid: 5, ProductID: "bed_bath_table:9", CustomerID: "6058d", PurchaseDate: utils.QuickParse("2017-03-05"), DocumentNumber: 100000500, DocumentLineNumber: 1, Price: 106.23},
	{Uuid: 6, ProductID: "bed_bath_table:8", CustomerID: "f4c13", PurchaseDate: utils.QuickParse("2017-03-06"), DocumentNumber: 100000600, DocumentLineNumber: 1, Price: 101.14},
	{Uuid: 7, ProductID: "bed_bath_table:8", CustomerID: "d5f2b", PurchaseDate: utils.QuickParse("2017-03-06"), DocumentNumber: 100000700, DocumentLineNumber: 1, Price: 101.14},
	{Uuid: 8, ProductID: "bed_bath_table:8", CustomerID: "0d554", PurchaseDate: utils.QuickParse("2017-03-08"), DocumentNumber: 100000800, DocumentLineNumber: 1, Price: 101.14},
	{Uuid: 9, ProductID: "bed_bath_table:8", CustomerID: "6d52f", PurchaseDate: utils.QuickParse("2017-03-09"), DocumentNumber: 100000900, DocumentLineNumber: 1, Price: 115.02},
	{Uuid: 10, ProductID: "bed_bath_table:9", CustomerID: "679f8", PurchaseDate: utils.QuickParse("2017-03-11"), DocumentNumber: 100001000, DocumentLineNumber: 1, Price: 106.23},
	{Uuid: 11, ProductID: "bed_bath_table:9", CustomerID: "5af63", PurchaseDate: utils.QuickParse("2017-03-13"), DocumentNumber: 100001100, DocumentLineNumber: 1, Price: 101.14},
	{Uuid: 12, ProductID: "bed_bath_table:9", CustomerID: "61e64", PurchaseDate: utils.QuickParse("2017-03-13"), DocumentNumber: 100001200, DocumentLineNumber: 1, Price: 104.70},
	{Uuid: 13, ProductID: "bed_bath_table:9", CustomerID: "5af63", PurchaseDate: utils.QuickParse("2017-03-16"), DocumentNumber: 100001300, DocumentLineNumber: 1, Price: 102.18},
	{Uuid: 14, ProductID: "bed_bath_table:9", CustomerID: "68fe3", PurchaseDate: utils.QuickParse("2017-03-16"), DocumentNumber: 100001400, DocumentLineNumber: 1, Price: 101.14},
	{Uuid: 15, ProductID: "bed_bath_table:8", CustomerID: "f4c13", PurchaseDate: utils.QuickParse("2017-03-20"), DocumentNumber: 100001500, DocumentLineNumber: 1, Price: 104.70},
	{Uuid: 16, ProductID: "bed_bath_table:9", CustomerID: "d98e2", PurchaseDate: utils.QuickParse("2017-03-20"), DocumentNumber: 100001600, DocumentLineNumber: 1, Price: 101.14},
	{Uuid: 17, ProductID: "bed_bath_table:9", CustomerID: "d98e2", PurchaseDate: utils.QuickParse("2017-03-20"), DocumentNumber: 100001600, DocumentLineNumber: 2, Price: 101.18},
	{Uuid: 18, ProductID: "bed_bath_table:8", CustomerID: "4ab4d", PurchaseDate: utils.QuickParse("2017-03-23"), DocumentNumber: 100001700, DocumentLineNumber: 1, Price: 101.14},
	{Uuid: 19, ProductID: "bed_bath_table:9", CustomerID: "20dcb", PurchaseDate: utils.QuickParse("2017-03-27"), DocumentNumber: 100001800, DocumentLineNumber: 1, Price: 106.23},
}

func verifyIntegerOrder(want, got []int64) (bool, string, error) {
	checkLen := len(want)
	var allSame bool = true

	if checkLen != len(got) {
		return false, "", errors.New("Different lengths of original records and sorted records")
	}

	outputString := fmt.Sprintf("\n%9s %-9s\n", "wanted:", "got:")
	for i := 0; i < checkLen; i++ {
		w := want[i]
		g := got[i]

		if w != g {
			allSame = false
			wString := fmt.Sprintf("***%6d", w)
			gString := fmt.Sprintf("%-6d***", g)
			outputString = fmt.Sprintf("%s%s %s\n", outputString, wString, gString)
		} else {
			outputString = fmt.Sprintf("%s%9d %-9d\n", outputString, w, g)
		}
	}

	return allSame, outputString, nil
}

func verifyFloatOrder(want, got []float64) (bool, string, error) {
	checkLen := len(want)
	var allSame bool = true

	if checkLen != len(got) {
		return false, "", errors.New("Different lengths of original records and sorted records")
	}

	outputString := fmt.Sprintf("\n%9s %-9s\n", "wanted:", "got:")
	for i := 0; i < checkLen; i++ {
		w := want[i]
		g := got[i]

		if w != g {
			allSame = false
			wString := fmt.Sprintf("***%6.2f", w)
			gString := fmt.Sprintf("%-6.2f***", g)
			outputString = fmt.Sprintf("%s%s %s\n", outputString, wString, gString)
		} else {
			outputString = fmt.Sprintf("%s%9.2f %-9.2f\n", outputString, w, g)
		}
	}

	return allSame, outputString, nil
}

func verifyStringOrder(want, got []string) (bool, string, error) {
	checkLen := len(want)
	var allSame bool = true

	if checkLen != len(got) {
		return false, "", errors.New("Different lengths of original records and sorted records")
	}

	outputString := fmt.Sprintf("\n%19s %-19s\n", "wanted:", "got:")
	for i := 0; i < checkLen; i++ {
		w := want[i]
		g := got[i]

		if w != g {
			allSame = false
			wString := fmt.Sprintf("***%16s", w)
			gString := fmt.Sprintf("%-16s***", g)
			outputString = fmt.Sprintf("%s%s %s\n", outputString, wString, gString)
		} else {
			outputString = fmt.Sprintf("%s%19s %-19s\n", outputString, w, g)
		}
	}

	return allSame, outputString, nil
}

func TestOrderByUuid(t *testing.T) {
	want := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	s := make([]structs.PriceRecord, len(RawRecords))
	copy(s, RawRecords)

	OrderedBy(ByUuid).Sort(s)

	r := make([]int64, len(s))
	for i := range r {
		r[i] = s[i].Uuid
	}

	same, outputString, err := verifyIntegerOrder(want, r)

	if err != nil {
		t.Fatal(err)
	}

	if !same {
		t.Errorf(outputString)
	}
}

func TestOrderByProductID(t *testing.T) {
	want := []string{"bed_bath_table:8", "bed_bath_table:8", "bed_bath_table:8", "bed_bath_table:8", "bed_bath_table:8", "bed_bath_table:8", "bed_bath_table:8", "bed_bath_table:8", "bed_bath_table:8", "bed_bath_table:8", "bed_bath_table:9", "bed_bath_table:9", "bed_bath_table:9", "bed_bath_table:9", "bed_bath_table:9", "bed_bath_table:9", "bed_bath_table:9", "bed_bath_table:9", "bed_bath_table:9", "bed_bath_table:9"}
	s := make([]structs.PriceRecord, len(RawRecords))
	copy(s, RawRecords)

	OrderedBy(ByProduct).Sort(s)

	r := make([]string, len(s))
	for i := range r {
		r[i] = s[i].ProductID
	}

	same, outputString, err := verifyStringOrder(want, r)

	if err != nil {
		t.Fatal(err)
	}

	if !same {
		t.Errorf(outputString)
	}
}

func TestOrderByCustomerID(t *testing.T) {
	want := []string{"0d554", "0dc4b", "15df0", "20dcb", "2ed85", "4ab4d", "5af63", "5af63", "6058d", "61e64", "679f8", "68fe3", "6d52f", "d5f2b", "d98e2", "d98e2", "d98e2", "f4c13", "f4c13", "f4c13"}
	s := make([]structs.PriceRecord, len(RawRecords))
	copy(s, RawRecords)

	OrderedBy(ByCustomer).Sort(s)

	r := make([]string, len(s))
	for i := range r {
		r[i] = s[i].CustomerID
	}

	same, outputString, err := verifyStringOrder(want, r)

	if err != nil {
		t.Fatal(err)
	}

	if !same {
		t.Errorf(outputString)
	}
}

func TestOrderByPurchaseDate(t *testing.T) {
	want := []string{"2017-02-28", "2017-02-28", "2017-03-01", "2017-03-02", "2017-03-04", "2017-03-05", "2017-03-06", "2017-03-06", "2017-03-08", "2017-03-09", "2017-03-11", "2017-03-13", "2017-03-13", "2017-03-16", "2017-03-16", "2017-03-20", "2017-03-20", "2017-03-20", "2017-03-23", "2017-03-27"}
	s := make([]structs.PriceRecord, len(RawRecords))
	copy(s, RawRecords)

	OrderedBy(ByDate).Sort(s)

	r := make([]string, len(s))
	for i := range r {
		r[i] = s[i].PurchaseDate.Format("2006-01-02")
	}

	same, outputString, err := verifyStringOrder(want, r)

	if err != nil {
		t.Fatal(err)
	}

	if !same {
		t.Errorf(outputString)
	}
}

func TestOrderByDocumentNumber(t *testing.T) {
	want := []int64{100000000, 100000100, 100000200, 100000300, 100000400, 100000500, 100000600, 100000700, 100000800, 100000900, 100001000, 100001100, 100001200, 100001300, 100001400, 100001500, 100001600, 100001600, 100001700, 100001800}
	s := make([]structs.PriceRecord, len(RawRecords))
	copy(s, RawRecords)

	OrderedBy(ByDocumentNumber).Sort(s)

	r := make([]int64, len(s))
	for i := range r {
		r[i] = s[i].DocumentNumber
	}

	same, outputString, err := verifyIntegerOrder(want, r)

	if err != nil {
		t.Fatal(err)
	}

	if !same {
		t.Errorf(outputString)
	}
}

func TestOrderByDocumentLineNumber(t *testing.T) {
	want := []int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2}
	s := make([]structs.PriceRecord, len(RawRecords))
	copy(s, RawRecords)

	OrderedBy(ByDocumentLineNumber).Sort(s)

	r := make([]int64, len(s))
	for i := range r {
		r[i] = s[i].DocumentLineNumber
	}

	same, outputString, err := verifyIntegerOrder(want, r)

	if err != nil {
		t.Fatal(err)
	}

	if !same {
		t.Errorf(outputString)
	}
}

func TestOrderByPrice(t *testing.T) {
	want := []float64{101.14, 101.14, 101.14, 101.14, 101.14, 101.14, 101.14, 101.14, 101.14, 101.14, 101.18, 102.18, 104.7, 104.7, 104.7, 104.7, 106.23, 106.23, 106.23, 115.02}
	s := make([]structs.PriceRecord, len(RawRecords))
	copy(s, RawRecords)

	OrderedBy(ByPrice).Sort(s)

	r := make([]float64, len(s))
	for i := range r {
		r[i] = s[i].Price
	}

	same, outputString, err := verifyFloatOrder(want, r)

	if err != nil {
		t.Fatal(err)
	}

	if !same {
		t.Errorf(outputString)
	}
}

func TestOrderByProductCustomerDocumentDatePrice(t *testing.T) {
	want := []int64{8, 0, 4, 18, 9, 7, 3, 1, 6, 15, 2, 19, 11, 13, 5, 12, 10, 14, 16, 17}
	s := make([]structs.PriceRecord, len(RawRecords))
	copy(s, RawRecords)

	OrderedBy(ByProduct, ByCustomer, ByDate, ByDocumentNumber, ByDocumentLineNumber, ByPrice).Sort(s)

	r := make([]int64, len(s))
	for i := range r {
		r[i] = s[i].Uuid
	}

	same, outputString, err := verifyIntegerOrder(want, r)

	if err != nil {
		t.Fatal(err)
	}

	if !same {
		t.Errorf(outputString)
	}
}

func TestOrderByDocumentAndLineNumber(t *testing.T) {
	want := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	s := make([]structs.PriceRecord, len(RawRecords))
	copy(s, RawRecords)

	OrderedBy(ByDocumentNumber, ByDocumentLineNumber).Sort(s)

	r := make([]int64, len(s))
	for i := range r {
		r[i] = s[i].Uuid
	}

	same, outputString, err := verifyIntegerOrder(want, r)

	if err != nil {
		t.Fatal(err)
	}

	if !same {
		t.Errorf(outputString)
	}
}
