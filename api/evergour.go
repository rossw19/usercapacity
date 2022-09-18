package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"rosswilson/usercapacity/utility"
)

type everhourStrategy struct {
	url        string
	authKey    string
	apiVersion string
	requestUri string
	logger     *utility.Logger
	request    *http.Request
}

func CreateEverhourStrategy() *everhourStrategy {
	return &everhourStrategy{
		url:        utility.GetEnvOrExit("EVERHOUR_URL"),
		authKey:    utility.GetEnvOrExit("EVERHOUR_AUTH_KEY"),
		apiVersion: utility.GetEnvOrExit("EVERHOUR_API_VERSION"),
		logger:     utility.GetLogger(),
	}
}

func (e everhourStrategy) execute() {
	e.processRequest()
	e.processResponse()
}

func (e *everhourStrategy) processRequest() {
	req, err := http.NewRequest("GET", e.url+e.requestUri, nil)
	if err != nil {
		e.logger.Write(fmt.Sprintf("api: bad http request %+v", req))
		os.Exit(1)
	}

	req.Header.Add("X-Api-Key", e.authKey)
	req.Header.Add("X-Accept-Version", e.apiVersion)
	e.request = req
}

func (e everhourStrategy) processResponse() []byte {
	client := &http.Client{}
	resp, err := client.Do(e.request)
	if err != nil {
		e.logger.Write(fmt.Sprintf("api: bad http response from %s", e.requestUri))
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		e.logger.Write("api: could not read body of request")
		os.Exit(1)
	}

	return body
}

func (e *everhourStrategy) SetRequestUri(requestUri string) {
	e.requestUri = requestUri
}
