package commands

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/robvanmieghem/poloniex/poloniexclient"
)

//OrderBookCommand returns the orderbook for a given currencypair
type OrderBookCommand struct {
	Credentials
	CurrencyPair string
	Depth        int
}

//Execute an OrderBookCommand
func (command *OrderBookCommand) Execute() (err error) {
	log.Debug("Executing OrderBook Command:\n\tCurrency pair: ", command.CurrencyPair)

	c, err := poloniexclient.NewClient(command.Credentials.Key, command.Credentials.Secret)
	if err != nil {
		return
	}

	orderbook, err := c.ReturnOrderBook(command.CurrencyPair, command.Depth)
	if err != nil {
		return
	}
	fmt.Println(orderbook)
	return
}
