package api

import (
	"rosswilson/usercapacity/utility"
	"testing"
)

func createMockFileStrategy() *FileStrategy {
	urlScope := utility.CreateScope("api_url_ics", "test")

	config := utility.CreateConfig()
	config.AddScope(*urlScope)

	utility.GetConfigProxy().SetConfig(config)
	return CreateFileStrategy()
}

func TestCreateFileStrategy(t *testing.T) {
	strategy := createMockFileStrategy()

	if strategy == nil {
		t.Errorf("api: CreateFileStrategy() returned nil")
	}

	if strategy.url != "test" {
		t.Errorf("api: CreateFileStrategy() url was not set correctly")
	}
}

func TestFileStrategy_ProcessRequest(t *testing.T) {
	strategy := createMockFileStrategy()
	strategy.SetRequestUri("test")
	strategy.processRequest()

	if strategy.requestUrl != "testtest" {
		t.Errorf("api: processRequest() request url was not set correctly")
	}
}

func TestFileStrategy_SetRequestUri(t *testing.T) {
	strategy := createMockFileStrategy()
	strategy.SetRequestUri("test")

	if strategy.requestUri != "test" {
		t.Errorf("api: SetRequestUri() request uri was not set correctly")
	}
}

func TestFileStrategy_GetFileContents(t *testing.T) {
	strategy := createMockFileStrategy()
	strategy.file = []byte("test")

	if string(strategy.GetFileContents()) != "test" {
		t.Errorf("api: GetFileContents() file contents were not set correctly")
	}
}