package api

import (
	"rosswilson/usercapacity/utility"
	"testing"
)

func createMockEverhourStrategy() *EverhourStrategy {
	urlScope := utility.CreateScope("api_url_everhour", "test")
	authScope := utility.CreateScope("api_auth_everhour", "test")
	versionScope := utility.CreateScope("api_version_everhour", "test")

	config := utility.CreateConfig()
	config.AddScope(*urlScope)
	config.AddScope(*authScope)
	config.AddScope(*versionScope)

	utility.GetConfigProxy().SetConfig(config)
	return CreateEverhourStrategy()
}

func TestCreateEverhourStrategy(t *testing.T) {
	strategy := createMockEverhourStrategy()

	if strategy == nil {
		t.Errorf("api: CreateEverhourStrategy() returned nil")
	}

	if strategy.url != "test" {
		t.Errorf("api: CreateEverhourStrategy() url was not set correctly")
	}

	if strategy.authKey != "test" {
		t.Errorf("api: CreateEverhourStrategy() authKey was not set correctly")
	}

	if strategy.apiVersion != "test" {
		t.Errorf("api: CreateEverhourStrategy() apiVersion was not set correctly")
	}
}

func TestEverhourStrategy_ProcessRequest(t *testing.T) {
	strategy := createMockEverhourStrategy()
	strategy.processRequest()

	if strategy.request == nil {
		t.Errorf("api: processRequest() request was not set")
	}

	if strategy.request.Method != "GET" {
		t.Errorf("api: processRequest() request method was not set correctly")
	}

	if strategy.request.URL.String() != "test" {
		t.Errorf("api: processRequest() request url was not set correctly")
	}

	if strategy.request.Header.Get("X-Api-Key") != "test" {
		t.Errorf("api: processRequest() request auth was not set correctly")
	}

	if strategy.request.Header.Get("X-Accept-Version") != "test" {
		t.Errorf("api: processRequest() request content type was not set correctly")
	}
}

func TestEverhourStrategy_SetRequestUri(t *testing.T) {
	strategy := createMockEverhourStrategy()
	strategy.SetRequestUri("test")

	if strategy.requestUri != "test" {
		t.Errorf("api: setRequestUri() request uri was not set correctly")
	}
}

func TestEverhourStrategy_TestGetResponseBody(t *testing.T) {
	strategy := createMockEverhourStrategy()
	strategy.responseBody = []byte("test")

	if string(strategy.GetResponseBody()) != "test" {
		t.Errorf("api: getResponseBody() response body was not set correctly")
	}
}
