package poloniexclient

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"

	log "github.com/Sirupsen/logrus"
)

const poloniexPublicAPIUrl = "http://poloniex.com/public"

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
		data, err := httputil.DumpRequest(request, true)
		if err != nil {
			log.Error("Error dumping request: ", err)
			return
		}
		log.Debugf("%s\n\n", data)
	}
}

func logResponse(response *http.Response, err error) {
	if err != nil {
		log.Error("HTTP reponse error: ", err)
		return
	}
	if log.GetLevel() == log.DebugLevel {
		data, err := httputil.DumpResponse(response, true)
		if err != nil {
			log.Error("Error dumping response: ", err)
			return
		}
		log.Debugf("%s\n\n", data)
	}
}

func (poloniexClient *PoloniexClient) executePublicAPICommand(command string, parameters map[string]string, target interface{}) (err error) {
	log.Debug("Executing public API command: ", command)
	req, err := http.NewRequest("GET", poloniexPublicAPIUrl, nil)
	if err != nil {
		return
	}
	query := req.URL.Query()
	query.Set("command", command)
	for key, value := range parameters {
		query.Set(key, value)
	}

	req.URL.RawQuery = query.Encode()

	logRequest(req)
	resp, err := poloniexClient.httpClient.Do(req)
	logResponse(resp, err)

	if err != nil {
		return
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(target)
	return
}
