package api

import (
	"encoding/base64"
	"fmt"
	"rosswilson/usercapacity/utility"
)

type jiraStrategy struct {
	url      string
	username string
	password string
	encoded  string
}

func CreateJiraStrategy() *jiraStrategy {
	return &jiraStrategy{
		url:      utility.GetEnvOrExit("JIRA_URL"),
		username: utility.GetEnvOrExit("JIRA_USERNAME"),
		password: utility.GetEnvOrExit("JIRA_PASSWORD"),
	}
}

func (j jiraStrategy) execute() {
	j.encodeUsernamePassword()
}

func (j *jiraStrategy) encodeUsernamePassword() {
	usernamePassword := fmt.Sprintf("%s:%s", j.username, j.password)
	j.encoded = base64.StdEncoding.EncodeToString([]byte(usernamePassword))
}

func (j *jiraStrategy) createRequest() {

}

func (j jiraStrategy) processResponse() []byte {
	return make([]byte, 0)
}
