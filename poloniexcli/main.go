package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/robvanmieghem/poloniex/poloniexcli/commands"
)

//Version
var Version = "0.1-Dev"

//ApplicationName is the name of the application
var ApplicationName = "Poloniex client"

func main() {

	var (
		credentials      = commands.Credentials{}
		orderbookCommand = &commands.OrderBookCommand{}
	)

	app := cli.NewApp()
	app.Name = ApplicationName
	app.Version = Version
	app.Usage = "Poloniex command line tool to demonstrate the usage of the api client"

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	var debugLogging bool
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "Enable debug logging",
			Destination: &debugLogging,
		},
		cli.StringFlag{
			Name:        "apikey",
			Usage:       "Poloniex API key",
			Destination: &credentials.Key,
			EnvVar:      "POLONIEX_API_KEY",
		},

		cli.StringFlag{
			Name:        "apisecret",
			Usage:       "Poloniex API secret",
			Destination: &credentials.Secret,
			EnvVar:      "POLONIEX_API_SECRET",
		},
	}
	app.Before = func(c *cli.Context) error {
		if debugLogging {
			log.SetLevel(log.DebugLevel)
			log.Debug("Debug logging enabled")
			log.Debug(ApplicationName, "-", Version)
		}
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:  "orderbook",
			Usage: "Print the orderbook",
			Action: func(c *cli.Context) {
				orderbookCommand.Credentials = credentials
				if err := orderbookCommand.Execute(); err != nil {
					log.Error(err)
				}
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "currencypair,c",
					Usage:       "Currency pair",
					Value:       "BTC_ETH",
					Destination: &orderbookCommand.CurrencyPair,
				},
				cli.IntFlag{
					Name:        "depth, d",
					Usage:       "Depth of the orderbook to print",
					Value:       50,
					Destination: &orderbookCommand.Depth,
				},
			},
		},
	}

	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}

	app.Run(os.Args)

}
