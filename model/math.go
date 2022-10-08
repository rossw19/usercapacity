package model

import (
	"rosswilson/usercapacity/utility"
)

type MathModel struct {
	users    map[int]User
	previous Modeler
	clock    utility.Clocker
}

func (m *MathModel) buildModel() {
	m.users = map[int]User{}

	for i, t := range m.previous.GetUsers() {
		m.users[i] = user{
			name:        t.GetName(),
			trackedTime: t.GetTimeTracked(),
			averageTime: t.GetTimeTracked() / m.clock.GetAverageOver(),
		}
	}

	utility.GetLogger().Write("model: built MathModel")
}

func (m *MathModel) GetPrevious() Modeler {
	return m.previous
}

func (m *MathModel) GetUsers() map[int]User {
	return m.users
}

func CreateMathModel(previous Modeler, clock utility.Clocker) *MathModel {
	return &MathModel{
		previous: previous,
		clock:    clock,
	}
}
