# go Poloniex client

[![Build Status](https://travis-ci.org/robvanmieghem/poloniex.svg?branch=master)](https://travis-ci.org/robvanmieghem/poloniex)

## Use as a client library

`import github.com/robvanmieghem/poloniex/poloniexclient`

## Command line client

As an example to use the library, a command line poloniex client is implemented in the `poloniexcli` folder.
Execute without parameters of with `-h` will show a help page.

Getting the order book for the BTC_ETH currency pair and a depth of 10 for example:
```
$ ./poloniexcli orderbook -c BTC_ETH -d 10
  Sell                                              Buy
  Price         ETH         BTC         Sum(BTC)    Price         ETH         BTC         Sum(BTC)
  0.00629927   24.41003084  0.15376537  0.15376537  0.00624002    1.19000000  0.00742562  0.00742562
  0.00629928    0.49882630  0.00314225  0.15690762  0.00624001   87.91974500  0.54862009  0.55604571
  0.00629929   54.00889312  0.34021768  0.49712530  0.00624000    0.28961699  0.00180721  0.55785292
  0.00629930  100.00000000  0.62993000  1.12705530  0.00623420    0.64405726  0.00401518  0.56186810
  0.00629992  107.51045736  0.67730728  1.80436258  0.00623419    1.19000000  0.00741869  0.56928679
  0.00629997    0.66629500  0.00419764  1.80856022  0.00623418    5.27108140  0.03286087  0.60214766
  0.00630000  286.35120921  1.80401262  3.61257284  0.00623410  100.00000000  0.62341000  1.22555766
  0.00630014    1.39495095  0.00878839  3.62136123  0.00623408    0.94596804  0.00589724  1.23145490
  0.00630015   59.07936611  0.37220887  3.99357009  0.00623405  307.76046551  1.91859413  3.15004903
  0.00630025    5.52252123  0.03479326  4.02836336  0.00623401   20.04000000  0.12492956  3.27497859
```

By default, the output is in a table format but by passing the format parameter, you can also have it output json.
