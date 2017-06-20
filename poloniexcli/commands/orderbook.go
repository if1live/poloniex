package commands

import (
	"os"
	"strings"
	"text/tabwriter"
	"text/template"

	log "github.com/sirupsen/logrus"
	"github.com/robvanmieghem/poloniex/poloniexclient"
)

//OrderBookCommand returns the orderbook for a given currencypair
type OrderBookCommand struct {
	CurrencyPair string
	Depth        int
	Format       string
}

const orderbookHeader = "Sell      \t\t\t\tBuy       \t\t\t\t\nPrice     \t{{.Currency}}       \t{{.BaseCurrency}}       \tSum({{.BaseCurrency}})  \tPrice     \t{{.Currency}}       \t{{.BaseCurrency}}       \t  Sum({{.BaseCurrency}})\n"
const orderbookTableFormat = "{{printf \"%.8f\" .Ask.Price}}\t{{printf \"%.8f\" .Ask.Amount}}\t{{printf \"%.8f\" .AskTotal}}\t{{printf \"%.8f\" .AskSum}}\t{{printf \"%.8f\" .Bid.Price}}\t{{printf \"%.8f\" .Bid.Amount}}\t{{printf \"%.8f\" .BidTotal}}\t  {{printf \"%.8f\" .BidSum}}\n"

func formatOrderBookAsTable(orderbook *poloniexclient.OrderBook) (err error) {
	log.Debug("Formatting as a table")

	t := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)
	defer func() {
		t.Flush()
	}()

	tmpl, err := template.New("").Parse(orderbookHeader)
	if err != nil {
		log.Error("Template parsing error: ", err)
		return
	}
	var headerData struct {
		Currency     string
		BaseCurrency string
	}
	currencies := strings.Split(orderbook.CurrencyPair, "_")
	headerData.BaseCurrency = currencies[0]
	headerData.Currency = currencies[1]

	err = tmpl.Execute(t, headerData)
	if err != nil {
		log.Error("Error executing template: ", err)
		return
	}

	tmpl, err = template.New("").Parse(orderbookTableFormat)
	if err != nil {
		log.Error("Template parsing error: ", err)
		return
	}

	var askSum, bidSum float64
	for i := 0; i < len(orderbook.Asks) || i < len(orderbook.Bids); i++ {
		askEntry := orderbook.Asks[i]
		bidEntry := orderbook.Bids[i]
		//TODO: if not equal number of Asks and Bids
		askTotal := askEntry.Price * askEntry.Amount
		bidTotal := bidEntry.Price * bidEntry.Amount
		askSum += askTotal
		bidSum += bidTotal
		data := struct {
			Ask      poloniexclient.OrderBookEntry
			Bid      poloniexclient.OrderBookEntry
			AskTotal float64
			AskSum   float64
			BidTotal float64
			BidSum   float64
		}{
			askEntry,
			bidEntry,
			askTotal,
			askSum,
			bidTotal,
			bidSum,
		}
		err = tmpl.Execute(t, data)
		if err != nil {
			log.Error("Error executing template: ", err)
			return
		}
	}
	return
}

//Execute an OrderBookCommand
func (command *OrderBookCommand) Execute() (err error) {
	log.Debug("Executing OrderBook Command:\n\tCurrency pair: ", command.CurrencyPair, "\n\tDepth: ", command.Depth)

	c, err := poloniexclient.NewClient("", "")
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
		err = formatOrderBookAsTable(orderbook)
	}
	return
}
