package api

type Context interface {
	SetApiStrategy(ApiStrategy)
	ExecuteApi()
}

type ApiContext struct {
	apiStrategy ApiStrategy
}

func CreateApiContext() *ApiContext {
	return &ApiContext{}
}

func (a *ApiContext) SetApiStrategy(apiStrategy ApiStrategy) {
	a.apiStrategy = apiStrategy
}

func (a ApiContext) ExecuteApi() {
	a.apiStrategy.execute()
}

type ApiStrategy interface {
	execute()
	processRequest()
	processResponse() []byte
}
