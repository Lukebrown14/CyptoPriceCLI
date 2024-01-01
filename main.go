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

func apiGetRequest(nameInput string, currencyInput string) (string, string, float32) {

	client := &http.Client {}
	req, err := http.NewRequest("GET", "https://rest.coinapi.io/v1/exchangerate/"+ nameInput + "/" + currencyInput, nil)

	checkError(err)

	req.Header.Add("Accept", "text/plain")
	req.Header.Add("X-CoinAPI-Key", "03F605A1-F710-41A8-8D5A-601C87BFD070")
	
	res, err := client.Do(req)

	checkError(err)
	
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	checkError(err)

	var cypto CyptoPrice
	err = json.Unmarshal(body, &cypto)

	checkError(err)
	
	// return data 
	name, currency, price := cypto.Name, cypto.Currency, cypto.Price

	return name, currency, price
}

func checkError (err error){
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {

	// Defining Variables 
	cyptoInput := "BTC"
	currencyInput := "GBP"
	
	// Handling Flags 
	if len(os.Args) >= 2 {
		cyptoInput = os.Args[1]
		currencyInput = os.Args[2]
	}

	// API Request
	name, currency, price := apiGetRequest(cyptoInput, currencyInput)

	switch currency {
	case "GBP":
		currency = "£"
	case "USD":
		currency = "$"
	default:
		fmt.Println("Only GBP or USD Currency Allowed")
		return
	}


	// Outputting results to CLI 
	fmt.Printf("%s%.0f %s \n", currency, price, name)
}