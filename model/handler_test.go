package model

import "testing"

func TestCreateHandler(t *testing.T) {
	handler := CreateHandler([]Modeler{})
	if handler == nil {
		t.Error("handler is nil")
	}
}

func TestHandle(t *testing.T) {
	mockModel := &mockModel{}
	handler := CreateHandler([]Modeler{mockModel})
	handler.Handle()

	if !mockModel.built {
		t.Error("model was not built")
	}
}

func TestAdd(t *testing.T) {
	mockModel := &mockModel{}
	handler := CreateHandler([]Modeler{})
	handler.Add(mockModel)

	if len(handler.models) != 1 {
		t.Error("model was not added")
	}
}

func TestGetLastModel(t *testing.T) {
	mockModel := &mockModel{}
	handler := CreateHandler([]Modeler{mockModel})
	lastModel, err := handler.GetLastModel()

	if err != nil {
		t.Error(err)
	}

	if lastModel != mockModel {
		t.Error("last model is not the mock model")
	}
}
