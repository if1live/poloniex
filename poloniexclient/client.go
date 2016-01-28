package poloniexclient

import (
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
