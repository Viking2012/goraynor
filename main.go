package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Viking2012/goraynor/src/countr"
	"github.com/Viking2012/goraynor/src/getr"
	"github.com/Viking2012/goraynor/src/quantilr"
)

// const randSeed uint64 = 123456

// func decileToIndex(d int) int {
// 	return d - 1
// }

// func indexToDecile(i int) int {
// 	return i + 1
// }

// type sparseArray map[int]float64

// func newSparseArray() sparseArray {
// 	sp := sparseArray{}
// 	for i := 1; i <= 10; i++ {
// 		sp[i] = 0
// 	}
// 	return sp
// }

// type transitionMatrix map[int]sparseArray

// func newTransitionMatrix() transitionMatrix {
// 	tm := transitionMatrix{}
// 	for i := 1; i <= 10; i++ {
// 		tm[i] = newSparseArray()
// 	}

// 	return tm
// }

// func (tm transitionMatrix) Add(i, j int) {
// 	currentCount := tm[i][j]
// 	currentCount += 1
// 	tm[i][j] = currentCount
// }

// func (tm transitionMatrix) Load(transitions []decileTransition) {
// 	for i := 0; i < len(transitions); i++ {
// 		t := transitions[i]
// 		tm.Add(t.this, t.next)
// 	}
// }

// type transitionModel map[int]distuv.Categorical

// func newTransitionModel(tm transitionMatrix) transitionModel {
// 	newModel := transitionModel{}
// 	src := rand.NewSource(randSeed)

// 	for thisDecile := 1; thisDecile <= 10; thisDecile++ {
// 		// collect the weights from this row of the transitionMatrix
// 		weights := make([]float64, 10)
// 		for thisWeight := 0; thisWeight < 10; thisWeight++ {
// 			weights[thisWeight] = tm[thisDecile][indexToDecile(thisWeight)]
// 		}

// 		newModel[decileToIndex(thisDecile)] = distuv.NewCategorical(weights, src)
// 	}

// 	return newModel
// }

// func (tm transitionModel) Simulate(start int8, lifespan int16) countr.Counter {
// 	var simulationResults []float64 = make([]float64, lifespan)
// 	var thisPeriod int16
// 	var thisStep int8 = start
// 	// var randSource := rand.New
// 	for thisPeriod = 0; thisPeriod < lifespan; thisPeriod++ {
// 		fmt.Printf("For period %d, beginning simulation with value %d", thisPeriod, thisStep)
// 		simulationResults[thisPeriod] = float64(thisStep + 1)
// 		modelToBeUsed := tm[int(thisStep)-1]
// 		thisStep = int8(modelToBeUsed.Rand())
// 		fmt.Printf("For period %d, after simulation got value %d", thisPeriod, thisStep)
// 	}
// 	return countr.Count(simulationResults)
// }

// type decileTransition struct {
// 	this int
// 	next int
// }

// var transitions []decileTransition = []decileTransition{
// 	{this: 1, next: 2},
// 	{this: 2, next: 3},
// 	{this: 3, next: 4},
// 	{this: 4, next: 5},
// 	{this: 5, next: 1},
// 	{this: 1, next: 2},
// 	{this: 2, next: 3},
// 	{this: 3, next: 4},
// 	{this: 4, next: 2},
// 	{this: 2, next: 1},
// }

func main() {
	today := "20210819" // otherwise, time.Now().Format("20060102")
	saveDir := filepath.Join(".", "data", today)
	_ = os.Mkdir(saveDir, os.ModeDir) // TODO(ajo): lazy ignoring of errors. Fix This!

	// err := getr.DownloadTickers(saveDir)
	// if err != nil {
	// 	panic(err)
	// }

	pRecords, err := getr.GetTickers(saveDir)
	if err != nil {
		panic(err)
	}

	var allPrices []float64
	for _, data := range *pRecords {
		records := *data
		// lastRecord := records[len(records)-1]
		for i := 0; i < len(records); i++ {
			allPrices = append(allPrices, records[i].PriceReturn)
		}
		// fmt.Printf("\tfor ticker: %s, got %5d monthly records (%v)\n", ticker, len(*data), lastRecord)
	}

	c := countr.Count(allPrices)
	d, err := quantilr.NewDeciles(c, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("All Prices\n%v\n", d.Pairs)
	pRecords.SetDeciles(&d)

	for ticker, data := range *pRecords {
		fmt.Printf("\tfor ticker: %s, got %5d monthly records and returns:\n", ticker, len(*data))
		records := *data
		for i := 0; i < len(records); i++ {
			thisRecord := records[i]
			if thisRecord.DecileOfPrice != 0 {
				fmt.Printf("\t\t(%v)\n", thisRecord)
			}
		}
	}

	// raw, err := readr.ParseCSV("./test/test_data.csv", 1, &readr.DefaultFieldMap)
	// if err != nil {
	// 	panic(err)
	// }

	// organizr.OrderedBy(organizr.ByProduct, organizr.ByCustomer, organizr.ByDate, organizr.ByDocumentNumber, organizr.ByDocumentLineNumber).Sort(raw)

	// newMatr := newTransitionMatrix()
	// newMatr.Load(transitions)

	// fmt.Println(newMatr)

	// _ = newTransitionModel(newMatr)

	// ALL PREVIOUSLY COMMENTED
	// var numSimulations int = 1
	// var cp countr.Counter
	// for i := 0; i < numSimulations; i++ {
	// 	cp = models.Simulate(1, 100)
	// }
	// fmt.Println(cp)

	// toCounter := make([]float64, numSimulations)

	// weights := []float64{1, 10, 89}

	// source := rand.NewSource(randSeed)
	// dist := distuv.NewCategorical(weights, source)
	// for i := 0; i < numSimulations; i++ {
	// 	r := dist.Rand()
	// 	toCounter[i] = r + 1 // we use plus one here, since the categorical distrubution is zeo indexed
	// }

	// cp := countr.Count(toCounter)

	// fmt.Println(cp)

}
