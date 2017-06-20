package poloniexclient

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
)

const (
	poloniexPublicAPIUrl  = "http://poloniex.com/public"
	poloniexTradingAPIUrl = "https://poloniex.com/tradingApi"
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
		data, err := httputil.DumpRequestOut(request, true)
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

func (poloniexClient *PoloniexClient) executeTradingAPICommand(command string, parameters map[string]string, target interface{}) (err error) {
	log.Debug("Executing trading API command: ", command)

	form := url.Values{}
	form.Add("command", command)
	for key, value := range parameters {
		form.Add(key, value)
	}
	//TODO: really small chance of collision
	form.Add("nonce", strconv.FormatInt(time.Now().UnixNano(), 10))

	body := form.Encode()
	log.Debug("BODY:", body, "-ENDBODY")
	req, err := http.NewRequest("POST", poloniexTradingAPIUrl, strings.NewReader(body))
	if err != nil {
		return
	}
	req.Header.Add("Key", poloniexClient.Key)

	mac := hmac.New(sha512.New, []byte(poloniexClient.Secret))
	mac.Write([]byte(body))
	signature := hex.EncodeToString(mac.Sum(nil))

	req.Header.Add("Sign", signature)

	// with content-type, "error: invalid command" raise
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
