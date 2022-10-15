package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"rosswilson/usercapacity/utility"
)

const (
	defaultUrl     = "https://api.everhour.com"
	defaultVersion = "1.2"
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
	config := utility.GetConfigProxy()

	url, ok := config.GetScope("api_url_everhour").ResolveString()
	if !ok {
		url = defaultUrl
		utility.GetLogger().Write(fmt.Errorf("api: could not resolve api_url_everhour, using %s", url))
	}

	authKey, ok := config.GetScope("api_auth_everhour").ResolveString()
	if !ok {
		utility.GetLogger().Write(errors.New("api: could not resolve api_auth_everhour"))
		os.Exit(1)
	}

	apiVersion, ok := config.GetScope("api_version_everhour").ResolveString()
	if !ok {
		apiVersion = defaultVersion
		utility.GetLogger().Write(fmt.Errorf("api: could not resolve api_version_everhour, using version %s", apiVersion))
	}

	return &EverhourStrategy{
		url:        url,
		authKey:    authKey,
		apiVersion: apiVersion,
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

func (e *EverhourStrategy) processResponse() []byte {
	client := &http.Client{}
	resp, err := client.Do(e.request)
	if err != nil {
		utility.GetLogger().Write(fmt.Sprintf("api: bad http response from %s", e.requestUri))
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		utility.GetLogger().Write("api: could not read body of response")
		os.Exit(1)
	}

	return body
}

func (e *EverhourStrategy) SetRequestUri(requestUri string) {
	e.requestUri = requestUri
}

func (e *EverhourStrategy) GetResponseBody() []byte {
	return e.responseBody
}
