package api

type Contexter interface {
	SetApiStrategy(Strategizer)
	ExecuteApi()
}

type Context struct {
	apiStrategy Strategizer
}

func CreateApiContext() *Context {
	return &Context{}
}

func (a *Context) SetApiStrategy(apiStrategy Strategizer) {
	a.apiStrategy = apiStrategy
}

func (a Context) ExecuteApi() {
	a.apiStrategy.execute()
}

type Strategizer interface {
	execute()
	processRequest()
	processResponse() []byte
	SetRequestUri(string)
}
