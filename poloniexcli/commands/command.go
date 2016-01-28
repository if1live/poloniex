package commands

import (
	"encoding/json"
	"os"

	"github.com/prometheus/common/log"
)

//Credentials are the poloniex api credentials
type Credentials struct {
	Key    string
	Secret string
}

//Command is a toplevel command to be executed by the cli's main routine
type Command interface {
	//Execute the command
	Execute() error
}

const (
	//FormatAsTable makes the command output the result in a nice table
	FormatAsTable = "table"
	//FormatAsJSON makes the command output the result as json
	FormatAsJSON = "json"
)

func formatAsJSON(data interface{}) error {
	log.Debug("Formatting as json")
	return json.NewEncoder(os.Stdout).Encode(data)
}
