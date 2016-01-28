package poloniexclient

import (
	"encoding/json"
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

//ReturnOrderBook returns the orderbook for a specific currencypair up to a certain depth
func (poloniexClient *PoloniexClient) ReturnOrderBook(currencypair string, depth int) (orderbook *OrderBook, err error) {
	log.Debug("ReturnOrderBook")

	orderbook = new(OrderBook)
	orderbook.CurrencyPair = currencypair

	commandParameters := make(map[string]string)

	commandParameters["currencyPair"] = currencypair
	commandParameters["depth"] = strconv.Itoa(depth)

	err = poloniexClient.executePublicAPICommand("returnOrderBook", commandParameters, orderbook)
	return
}
