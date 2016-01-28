package poloniexclient

import (
	"encoding/json"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
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

	req, err := http.NewRequest("GET", poloniexPublicAPIUrl, nil)
	if err != nil {
		return
	}
	query := req.URL.Query()
	query.Set("command", "returnLoanOrders")
	query.Set("currency", currency)
	req.URL.RawQuery = query.Encode()

	logRequest(req)
	resp, err := poloniexClient.httpClient.Do(req)
	logResponse(resp, err)

	if err != nil {
		return
	}

	defer resp.Body.Close()
	loanorders = new(LoanOrders)
	loanorders.Currency = currency
	err = json.NewDecoder(resp.Body).Decode(&loanorders)
	return
}
