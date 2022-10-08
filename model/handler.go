package model

import "errors"

type ModelHandler interface {
	Handle() ModelHandler
	Add(model Modeler)
}

type Handler struct {
	models []Modeler
}

func (h *Handler) Handle() *Handler {
	for _, m := range h.models {
		m.buildModel()
	}

	return h
}

func (h *Handler) Add(model Modeler) {
	h.models = append(h.models, model)
}

func (h *Handler) GetLastModel() (Modeler, error) {
	if len(h.models) == 0 {
		return nil, errors.New("model: no models present")
	}

	return h.models[len(h.models)-1], nil
}

func CreateHandler(models []Modeler) *Handler {
	return &Handler{
		models: models,
	}
}
