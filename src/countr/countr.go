package countr

import "sort"

// CounterPair holds a key, value pair of numbers and counts (basically a map)
// generally, this is used to hold unique prices and their number of occurances
// but can be generalized to any key/value pair
type CounterPair struct {
	Value float64
	Count float64
}

// Counter is an array of CounterPair, which usually represents a list of all unique
// prices and the number of occurances within a data stream
type Counter []CounterPair

func (c Counter) Len() int           { return len(c) }
func (c Counter) Less(i, j int) bool { return c[i].Value < c[j].Value }
func (c Counter) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

// NewCounter takes a map and converts it into the internal Counter struct
// which implements the right sorting methods we need for downstream calculation
// Note: the return value is sorted for increasing order of values (not counts)
func NewCounter(m map[float64]float64) Counter {
	var c = make(Counter, len(m))
	i := 0
	for k := range m {
		c[i] = CounterPair{Value: k, Count: m[k]}
		i++
	}

	sort.Sort(c)
	return c
}

// Count takes an array of floats and returns a Counter of unique values and the
// number of occurances of these values
func Count(values []float64) Counter {
	var tempCounts = make(map[float64]float64)
	for _, v := range values {
		tempCounts[v]++
	}

	return NewCounter(tempCounts)
}

// Returns
func (c Counter) GetValues() []float64 {
	var v = make([]float64, c.Len())
	for i := 0; i < c.Len(); i++ {
		v[i] = c[i].Value
	}
	return v
}

func (c Counter) GetCounts() []float64 {
	var v = make([]float64, c.Len())
	for i := 0; i < c.Len(); i++ {
		v[i] = c[i].Count
	}
	return v
}
