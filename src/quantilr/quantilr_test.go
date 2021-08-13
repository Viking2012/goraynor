package quantilr

import (
	"sort"
	"testing"

	"github.com/Viking2012/goraynor/src/countr"
)

var (
	v        []float64      = []float64{1, 1, 1, 1, 1, 1, 1, 1.5, 1.5, 1.5, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	vNums    []float64      = []float64{1, 1.5, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	vWeights []float64      = []float64{7, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	probs    []float64      = []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0}
	dNums    []float64      = []float64{1, 1, 1, 1.5, 2, 4, 6, 8, 10, 12}
	c        countr.Counter = countr.Counter{
		countr.CounterPair{Value: 1, Count: 7},
		countr.CounterPair{Value: 1.5, Count: 3},
		countr.CounterPair{Value: 2, Count: 1},
		countr.CounterPair{Value: 3, Count: 1},
		countr.CounterPair{Value: 4, Count: 1},
		countr.CounterPair{Value: 5, Count: 1},
		countr.CounterPair{Value: 6, Count: 1},
		countr.CounterPair{Value: 7, Count: 1},
		countr.CounterPair{Value: 8, Count: 1},
		countr.CounterPair{Value: 9, Count: 1},
		countr.CounterPair{Value: 10, Count: 1},
		countr.CounterPair{Value: 11, Count: 1},
		countr.CounterPair{Value: 12, Count: 1},
	}
)

func TestDecileLen(t *testing.T) {
	var d = Deciles{
		Pairs: []DecilePair{
			{Decile: 1, Weight: 20.0},
			{Decile: 2, Weight: 10.0},
		},
	}

	var want int = 2
	got := d.Len()

	if got != want {
		t.Errorf("Len should have returned %3d, got %3d instead", want, got)
	}
}
func TestDecileLess(t *testing.T) {
	var d = Deciles{
		Pairs: []DecilePair{
			{Decile: 1, Weight: 20.0},
			{Decile: 2, Weight: 10.0},
		},
	}

	var want bool = true
	var got bool = d.Less(0, 1)

	if got != want {
		t.Errorf("Less for indices 0 and 1 should have returned %t, got %t instead", want, got)
	}
}
func TestDecileSwap(t *testing.T) {
	var d = Deciles{
		Pairs: []DecilePair{
			{Decile: 1, Weight: 20.0},
			{Decile: 2, Weight: 10.0},
		},
	}

	var wantIndices []int8 = []int8{2, 1}
	var gotIndices []int8 = make([]int8, d.Len())

	d.Swap(0, 1)
	for i := 0; i < d.Len(); i++ {
		gotIndices[i] = d.Pairs[i].Decile
	}

	for i := 0; i < d.Len(); i++ {
		w := wantIndices[i]
		g := gotIndices[i]

		if g != w {
			t.Errorf("After Swap, index %d should been %3d, got %3d instead", i, w, g)
		}
	}
}

func TestNewDeciles(t *testing.T) {
	var want = make([]DecilePair, len(dNums))
	for i := 0; i < len(dNums); i++ {
		thisProb := int8(probs[i] * 10)
		thisWeight := dNums[i]
		want[i] = DecilePair{Decile: thisProb, Weight: thisWeight}

	}

	got, err := NewDeciles(c, false)
	if err != nil {
		t.Errorf("NewDecile errored unexpectedly with %s", err)
	}

	sort.Sort(got)

	for i, g := range got.Pairs {
		w := want[i]
		if g.Decile != w.Decile {
			t.Errorf("For index %d, wanted decile %d, but got %d", i, w.Decile, g.Decile)
		}
		if g.Weight != w.Weight {
			t.Errorf("For index %d (decile %d), wanted weight %3.1f, but got %3.1f", i, w.Decile, w.Weight, g.Weight)
		}
	}
}

func TestNewDecilesDeduplicatesPairsWithSameValue(t *testing.T) {
	var want []DecilePair = []DecilePair{
		{Decile: 3, Weight: 1.0},
		{Decile: 4, Weight: 1.5},
		{Decile: 5, Weight: 2.0},
		{Decile: 6, Weight: 4.0},
		{Decile: 7, Weight: 6.0},
		{Decile: 8, Weight: 8.0},
		{Decile: 9, Weight: 10.0},
		{Decile: 10, Weight: 12.0},
	}

	got, err := NewDeciles(c, true)
	if err != nil {
		t.Errorf("NewDecile errored unexpectedly with %s", err)
	}

	sort.Sort(got)

	for i, g := range got.Pairs {
		w := want[i]
		if g.Decile != w.Decile {
			t.Errorf("For index %d, wanted decile %d, but got %d", i, w.Decile, g.Decile)
		}
		if g.Weight != w.Weight {
			t.Errorf("For index %d (decile %d), wanted weight %3.1f, but got %3.1f", i, w.Decile, w.Weight, g.Weight)
		}
	}
}

func TestLookupValueReturnsCorrectDecile(t *testing.T) {
	var lookupHere = Deciles{
		Pairs: []DecilePair{
			{Decile: 3, Weight: 1.0},
			{Decile: 4, Weight: 1.5},
			{Decile: 5, Weight: 2.0},
			{Decile: 6, Weight: 4.0},
			{Decile: 7, Weight: 6.0},
			{Decile: 8, Weight: 8.0},
			{Decile: 9, Weight: 10.0},
			{Decile: 10, Weight: 12.0},
		},
		isSorted:  true,
		isDeduped: true,
	}

	var want int8 = 7
	var lookupValue float64 = 5.0
	got, err := lookupHere.LookupValue(lookupValue)
	if err != nil {
		t.Errorf("Wanted decile %d, but got an error instead: %s", want, err)
	}

	if got != want {
		t.Errorf("When lookup up value %3.1f, wanted decile %d, but got %d", lookupValue, want, got)
	}
}

func TestLookupValueErrorsOnBadValue(t *testing.T) {
	var lookupHere = Deciles{
		Pairs: []DecilePair{
			{Decile: 3, Weight: 1.0},
			{Decile: 4, Weight: 1.5},
			{Decile: 5, Weight: 2.0},
			{Decile: 6, Weight: 4.0},
			{Decile: 7, Weight: 6.0},
			{Decile: 8, Weight: 8.0},
			{Decile: 9, Weight: 10.0},
			{Decile: 10, Weight: 12.0},
		},
		isSorted:  true,
		isDeduped: true,
	}

	var lookupValue float64 = 100.0
	_, err := lookupHere.LookupValue(lookupValue)
	if err == nil {
		t.Error("When looking up a value out of the decile range, should have returned an error but didn't")
	}
	if err != ValueNotFound {
		t.Errorf("When looking up a value out of the decile range, should have returned a ValueNotFound error, but got %s", err)
	}
}

func TestLookupValueDeduplicatesAutomatically(t *testing.T) {
	var lookupHere = Deciles{
		Pairs: []DecilePair{
			{Decile: 3, Weight: 1.0},
			{Decile: 4, Weight: 1.5},
			{Decile: 5, Weight: 2.0},
			{Decile: 6, Weight: 4.0},
			{Decile: 7, Weight: 6.0},
			{Decile: 8, Weight: 8.0},
			{Decile: 9, Weight: 10.0},
			{Decile: 10, Weight: 12.0},
		},
		isSorted: true,
		// although the above decilesPairs are already deduped,
		// the struct relies upon this flag for checking
		isDeduped: false,
	}

	var lookupValue float64 = 5.0
	_, err := lookupHere.LookupValue(lookupValue)
	if err == nil {
		t.Error("When sending a non-deduplicated set of DecilePairs, should have warned this was applied but didn't")
	}
	if err != WarnDecilesNotDeduplicated {
		t.Errorf("When sending a non-deduplicated set of DecilePairs, should have returned a WarnDecilesNotDeduplicated error, but got %s", err)
	}
}

func TestLookupValueSortsAutomatically(t *testing.T) {
	var lookupHere = Deciles{
		Pairs: []DecilePair{
			{Decile: 3, Weight: 1.0},
			{Decile: 4, Weight: 1.5},
			{Decile: 5, Weight: 2.0},
			{Decile: 6, Weight: 4.0},
			{Decile: 7, Weight: 6.0},
			{Decile: 8, Weight: 8.0},
			{Decile: 9, Weight: 10.0},
			{Decile: 10, Weight: 12.0},
		},
		// although the above decilesPairs are already sorted,
		// the struct relies upon this flag for checking
		isSorted:  false,
		isDeduped: true,
	}

	var lookupValue float64 = 5.0
	_, err := lookupHere.LookupValue(lookupValue)
	if err == nil {
		t.Error("When sending a non-deduplicated set of DecilePairs, should have warned this was applied but didn't")
	}
	if err != WarnDecilesNotSorted {
		t.Errorf("When sending a non-deduplicated set of DecilePairs, should have returned a WarnDecilesNotSorted error, but got %s", err)
	}
}
