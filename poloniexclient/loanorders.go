package poloniexclient

import (
	"encoding/json"
	"strconv"

	log "github.com/sirupsen/logrus"
)

//LoanOrder is a single an order (offer or demand) for a specific rate
type LoanOrder struct {
	Rate     float64
	Amount   float64
	RangeMin int
	RangeMax int
}

//UnmarshalJSON is a custom unmarshaller to convert the strings poloniex returns as Rate and Amount to floats
func (order *LoanOrder) UnmarshalJSON(b []byte) (err error) {
	var poloOrder struct {
		Rate     string
		Amount   string
		RangeMin int
		RangeMax int
	}
	err = json.Unmarshal(b, &poloOrder)
	if err != nil {
		return
	}
	order.Amount, err = strconv.ParseFloat(poloOrder.Amount, 64)
	if err != nil {
		return
	}
	order.Rate, err = strconv.ParseFloat(poloOrder.Rate, 64)
	if err != nil {
		return
	}
	order.RangeMin = poloOrder.RangeMin
	order.RangeMax = poloOrder.RangeMax
	return
}

//LoanOrders are the loan offers and demands for a specific currency
type LoanOrders struct {
	Currency string
	Offers   []LoanOrder
	Demands  []LoanOrder
}

//ReturnLoanOrders returns the loan offers and demands for a specific currency
func (poloniexClient *PoloniexClient) ReturnLoanOrders(currency string) (loanorders *LoanOrders, err error) {
	log.Debug("ReturnLoanOrders")

	loanorders = new(LoanOrders)
	loanorders.Currency = currency

	commandParameters := map[string]string{
		"currency": currency,
	}

	err = poloniexClient.executePublicAPICommand("returnLoanOrders", commandParameters, loanorders)

	return
}
