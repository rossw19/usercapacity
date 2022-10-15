package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"rosswilson/usercapacity/utility"
)

type FileStrategy struct {
	url        string
	requestUri string
	requestUrl string
	file       []byte
}

func CreateFileStrategy() *FileStrategy {
	url, ok := utility.GetConfigProxy().GetScope("api_url_ics").ResolveString()
	if !ok {
		utility.GetLogger().Write(errors.New("api: could not resolve api_url_ics"))
		os.Exit(1)
	}

	return &FileStrategy{
		url: url,
	}
}

func (i *FileStrategy) execute() {
	i.processRequest()
	i.processResponse()
}

func (i *FileStrategy) processRequest() {
	i.requestUrl = i.url + i.requestUri
}

func (i *FileStrategy) processResponse() {
	resp, err := http.Get(i.requestUrl)
	if err != nil {
		utility.GetLogger().Write(fmt.Sprintf("api: bad http request %+v", err))
		os.Exit(1)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		utility.GetLogger().Write(fmt.Sprintf("api: could not read body %+v", err))
		os.Exit(1)
	}

	i.file = body
}

func (i *FileStrategy) SetRequestUri(uri string) {
	i.requestUri = uri
}

func (i *FileStrategy) GetFileContents() []byte {
	return i.file
}
