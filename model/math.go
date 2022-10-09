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

	for i, t := range m.GetPrevious().GetUsers() {
		averageTime := m.calculateAverageTime(t.GetTimeTracked(), t.GetDaysHadOff(), t.GetDaysHaveOff())

		m.users[i] = user{
			name:        t.GetName(),
			trackedTime: t.GetTimeTracked(),
			averageTime: averageTime,
			daysHadOff:  t.GetDaysHadOff(),
			daysHaveOff: t.GetDaysHaveOff(),
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

// Works out how much time a user
// should be working the following period
func (m *MathModel) calculateAverageTime(timeTracked int, daysHadOff int, daysHaveOff int) int {
	workingDaysOverPeriod := m.clock.GetAverageOver() * m.clock.GetWorkingDays()
	averageTimeWorkedPerDay := timeTracked / (workingDaysOverPeriod - daysHadOff)
	return averageTimeWorkedPerDay * (m.clock.GetWorkingDays() - daysHaveOff)
}
