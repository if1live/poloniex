package poloniexclient

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoanOrdersUnmarshalJSON(t *testing.T) {
	input := `{"offers":[{"rate":"0.00144800","amount":"1.70108234","rangeMin":2,"rangeMax":2}],"demands":[{"rate":"0.00100000","amount":"0.00073183","rangeMin":2,"rangeMax":2},{"rate":"0.00005000","amount":"14.15437297","rangeMin":2,"rangeMax":2}]}`
	loanorders := LoanOrders{}
	err := json.Unmarshal([]byte(input), &loanorders)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(loanorders.Offers))
	assert.Equal(t, 2, len(loanorders.Demands))
	assert.Equal(t, 0.00144800, loanorders.Offers[0].Rate)
	assert.Equal(t, 1.70108234, loanorders.Offers[0].Amount)
}
