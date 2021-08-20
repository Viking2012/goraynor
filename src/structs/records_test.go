package structs

import (
	"testing"

	"github.com/Viking2012/goraynor/src/quantilr"
)

var testDeciles quantilr.Deciles = quantilr.Deciles{
	Pairs: []quantilr.DecilePair{
		{Decile: 1, Weight: 0.0},
		{Decile: 2, Weight: 10.0},
		{Decile: 3, Weight: 20.0},
		{Decile: 4, Weight: 30.0},
		{Decile: 5, Weight: 40.0},
		{Decile: 6, Weight: 50.0},
		{Decile: 7, Weight: 60.0},
		{Decile: 8, Weight: 70.0},
		{Decile: 9, Weight: 80.0},
		{Decile: 10, Weight: 90.0},
	},
}

func TestPriceRecordHasDecileSet(t *testing.T) {
	type testCase struct {
		PR   *PriceRecord
		Want int8
	}
	var testCases []testCase = []testCase{
		{PR: &PriceRecord{PriceReturn: -1}, Want: 1},
		{PR: &PriceRecord{PriceReturn: 9}, Want: 2},
		{PR: &PriceRecord{PriceReturn: 11}, Want: 3},
		{PR: &PriceRecord{PriceReturn: 21}, Want: 4},
		{PR: &PriceRecord{PriceReturn: 39}, Want: 5},
		{PR: &PriceRecord{PriceReturn: 41}, Want: 6},
		{PR: &PriceRecord{PriceReturn: 59.999}, Want: 7},
		{PR: &PriceRecord{PriceReturn: 69.999}, Want: 8},
		{PR: &PriceRecord{PriceReturn: 80}, Want: 9},
		{PR: &PriceRecord{PriceReturn: 89}, Want: 10},
	}

	for i := 0; i < len(testCases); i++ {
		pr, want := testCases[i].PR, testCases[i].Want
		err := pr.SetDecile(&testDeciles)
		if err != nil {
			t.Error(err)
		}
		if pr.DecileOfPrice != want {
			t.Errorf("For price record %v, wanted decile %d, but got %d", pr, want, pr.DecileOfPrice)
		}
	}
}
