package api

import (
	"encoding/base64"
	"fmt"
	"rosswilson/usercapacity/utility"
)

type JiraStrategy struct {
	url        string
	username   string
	password   string
	encoded    string
	requestUri string
}

func CreateJiraStrategy() *JiraStrategy {
	return &JiraStrategy{
		url:      utility.GetEnvOrExit("JIRA_URL"),
		username: utility.GetEnvOrExit("JIRA_USERNAME"),
		password: utility.GetEnvOrExit("JIRA_PASSWORD"),
	}
}

func (j JiraStrategy) execute() {
	j.encodeUsernamePassword()
}

func (j *JiraStrategy) encodeUsernamePassword() {
	usernamePassword := fmt.Sprintf("%s:%s", j.username, j.password)
	j.encoded = base64.StdEncoding.EncodeToString([]byte(usernamePassword))
}

func (j *JiraStrategy) processRequest() {

}

func (j JiraStrategy) processResponse() []byte {
	return make([]byte, 0)
}

func (e *JiraStrategy) SetRequestUri(requestUri string) {
	e.requestUri = requestUri
}
