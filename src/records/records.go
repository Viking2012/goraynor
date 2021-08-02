package records

import (
	"sort"
	"time"
)

const layout = "2006-01-02"

// A PriceRecord contains all of the information on individual purchases and their prices to identify outliers
type PriceRecord struct {
	Uuid               int64
	ProductID          string
	CustomerID         string
	PurchaseDate       time.Time
	DocumentNumber     int64
	DocumentLineNumber int64
	Price              float64
}

// implementation basics from https://pkg.go.dev/sort#example-package-SortKeys
// lessFunc is the type of a "less" function that defines the ordering of its PurchaseRecord arguments.
type lessFunc func(p1, p2 *PriceRecord) bool

// multiSorter implements the Sort interface, sorting the changes within.
type multiSorter struct {
	records []PriceRecord
	less    []lessFunc
}

// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (ms *multiSorter) Sort(records []PriceRecord) {
	ms.records = records
	sort.Sort(ms)
}

// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

// Len is part of sort.Interface.
func (ms *multiSorter) Len() int {
	return len(ms.records)
}

// Swap is part of sort.Interface
func (ms *multiSorter) Swap(i, j int) {
	ms.records[i], ms.records[j] = ms.records[j], ms.records[i]
}

// Less is part of sort.Interface. It is implemented by looping along the
// less functions until it finds a comparison that discriminates between
// the two items (one is less than the other). Note that it can call the
// less functions twice per call. We could change the functions to return
// -1, 0, 1 and reduce the number of calls for greater efficiency: an
// exercise for the reader.
func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.records[i], &ms.records[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return ms.less[k](p, q)
}

func byUuid(p1, p2 *PriceRecord) bool {
	return p1.Uuid < p2.Uuid
}

func byProduct(p1, p2 *PriceRecord) bool {
	return p1.ProductID < p2.ProductID
}

func byCustomer(p1, p2 *PriceRecord) bool {
	return p1.CustomerID < p2.CustomerID
}

func byDate(p1, p2 *PriceRecord) bool {
	return p1.PurchaseDate.Before(p2.PurchaseDate)
}

func byDocumentNumber(p1, p2 *PriceRecord) bool {
	return p1.DocumentNumber < p2.DocumentNumber
}

func byDocumentLineNumber(p1, p2 *PriceRecord) bool {
	return p1.DocumentLineNumber < p2.DocumentLineNumber
}

func byPrice(p1, p2 *PriceRecord) bool {
	return p1.Price < p2.Price
}

var ByUuid lessFunc = byUuid
var ByProduct lessFunc = byProduct
var ByCustomer lessFunc = byCustomer
var ByDate lessFunc = byDate
var ByDocumentNumber lessFunc = byDocumentNumber
var ByDocumentLineNumber lessFunc = byDocumentLineNumber
var ByPrice lessFunc = byPrice

// QuickParse simply reutrns a result of parsing a time string in YYYY-MM-DD format
func QuickParse(dateString string) time.Time {
	t, err := time.Parse(layout, dateString)
	if err != nil {
		panic(err)
	}
	return t
}
