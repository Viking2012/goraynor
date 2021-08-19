package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	data "github.com/Viking2012/goraynor/data/tickers"
	"github.com/Viking2012/goraynor/src/getr"
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

func downloadTickers(saveDir string) error {
	var wg sync.WaitGroup

	for _, ticker := range data.TICKERS {
		wg.Add(1)
		go downloadTicker(ticker, saveDir, &wg)
	}

	wg.Wait()
	return nil
}

func downloadTicker(ticker, saveDir string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("getting %5s\n", ticker)
	getr.DownloadMutualFundData(ticker, saveDir)
}

func main() {
	today := time.Now().Format("20060102")
	saveDir := filepath.Join(".", "data", today)
	_ = os.Mkdir(saveDir, os.ModeDir) // TODO(ajo): lazy ignoring of errors. Fix This!

	_ = downloadTickers(saveDir)

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
