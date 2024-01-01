package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "io/ioutil"
)

// type CyptoPrice struct {
//   Name string `json:"asset_id_base"`
//   Currency string `json:"asset_id_quote"`
//   Price float32 `json:"rate"`
// }

type HistoricalPrice struct {
  Price []interface{} `json:"prices"`
}

func main() {


  client := &http.Client {
  }
  req, err := http.NewRequest("GET", "https://api.coingecko.com/api/v3/coins/bitcoin/market_chart?vs_currency=usd&days=1&interval=daily&precision=6", nil)

  if err != nil {
    fmt.Println(err)
    return
  }

  res, err := client.Do(req)

  if err != nil {
    fmt.Println(err)
    return
  }

  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  
  if err != nil {
    fmt.Println(err)
    return
  }

  var histroy HistoricalPrice
  err = json.Unmarshal(body, &histroy)

  historicalPrice := histroy.Price[0]


  var interfaceSlice []interface{} = make([]interface{}, len(historicalPrice))
  for i, d := range historicalPrice {
    interfaceSlice[i] = d
    fmt.Println(interfaceSlice[i])
  }

  fmt.Println(historicalPrice)

}

// {"prices":[[1704067200000,42208.2],[1704125232000,42773.36]],"market_caps":[[1704067200000,827596236151.1959],[1704125232000,838184180847.9264]],"total_volumes":[[1704067200000,14183728910.169804],[1704125232000,13108507724.910725]]}



