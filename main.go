package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CyptoPrice struct {
	Name string `json:"asset_id_base"`
	Currency string `json:"asset_id_quote"`
	Price float32 `json:"rate"`
}

func main() {

	// Defining Variables 
	cyptoInput := "BTC"
	currencyInput := "GBP"
	client := &http.Client {}


	// Handling Flags 
	if len(os.Args) >= 2 {
		cyptoInput = os.Args[1]
		currencyInput = os.Args[2]
	}

	// API Request
	req, err := http.NewRequest("GET", "https://rest.coinapi.io/v1/exchangerate/"+ cyptoInput + "/" + currencyInput, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Accept", "text/plain")
	req.Header.Add("X-CoinAPI-Key", "03F605A1-F710-41A8-8D5A-601C87BFD070")
	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	// Handling Data
	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	var cypto CyptoPrice
	err = json.Unmarshal(body, &cypto)

	if err != nil {
		fmt.Println(err)
		return
	}

	name, currency, price := cypto.Name, cypto.Currency, cypto.Price

	// Outputting results to CLI 
	fmt.Printf("%s: %s %.0f\n", name, currency, price)
}