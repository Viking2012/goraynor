package structs

import (
	"github.com/Viking2012/goraynor/src/quantilr"
	"gonum.org/v1/gonum/stat/distuv"
)

type ProductContainer struct {
	ProductID       string
	DecileModel     distuv.Categorical
	CustomerRecords []CustomerContainer
}

type CustomerContainer struct {
	CustomerID   string
	PriceRecords []PriceRecord
	LifeSpan     int16
	DecileCounts []quantilr.DecilePair
}
