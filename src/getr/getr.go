package getr

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/Viking2012/goraynor/data/basis"
	"github.com/Viking2012/goraynor/src/structs"
)

const baseUrl = "https://api.tiingo.com/tiingo/daily/"
const DefaultTimeout = time.Second * 8

type rawTiingoResponse struct {
	TickerDate string  `json:"date"`
	Close      float32 `json:"close,float32"`
	AdjClose   float32 `json:"adjClose"`
}

func (tR *rawTiingoResponse) calculatePercentChange(previousPrice *rawTiingoResponse) float64 {
	var pChange float64
	pChange = float64((tR.AdjClose - previousPrice.AdjClose) / previousPrice.AdjClose)
	return pChange
}

func DownloadTickers(saveDir string) error {
	var wg sync.WaitGroup

	for _, ticker := range basis.TICKERS {
		wg.Add(1)
		go downloadTicker(ticker, saveDir, &wg)
	}

	wg.Wait()
	return nil
}

func downloadTicker(ticker, saveDir string, wg *sync.WaitGroup) {
	fmt.Printf("getting %5s\n", ticker)
	_ = downloadMutualFundData(ticker, saveDir)
	wg.Done()
}

func downloadMutualFundData(ticker, saveDir string) error {
	// internal variables
	tickerUrl := baseUrl + ticker + "/prices"
	filename := filepath.Join(saveDir, ticker+".json")
	out, _ := os.Create(filename)
	defer out.Close()

	req, err := http.NewRequest(http.MethodGet, tickerUrl, nil)
	if err != nil {
		return err
	}

	// Query Variables
	today := time.Now().Format("2006-01-02")
	q := req.URL.Query()
	q.Add("endDate", today)
	q.Add("token", basis.ApiToken)
	q.Add("format", basis.ResponseFormatJson)
	q.Add("resampleFreq", basis.ResponseFrequencyMonthly)
	q.Add("columns", basis.ResponseColumns)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// save the data to file
	if _, err := io.Copy(out, resp.Body); err != nil {
		return err
	}

	return nil
}

func GetTickers(lookupDir string) (*structs.AllPerformers, error) {
	var allPerf structs.AllPerformers = make(structs.AllPerformers, len(basis.TICKERS))
	for _, ticker := range basis.TICKERS {
		data, err := getTicker(ticker, lookupDir)
		if err != nil {
			return nil, err
		}
		allPerf[ticker] = data
	}

	return &allPerf, nil
}

func getTicker(ticker, lookupDir string) (*structs.PriceRecords, error) {
	var records *structs.PriceRecords
	records, err := readDataIntoPriceRecords(ticker, lookupDir)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func readDataIntoPriceRecords(ticker, lookupDir string) (*structs.PriceRecords, error) {
	filename := filepath.Join(lookupDir, ticker+".json")
	in, err := os.OpenFile(filename, os.O_RDONLY, 0400) // 0400 becuase we only need to read from the file
	defer in.Close()

	// read the data into memory to convert it into a PriceRecords object
	body, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}

	var rawResp []rawTiingoResponse // to hold the API response call
	if err = json.Unmarshal(body, &rawResp); err != nil {
		return nil, err
	}

	records := make(structs.PriceRecords, len(rawResp)-1)

	for i := 1; i < len(rawResp); i++ {
		prevRecord := rawResp[i-1]
		thisRecord := rawResp[i]
		thisTime, err := time.Parse(time.RFC3339, thisRecord.TickerDate)
		if err != nil {
			return nil, err
		}
		records[i-1] = structs.PriceRecord{
			TickerDate:  thisTime,
			PriceReturn: thisRecord.calculatePercentChange(&prevRecord),
		}
	}

	return &records, nil
}
