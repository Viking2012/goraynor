package countr

import (
	"sort"
	"testing"
)

var (
	// quantiles []float64 = []float64{0.0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0}
	v        []float64 = []float64{1, 1, 1, 1, 1, 1, 1, 1.5, 1.5, 1.5, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	vNums    []float64 = []float64{1, 1.5, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	vWeights []float64 = []float64{7, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
)

func TestCount(t *testing.T) {
	want := make(map[float64]float64)
	for i := 0; i < len(vNums); i++ {
		u := vNums[i]
		w := vWeights[i]
		want[u] = w
	}

	q := Count(v)
	sort.Sort(q)

	for _, pair := range q {
		u, w := pair.Value, pair.Count
		if want[u] != w {
			t.Errorf("Count returned a CounterPair where, for key %3.1f, we wanted %3.1f but got %3.1f", u, want[u], w)
		}
	}
}

func TestNewCounter(t *testing.T) {
	var want = make(Counter, len(vNums))
	var toGet = make(map[float64]float64, len(vNums))
	for i := 0; i < len(vNums); i++ {
		want[i] = CounterPair{Value: vNums[i], Count: vWeights[i]}
		toGet[vNums[i]] = vWeights[i]
	}

	got := NewCounter(toGet)

	for i, g := range got {
		w := want[i]
		if g.Value != w.Value {
			t.Errorf("For value: %3.1f, wanted count %3.1f but got %3.1f", w.Value, w.Count, g.Count)
		}
	}
}
func TestGetValues(t *testing.T) {
	var c Counter = Counter{
		CounterPair{Value: 1.0, Count: 10.0},
		CounterPair{Value: 2.0, Count: 20.0},
		CounterPair{Value: 3.0, Count: 30.0},
	}

	var want []float64 = []float64{1.0, 2.0, 3.0}
	got := c.GetValues()

	for i, g := range got {
		w := want[i]
		if w != g {
			t.Errorf("GetValues should have returned %3.1f, but got %3.1f", w, g)
		}
	}

}
func TestGetCounts(t *testing.T) {
	var c Counter = Counter{
		CounterPair{Value: 1.0, Count: 10.0},
		CounterPair{Value: 2.0, Count: 20.0},
		CounterPair{Value: 3.0, Count: 30.0},
	}

	var want []float64 = []float64{10.0, 20.0, 30.0}
	got := c.GetCounts()

	for i, g := range got {
		w := want[i]
		if w != g {
			t.Errorf("GetValues should have returned %3.1f, but got %3.1f", w, g)
		}
	}
}
