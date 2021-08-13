package quantilr

import (
	"errors"
	"sort"

	"gonum.org/v1/gonum/stat"
)

var (
	NoWeightsError             error = errors.New("Encountered a decile with no weights (counts or frequencies were all 0)")
	NonStochasticWeights       error = errors.New("weights for decile provided did not (or could not scaled to) sum to 1")
	ValueNotFound              error = errors.New("the value provided does not appear in the calculated decile range")
	WarnDecilesNotDeduplicated error = errors.New("The decile pairs provided were not deduplicated - this was performed automatically and then resorted")
	WarnDecilesNotSorted       error = errors.New("The decile pairs provided were not deduplicated - this was performed automatically")
)

type CountedPairs interface {
	GetValues() []float64
	GetCounts() []float64
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type DecilePair struct {
	Decile int8
	Weight float64
}
type Deciles struct {
	Pairs     []DecilePair
	isSorted  bool
	isDeduped bool
}

func (d Deciles) Len() int           { return len(d.Pairs) }
func (d Deciles) Less(i, j int) bool { return d.Pairs[i].Decile < d.Pairs[j].Decile }
func (d Deciles) Swap(i, j int)      { d.Pairs[i], d.Pairs[j] = d.Pairs[j], d.Pairs[i] }
func (d *Deciles) Sort() {
	if !d.isSorted {
		sort.Sort(d)
		d.isSorted = true
	}
}

func NewDeciles(c CountedPairs, deduplicateDeciles bool) (Deciles, error) {
	deciles, pointValues := Quantiles(c, []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0})

	var d Deciles = Deciles{
		Pairs:     make([]DecilePair, len(deciles)),
		isSorted:  false,
		isDeduped: false,
	}

	for i := 0; i < len(deciles); i++ {
		thisDecile := int8(deciles[i] * 10)
		thisWeight := pointValues[i]
		d.Pairs[i] = DecilePair{Decile: thisDecile, Weight: thisWeight}
	}

	if deduplicateDeciles {
		d.depulicateDeciles()
	}

	// sort as a precuation that anything went awry in generating the deciles
	d.Sort()

	return d, nil
}

func (d *Deciles) depulicateDeciles() {
	var deduped = make(map[float64]int8)
	if !d.isSorted {
		sort.Sort(d)
		d.isSorted = true
	}
	for i := 0; i < d.Len(); i++ {
		thisPair := d.Pairs[i]
		thisWeight := thisPair.Weight
		thisDecile := thisPair.Decile

		deduped[thisWeight] = thisDecile
	}

	var newDeciles []DecilePair = make([]DecilePair, len(deduped))
	i := 0
	for value, decile := range deduped {
		newPair := DecilePair{Decile: decile, Weight: value}
		newDeciles[i] = newPair
		i++
	}

	d.Pairs = newDeciles
	d.isSorted = false
}

func Quantiles(c CountedPairs, probs []float64) (quantiles, pointValues []float64) {
	// return value instantiation
	quantiles = make([]float64, len(probs))
	pointValues = make([]float64, len(probs))

	// we sort here to ensure that the values are in increasing order,
	// which is required for calculating and collecting quantiles
	// However, this is probably more a guard than a requirement, since
	// the Countr package should sort the CountedPairs prior to return.
	sort.Sort(c)

	values := c.GetValues()
	weights := c.GetCounts()

	for i, p := range probs {
		quantiles[i] = p
		pointValues[i] = stat.Quantile(p, stat.Empirical, values, weights)
	}

	return quantiles, pointValues
}

func (d Deciles) ScaleToOne() error {
	var denominator float64
	for i := 0; i < d.Len(); i++ {
		e := &d.Pairs[i]
		denominator += e.Weight
	}

	if denominator == 0 {
		return NoWeightsError
	}

	for i := 0; i < d.Len(); i++ {
		e := &d.Pairs[i] // pointer to underlying struct, since we need to modify it
		// we don't need to check for zero division here, since we
		// check for zero weights (totalValue == 0) above
		scaledValue := e.Weight / denominator
		e.Weight = scaledValue
	}

	var finalValue int8
	for i := 0; i < d.Len(); i++ {
		e := &d.Pairs[i]
		finalValue += int8(e.Weight)
	}

	if finalValue != 1 {
		return NonStochasticWeights
	}

	return nil
}

func (d Deciles) LookupValue(v float64) (decileOfValue int8, errorWarning error) {
	errorWarning = nil
	var sortWarningSafeToApply bool = true // needed so that the sorting done after deduping (if performed) is not overwritten
	if !d.isDeduped {
		d.depulicateDeciles()
		errorWarning = WarnDecilesNotDeduplicated
		sortWarningSafeToApply = false
	}
	if !d.isSorted {
		d.Sort()
		if sortWarningSafeToApply {
			errorWarning = WarnDecilesNotSorted
		}
	}
	for i := 0; i < d.Len(); i++ {
		if v <= d.Pairs[i].Weight {
			return d.Pairs[i].Decile, errorWarning
		}
	}

	return -1, ValueNotFound
}
