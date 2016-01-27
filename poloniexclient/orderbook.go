package poloniexclient

import (
	"encoding/json"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

//OrderBookEntry is a value/amount combination representing an entry in an OrderBook
type OrderBookEntry struct {
	Price  float64
	Amount float64
}

//UnmarshalJSON is a custom unmarshaller to handle poloniex special json format
func (entry *OrderBookEntry) UnmarshalJSON(b []byte) (err error) {
	var poloEntry []interface{}
	err = json.Unmarshal(b, &poloEntry)
	if err != nil {
		return
	}

	entry.Price, err = strconv.ParseFloat(poloEntry[0].(string), 64)
	entry.Amount = poloEntry[1].(float64)
	return
}

//OrderBook is the asks and bids for a currency pair
type OrderBook struct {
	CurrencyPair string
	Asks         []OrderBookEntry
	Bids         []OrderBookEntry
}

//ReturnOrderBook return the orderbook for a specific currencypair up to a certain depth
func (poloniexClient *PoloniexClient) ReturnOrderBook(currencypair string, depth int) (orderbook *OrderBook, err error) {
	log.Debug("ReturnOrderBook")

	req, err := http.NewRequest("GET", "http://poloniex.com/public", nil)
	if err != nil {
		return
	}
	query := req.URL.Query()
	query.Set("command", "returnOrderBook")
	query.Set("currencyPair", currencypair)
	query.Set("depth", strconv.Itoa(depth))
	req.URL.RawQuery = query.Encode()

	logRequest(req)
	resp, err := poloniexClient.httpClient.Do(req)
	logResponse(resp, err)

	if err != nil {
		return
	}

	defer resp.Body.Close()
	orderbook = new(OrderBook)
	orderbook.CurrencyPair = currencypair
	err = json.NewDecoder(resp.Body).Decode(&orderbook)
	return
}
