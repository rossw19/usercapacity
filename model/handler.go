package model

type ModelHandler interface {
	Handle(model Model)
}

type everhourUserHandler struct {
	next ModelHandler
}

// Returns early so we can provide
// contextual information for next step
func (e everhourUserHandler) Handle(model Model) {
	if model.getType() == everhourUserType {
		model.buildModel()
		return
	}

	if e.next != nil {
		e.next.Handle(model)
	}
}

func CreateEverhourUserHandler(next ModelHandler) *everhourUserHandler {
	return &everhourUserHandler{
		next: next,
	}
}

type everhourTimeHandler struct {
	next ModelHandler
}

func (e everhourTimeHandler) Handle(model Model) {
	if model.getType() == everhourTimeType {
		model.buildModel()
	}

	if e.next != nil {
		e.next.Handle(model)
	}
}

func CreateEverhourTimeHandler(next ModelHandler) *everhourTimeHandler {
	return &everhourTimeHandler{
		next: next,
	}
}

type mathHandler struct {
	next ModelHandler
}

func (m mathHandler) Handle(model Model) {
	if model.getType() == mathType {
		model.buildModel()
	}

	if m.next != nil {
		m.next.Handle(model)
	}
}

func CreateMathHandler(next ModelHandler) *mathHandler {
	return &mathHandler{
		next: next,
	}
}

type jiraHandler struct {
	next ModelHandler
}

func (j jiraHandler) Handle(model Model) {
	if model.getType() == jiraType {
		model.buildModel()
	}

	if j.next != nil {
		j.next.Handle(model)
	}
}

func CreateJiraHandler(next ModelHandler) *jiraHandler {
	return &jiraHandler{
		next: next,
	}
}
