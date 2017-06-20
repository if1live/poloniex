package commands

import (
	"os"
	"text/tabwriter"
	"text/template"

	log "github.com/sirupsen/logrus"
	"github.com/robvanmieghem/poloniex/poloniexclient"
)

//BalancesCommand returns the available balances
type BalancesCommand struct {
	Credentials
	Format string
}

func formatBalancesAsTable(balances *map[string]float64) (err error) {
	log.Debug("Formatting as a table")

	t := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)
	defer func() {
		t.Flush()
	}()

	tmpl, err := template.New("").Parse("{{.Currency}}\t {{printf \"%.8f\" .Amount}}")
	if err != nil {
		log.Error("Template parsing error: ", err)
		return
	}
	for currency, amount := range *balances {

		data := struct {
			Currency string
			Amount   float64
		}{
			currency,
			amount,
		}
		err = tmpl.Execute(t, data)
		if err != nil {
			log.Error("Error executing template: ", err)
			return
		}
	}

	return
}

//Execute an BalancesCommand
func (command *BalancesCommand) Execute() (err error) {
	log.Debug("Executing BalancesCommand Command")

	c, err := poloniexclient.NewClient(command.Credentials.Key, command.Credentials.Secret)
	if err != nil {
		return
	}

	balances, err := c.ReturnBalances()
	if err != nil {
		return
	}
	log.Debug(balances)
	if command.Format == FormatAsJSON {
		err = formatAsJSON(balances)
	} else {
		err = formatBalancesAsTable(balances)
	}
	return
}
