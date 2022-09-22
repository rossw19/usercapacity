package model

import (
	"rosswilson/usercapacity/utility"
)

type mathModel struct {
	users    map[int]user
	previous Modeler
}

func (m *mathModel) buildModel() {
	m.users = map[int]user{}

	for i, t := range m.previous.GetUsers() {
		m.users[i] = user{
			name:        t.GetName(),
			trackedTime: t.GetTimeTracked(),
			averageTime: t.GetTimeTracked() / 3,
		}
	}

	utility.GetLogger().Write("model: built mathModel")
}

func (m mathModel) GetPrevious() Modeler {
	return m.previous
}

func (e mathModel) GetUsers() map[int]user {
	return e.users
}

func CreateMathModel(previous Modeler) *mathModel {
	return &mathModel{
		previous: previous,
	}
}
