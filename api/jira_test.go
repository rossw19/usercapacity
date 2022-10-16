package api

import (
	"rosswilson/usercapacity/model"
	"rosswilson/usercapacity/utility"
	"testing"
)

func createMockJiraStrategy() *JiraStrategy {
	urlScope := utility.CreateScope("api_url_jira", "test")
	emailScope := utility.CreateScope("api_email_jira", "test")
	authScope := utility.CreateScope("api_auth_jira", "test")

	config := utility.CreateConfig()
	config.AddScope(*urlScope)
	config.AddScope(*emailScope)
	config.AddScope(*authScope)

	user := model.CreateUser("test", 1, 1, "test", 1, 1)

	utility.GetConfigProxy().SetConfig(config)
	return CreateJiraStrategy(user)
}

func TestCreateJiraStrategy(t *testing.T) {
	strategy := createMockJiraStrategy()

	if strategy == nil {
		t.Errorf("api: CreateJiraStrategy() returned nil")
	}

	if strategy.url != "test" {
		t.Errorf("api: CreateJiraStrategy() url was not set correctly")
	}

	if strategy.user == nil {
		t.Errorf("api: CreateJiraStrategy() user was not set correctly")
	}
}

func TestJiraStrategy_ProcessRequest(t *testing.T) {
	strategy := createMockJiraStrategy()
	strategy.processRequest()

	if strategy.request == nil {
		t.Errorf("api: processRequest() request was not set")
	}

	if strategy.request.Method != "PUT" {
		t.Errorf("api: processRequest() request method was not set correctly")
	}

	if strategy.request.URL.String() != "testtest" {
		t.Errorf("api: processRequest() request url was not set correctly")
	}
}

func TestJiraStrategy_SetRequestUri(t *testing.T) {
	strategy := createMockJiraStrategy()
	strategy.SetRequestUri("test")

	if strategy.requestUri != "test" {
		t.Errorf("api: SetRequestUri() requestUri was not set correctly")
	}
}

func TestEncodeUsernameAndPassword(t *testing.T) {
	jira := JiraStrategy{
		username: "username",
		password: "password",
	}

	jira.encodeUsernamePassword()
	if jira.encoded != "dXNlcm5hbWU6cGFzc3dvcmQ=" {
		t.Fail()
	}
}
