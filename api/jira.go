package api

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"rosswilson/usercapacity/model"
	"rosswilson/usercapacity/utility"
)

type JiraStrategy struct {
	url          string
	username     string
	password     string
	encoded      string
	requestUri   string
	user         model.User
	request      *http.Request
	responseBody []byte
}

func CreateJiraStrategy(user model.User) *JiraStrategy {
	config := utility.GetConfig().Env.Jira

	return &JiraStrategy{
		url:      config.Url,
		username: config.Email,
		password: config.Auth,
		user:     user,
	}
}

func (j *JiraStrategy) execute() {
	j.encodeUsernamePassword()
	j.processRequest()
	j.responseBody = j.processResponse()
}

func (j *JiraStrategy) encodeUsernamePassword() {
	usernamePassword := fmt.Sprintf("%s:%s", j.username, j.password)
	j.encoded = base64.StdEncoding.EncodeToString([]byte(usernamePassword))
}

func (j *JiraStrategy) processRequest() {
	data := []byte(fmt.Sprintf("%v", j.user.GetAvgTime()))
	body := bytes.NewReader(data)

	req, err := http.NewRequest("PUT", j.url+j.requestUri+j.user.GetJiraId(), body)
	if err != nil {
		utility.GetLogger().Write(fmt.Sprintf("api: bad http request %+v", req))
		os.Exit(1)
	}

	auth := fmt.Sprintf("Basic %s", j.encoded)
	req.Header.Add("Authorization", auth)
	j.request = req
}

func (j *JiraStrategy) processResponse() []byte {
	client := &http.Client{}
	resp, err := client.Do(j.request)
	if err != nil {
		utility.GetLogger().Write(fmt.Sprintf("api: bad http response from %s", j.requestUri))
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

func (j *JiraStrategy) SetRequestUri(requestUri string) {
	j.requestUri = requestUri
}
