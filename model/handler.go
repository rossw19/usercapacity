package model

import "errors"

type ModelHandler interface {
	Handle() ModelHandler
	Add(model Modeler)
}

type handler struct {
	models []Modeler
}

func (h handler) Handle() *handler {
	for _, m := range h.models {
		m.buildModel()
	}

	return &h
}

func (h *handler) Add(model Modeler) {
	h.models = append(h.models, model)
}

func (h handler) GetLastModel() (Modeler, error) {
	if len(h.models) == 0 {
		return nil, errors.New("model: no models present")
	}

	return h.models[len(h.models)-1], nil
}

func CreateHandler(models []Modeler) *handler {
	return &handler{
		models: models,
	}
}
