package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	tickerList [9]string
	m          float64
)

type symbolPrice struct {
	Price string `json:"price"`
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	tickerList[0] = "BTTBNB"
	tickerList[1] = "BTTBRL"
	tickerList[2] = "BTTBUSD"
	tickerList[3] = "BTTEUR"
	tickerList[4] = "BTTTRX"
	tickerList[5] = "BTTTRY"
	tickerList[6] = "BTTTUSD"
	tickerList[7] = "BTTUSDC"
	tickerList[8] = "BTTUSDT"

	if len(os.Args) > 1 {
		howMatch := os.Args[1]
		if s, err := strconv.ParseFloat(howMatch, 64); err == nil {
			m = s
		}

	} else {
		fmt.Println("Sorry, but you are not set how tokens need to count, as sample ./bttto.exe 11.11")
		fmt.Println("Here default price for 1 (one token): ")
		m = 1.00
	}

	for l := 0; l < len(tickerList); l++ {
		getAvgPricebySymbol(tickerList[l], m)
	}

}

func getAvgPricebySymbol(symbol string, countTokens float64) {
	resp, err := http.Get("https://api.binance.com/api/v3/avgPrice?symbol=" + symbol)
	checkErr(err)

	body, err := io.ReadAll(resp.Body)
	checkErr(err)

	var getSymbolPriceRecord symbolPrice

	json.Unmarshal(body, &getSymbolPriceRecord)

	if s, err := strconv.ParseFloat(getSymbolPriceRecord.Price, 64); err == nil {
		fmt.Println(symbol+" : ", countTokens*s)
	}

}
