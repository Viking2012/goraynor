package utils

import "time"

const layout = "2006-01-02"

// QuickParse simply reutrns a result of parsing a time string in YYYY-MM-DD format
func QuickParse(dateString string) time.Time {
	t, err := time.Parse(layout, dateString)
	if err != nil {
		panic(err)
	}
	return t
}
