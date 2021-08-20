package structs

import (
	"time"

	"github.com/Viking2012/goraynor/src/quantilr"
)

// OLD VERSION OF A PRICE RECORD!!!
// A PriceRecord contains all of the information on individual purchases and their prices to identify outliers
// type PriceRecord struct {
// 	Uuid               int64
// 	ProductID          string
// 	CustomerID         string
// 	PurchaseDate       time.Time
// 	DocumentNumber     int64
// 	DocumentLineNumber int64
// 	Price              float64
// }

// TODO(ajo): remove this struct and below methods into somehwere else
// perhaps replacing "PriceRecord" in the structs folder?
type PriceRecord struct {
	TickerDate    time.Time
	PriceReturn   float64
	DecileOfPrice int8
}

func (pr *PriceRecord) SetDecile(d *quantilr.Deciles) error {
	thisDecile, err := d.LookupValue(pr.PriceReturn)
	if err != nil {
		return err
	}
	(*pr).DecileOfPrice = thisDecile
	return nil
}

// TODO(ajo): remove this struct and below methods into somehwere else
// perhaps replacing "PriceRecord" in the structs folder?
type PriceRecords []PriceRecord

func (a PriceRecords) Len() int               { return len(a) }
func (a PriceRecords) Swap(i, j int)          { a[i], a[j] = a[j], a[i] }
func (a PriceRecords) Less(i, j int) bool     { return a[i].TickerDate.Before(a[j].TickerDate) }
func (a PriceRecords) Get(i int) *PriceRecord { return &a[i] }

func (prs *PriceRecords) SetDeciles(d *quantilr.Deciles) error {
	for i := 0; i < len(*prs); i++ {
		pr := prs.Get(i)
		err := pr.SetDecile(d)
		if err != nil {
			return err
		}
	}
	return nil
}

// TODO(ajo): remove this struct and below methods into somehwere else
// perhaps replacing "PriceRecord" in the structs folder?
type AllPerformers map[string]*PriceRecords

func (ap *AllPerformers) SetDeciles(d *quantilr.Deciles) error {
	for _, pr := range *ap {
		err := pr.SetDeciles(d)
		if err != nil {
			return err
		}
	}
	return nil
}
