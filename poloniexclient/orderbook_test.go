package poloniexclient

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderBookUnmarshalJSON(t *testing.T) {
	input := `{"asks":[["0.00633989",24.25363453],["0.00633990",76.61232516]],"bids":[["0.00630016",18.28449166],["0.00630015",82.83]],"isFrozen":"0"}`
	orderbook := OrderBook{}
	err := json.Unmarshal([]byte(input), &orderbook)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(orderbook.Asks))
	assert.Equal(t, 0.00633989, orderbook.Asks[0].Price)
	assert.Equal(t, 24.25363453, orderbook.Asks[0].Amount)
}
