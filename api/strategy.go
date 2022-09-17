package api

import (
	"encoding/base64"
	"fmt"
	"rosswilson/usercapacity/utility"
)

type apiContext struct {
	apiStrategy ApiStrategy
}

func CreateApiContext() *apiContext {
	return &apiContext{}
}

func (a *apiContext) SetApiStrategy(apiStrategy ApiStrategy) {
	a.apiStrategy = apiStrategy
}

func (a *apiContext) ExecuteApi() {
	a.apiStrategy.execute()
}

type ApiStrategy interface {
	execute()
}

type everhourStrategy struct {
	url        string
	authKey    string
	apiVersion string
}

func CreateEverhourStrategy() *everhourStrategy {
	return &everhourStrategy{
		url:        utility.GetEnvOrPanic("EVERHOUR_URL"),
		authKey:    utility.GetEnvOrPanic("EVERHOUR_AUTH_KEY"),
		apiVersion: utility.GetEnvOrPanic("EVERHOUR_API_VERSION"),
	}
}

func (e *everhourStrategy) execute() {
}

type jiraStrategy struct {
	url      string
	username string
	password string
	encoded  string
}

func CreateJiraStrategy() *jiraStrategy {
	return &jiraStrategy{
		url:      utility.GetEnvOrPanic("JIRA_URL"),
		username: utility.GetEnvOrPanic("JIRA_USERNAME"),
		password: utility.GetEnvOrPanic("JIRA_PASSWORD"),
	}
}

func (j *jiraStrategy) execute() {
	j.encodeUsernamePassword()
}

func (j *jiraStrategy) encodeUsernamePassword() {
	usernamePassword := fmt.Sprintf("%s:%s", j.username, j.password)
	j.encoded = base64.StdEncoding.EncodeToString([]byte(usernamePassword))
}
