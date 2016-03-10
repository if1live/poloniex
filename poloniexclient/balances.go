package poloniexclient

import "github.com/prometheus/common/log"

//ReturnBalances returns the available balances
func (poloniexClient *PoloniexClient) ReturnBalances() (balances *map[string]float64, err error) {
	log.Debug("ReturnBalances")

	commandParameters := make(map[string]string)
	poloniexBalances := make(map[string]string)

	err = poloniexClient.executeTradingAPICommand("returnBalances", commandParameters, poloniexBalances)
	if err != nil {
		return
	}

	return
}
