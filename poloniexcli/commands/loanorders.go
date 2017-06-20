package commands

import (
	"os"
	"text/tabwriter"
	"text/template"

	log "github.com/sirupsen/logrus"
	"github.com/robvanmieghem/poloniex/poloniexclient"
)

//LoanOrdersCommand returns the orderbook for a given currencypair
type LoanOrdersCommand struct {
	Currency string
	Format   string
}

const loanOrdersHeader = "Rate\tAmount({{.Currency}})\tDuration\t  Sum({{.Currency}})\n"
const loanOrdersTableFormat = "{{printf \"%.3f\" .RatePercentage}}%\t{{printf \"%.8f\" .Order.Amount}}\t{{.Order.RangeMin}}-{{.Order.RangeMax}} Days\t  {{printf \"%.8f\" .Sum}}\n"

func formatLoanOrdersAsTable(loanorders *poloniexclient.LoanOrders) (err error) {
	log.Debug("Formatting as a table")

	t := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)
	defer func() {
		t.Flush()
	}()

	tmpl, err := template.New("").Parse(loanOrdersHeader)
	if err != nil {
		log.Error("Template parsing error: ", err)
		return
	}

	err = tmpl.Execute(t, loanorders)
	if err != nil {
		log.Error("Error executing template: ", err)
		return
	}

	tmpl, err = template.New("").Parse(loanOrdersTableFormat)
	if err != nil {
		log.Error("Template parsing error: ", err)
		return
	}
	var sum float64
	for _, order := range loanorders.Offers {
		sum += order.Amount
		data := struct {
			Order          poloniexclient.LoanOrder
			RatePercentage float64
			Sum            float64
		}{
			order,
			order.Rate * 100,
			sum,
		}
		err = tmpl.Execute(t, data)
		if err != nil {
			log.Error("Error executing template: ", err)
			return
		}
	}

	return
}

//Execute a LoanOrdersCommand
func (command *LoanOrdersCommand) Execute() (err error) {
	log.Debug("Executing LoanOrders Command:\n\tCurrency: ", command.Currency)

	c, err := poloniexclient.NewClient("", "")
	if err != nil {
		return
	}

	loanorders, err := c.ReturnLoanOrders(command.Currency)
	if err != nil {
		return
	}
	log.Debug(loanorders)
	if command.Format == FormatAsJSON {
		err = formatAsJSON(loanorders)
	} else {
		err = formatLoanOrdersAsTable(loanorders)
	}
	return
}
