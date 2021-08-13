package structs

import "time"

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
