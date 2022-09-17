package api

type apiContext struct {
	apiStrategy ApiStrategy
}

func CreateApiContext() *apiContext {
	return &apiContext{}
}

func (a *apiContext) SetApiStrategy(apiStrategy ApiStrategy) {
	a.apiStrategy = apiStrategy
}

func (a apiContext) ExecuteApi() {
	a.apiStrategy.execute()
}

type ApiStrategy interface {
	execute()
	createRequest()
	processResponse() []byte
}
