package api

import (
	"testing"
)

type mockStrategy struct {
	requestUri string
	executed   bool
}

func (m *mockStrategy) execute() {
	m.executed = true
}

func (m *mockStrategy) SetRequestUri(uri string) {
	m.requestUri = uri
}

func createMockStrategy() *mockStrategy {
	return &mockStrategy{}
}

func TestCreateApiContext(t *testing.T) {
	apiContext := CreateApiContext()
	if apiContext == nil {
		t.Errorf("apiContext is nil")
	}
}

func TestCreateApiStrategy(t *testing.T) {
	context := CreateApiContext()
	strategy := createMockStrategy()
	context.SetApiStrategy(strategy)

	if context.apiStrategy == nil {
		t.Errorf("apiStrategy is nil")
	}
}

func TestExecuteApi(t *testing.T) {
	context := CreateApiContext()
	strategy := createMockStrategy()
	context.SetApiStrategy(strategy)

	context.ExecuteApi()

	if !strategy.executed {
		t.Errorf("strategy was not executed")
	}
}
