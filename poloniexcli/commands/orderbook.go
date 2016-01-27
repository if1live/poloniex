package commands

import (
	"encoding/json"
	"errors"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/robvanmieghem/poloniex/poloniexclient"
)

const (
	//FormatAsTable makes the command output the result in a nice table
	FormatAsTable = "table"
	//FormatAsJSON makes the command output the result as json
	FormatAsJSON = "json"
)

//OrderBookCommand returns the orderbook for a given currencypair
type OrderBookCommand struct {
	Credentials
	CurrencyPair string
	Depth        int
	Format       string
}

func formatAsTable(orderbook *poloniexclient.OrderBook) error {
	log.Debug("Formatting as a table")
	return errors.New("Table output is not implemented")
}

func formatAsJSON(orderbook *poloniexclient.OrderBook) error {
	log.Debug("Formatting as json")
	return json.NewEncoder(os.Stdout).Encode(orderbook)
}

//Execute an OrderBookCommand
func (command *OrderBookCommand) Execute() (err error) {
	log.Debug("Executing OrderBook Command:\n\tCurrency pair: ", command.CurrencyPair, "\n\tDepth: ", command.Depth)

	c, err := poloniexclient.NewClient(command.Credentials.Key, command.Credentials.Secret)
	if err != nil {
		return
	}

	orderbook, err := c.ReturnOrderBook(command.CurrencyPair, command.Depth)
	if err != nil {
		return
	}
	log.Debug(orderbook)
	if command.Format == FormatAsJSON {
		err = formatAsJSON(orderbook)
	} else {
		err = formatAsTable(orderbook)
	}
	return
}
