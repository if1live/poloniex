package poloniexclient

import (
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

//PoloniexClient is a client to the poloniex (https://www.poloniex.com) api
type PoloniexClient struct {
	Key        string
	Secret     string
	httpClient *http.Client
}

//NewClient creates a new poloniex client given the api key and secret
func NewClient(key, secret string) (client *PoloniexClient, err error) {
	log.Debug("Creating poloniex client")
	client = &PoloniexClient{
		Key:        key,
		Secret:     secret,
		httpClient: &http.Client{}}
	return
}

func logRequest(request *http.Request) {
	if log.GetLevel() == log.DebugLevel {
		data, _ := httputil.DumpRequest(request, true)
		log.Debugf("%s\n\n", data)
	}
}

func logResponse(response *http.Response, err error) {
	if err != nil {
		log.Error("HTTP reponse error: ", err)
		return
	}
	if log.GetLevel() == log.DebugLevel {
		data, _ := httputil.DumpResponse(response, true)
		log.Debugf("%s\n\n", data)
	}
}

//ReturnOrderBook return the orderbook for a specific currencypair up to a certain depth
func (poloniexClient *PoloniexClient) ReturnOrderBook(currencypair string, depth int) (orderbook string, err error) {

	req, err := http.NewRequest("GET", "http://poloniex.com/public", nil)
	if err != nil {
		return
	}
	query := req.URL.Query()
	query.Set("command", "returnOrderBook")
	query.Set("currencyPair", currencypair)
	query.Set("depth", strconv.Itoa(depth))
	req.URL.RawQuery = query.Encode()

	logRequest(req)
	resp, err := poloniexClient.httpClient.Do(req)
	logResponse(resp, err)

	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	orderbook = string(body)
	return
}
