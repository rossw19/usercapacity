package model

import (
	"rosswilson/usercapacity/utility"
)

type mathModel struct {
	users     map[int]user
	prototype Modeler
}

func (m *mathModel) buildModel() {
	m.users = map[int]user{}

	for i, t := range m.prototype.GetUsers() {
		m.users[i] = user{
			name:        t.GetName(),
			trackedTime: t.GetTimeTracked(),
			averageTime: t.GetTimeTracked() / 3,
		}
	}

	utility.GetLogger().Write("model: built mathModel")
}

func (m mathModel) GetPrototype() Modeler {
	return m.prototype
}

func (e mathModel) GetUsers() map[int]user {
	return e.users
}

func CreateMathModel(prototype Modeler) *mathModel {
	return &mathModel{
		prototype: prototype,
	}
}
