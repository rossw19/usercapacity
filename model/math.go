package model

import (
	"rosswilson/usercapacity/utility"
)

type mathModel struct {
	users    map[int]user
	previous Modeler
	clock    utility.Clocker
}

func (m *mathModel) buildModel() {
	m.users = map[int]user{}

	for i, t := range m.previous.GetUsers() {
		m.users[i] = user{
			name:        t.GetName(),
			trackedTime: t.GetTimeTracked(),
			averageTime: t.GetTimeTracked() / m.clock.GetAverageOver(),
		}
	}

	utility.GetLogger().Write("model: built mathModel")
}

func (m mathModel) GetPrevious() Modeler {
	return m.previous
}

func (m mathModel) GetUsers() map[int]user {
	return m.users
}

func CreateMathModel(previous Modeler, clock utility.Clocker) *mathModel {
	return &mathModel{
		previous: previous,
		clock:    clock,
	}
}
