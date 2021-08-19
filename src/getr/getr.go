package getr

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const baseUrl = "https://query1.finance.yahoo.com/v7/finance/download/"
const DefaultTimeout = time.Second * 8

func DownloadMutualFundData(ticker, saveDir string) error {
	tickerUrl := baseUrl + ticker
	filename := filepath.Join(saveDir, ticker+".csv")
	out, _ := os.Create(filename)
	defer out.Close()

	req, err := http.NewRequest(http.MethodGet, tickerUrl, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("period1", "0")
	q.Add("period2", "9999999999")
	q.Add("interval", "1mo")
	q.Add("events", "history")
	q.Add("includeAdjustedClose", "true")
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return err
	}

	return nil

}
