package model

import "errors"

type ModelHandler interface {
	Handle() ModelHandler
	Add(model Model)
}

type handler struct {
	models []Model
}

func (h handler) Handle() *handler {
	for _, m := range h.models {
		m.buildModel()
	}

	return &h
}

func (h *handler) Add(model Model) {
	h.models = append(h.models, model)
}

func (h handler) GetLastModel() (Model, error) {
	if len(h.models) == 0 {
		return nil, errors.New("model: no models present")
	}

	return h.models[len(h.models)-1], nil
}

func CreateHandler(models []Model) *handler {
	return &handler{
		models: models,
	}
}
