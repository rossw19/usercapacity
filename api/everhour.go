package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"rosswilson/usercapacity/utility"
)

type EverhourStrategy struct {
	url          string
	authKey      string
	apiVersion   string
	requestUri   string
	request      *http.Request
	responseBody []byte
}

func CreateEverhourStrategy() *EverhourStrategy {
	config := utility.GetConfig().Env.Everhour

	return &EverhourStrategy{
		url:        config.Url,
		authKey:    config.Auth,
		apiVersion: config.Version,
	}
}

func (e *EverhourStrategy) execute() {
	e.processRequest()
	e.responseBody = e.processResponse()
}

func (e *EverhourStrategy) processRequest() {
	req, err := http.NewRequest("GET", e.url+e.requestUri, nil)
	if err != nil {
		utility.GetLogger().Write(fmt.Sprintf("api: bad http request %+v", req))
		os.Exit(1)
	}

	req.Header.Add("X-Api-Key", e.authKey)
	req.Header.Add("X-Accept-Version", e.apiVersion)
	e.request = req
}

func (e EverhourStrategy) processResponse() []byte {
	client := &http.Client{}
	resp, err := client.Do(e.request)
	if err != nil {
		utility.GetLogger().Write(fmt.Sprintf("api: bad http response from %s", e.requestUri))
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		utility.GetLogger().Write("api: could not read body of request")
		os.Exit(1)
	}

	return body
}

func (e *EverhourStrategy) SetRequestUri(requestUri string) {
	e.requestUri = requestUri
}

func (e EverhourStrategy) GetResponseBody() []byte {
	return e.responseBody
}
