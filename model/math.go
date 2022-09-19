package model

import "rosswilson/usercapacity/utility"

type mathModel struct {
	times     map[int]users
	prototype *everhourTimeModel
}

type users struct {
	name        string
	trackedTime int
	averageTime float32
}

func (m *mathModel) buildModel() {
	m.times = map[int]users{}

	for i, t := range m.prototype.times {
		m.times[i] = users{
			name:        t.name,
			trackedTime: t.trackedTime,
			averageTime: float32(t.trackedTime) / 10800,
		}
	}

	utility.GetLogger().Write("model: built mathModel")
}

func (m *mathModel) GetPrototype() Model {
	return m.prototype
}

func CreateMathModel(prototype *everhourTimeModel) *mathModel {
	return &mathModel{
		prototype: prototype,
	}
}
