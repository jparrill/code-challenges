/*
👋 Welcome Gophers! In this Go challenge, you are going to be working with JSON data and implementing a function that takes in a JSON string full of stock quotes and returns the one with the highest dividend return.

You will need to implement the HighestDividend function so that it takes in a JSON string. You will then need to somehow convert this JSON string into something you can traverse in Go using the encoding/json package.

Finally, you will need to return the string ticker for the stock that has the highest dividend.
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

type Stocks struct {
	Stocks []Stocks
}

type Stock struct {
	Ticker   string  `json:"ticker"`
	Dividend float64 `json:'dividend"`
}

type ByDiv []Stock

func (stocks ByDiv) Len() int           { return len(stocks) }
func (stocks ByDiv) Swap(x, y int)      { stocks[x], stocks[y] = stocks[y], stocks[x] }
func (stocks ByDiv) Less(x, y int) bool { return stocks[x].Dividend > stocks[y].Dividend }

// HighestDividend iterates through a JSON string of stocks
// unmarshalls them into a struct and returns the Stock with
// the highest dividend
func HighestDividend(stocks_json string) string {
	var sts []Stock

	if err := json.Unmarshal([]byte(stocks_json), &sts); err != nil {
		log.Fatal(err)
	}

	sort.Sort(ByDiv(sts))

	return sts[0].Ticker
}

func main() {
	fmt.Println("Stock Price AI")

	stocks_json := `[
    {"ticker":"APPL","dividend":0.5},
    {"ticker":"GOOG","dividend":0.2},
    {"ticker":"MSFT","dividend":0.3}
  ]`

	fmt.Println(HighestDividend(stocks_json))
}
