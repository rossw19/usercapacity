package model

type Modeler interface {
	buildModel()
	GetPrototype() Modeler
	GetUsers() map[int]user
}

type User interface {
	GetName() string
	GetTrackedTime() int
	GetAvgTime() float32
}

type user struct {
	name        string
	trackedTime int
	averageTime int
}

func (u user) GetName() string {
	return u.name
}

func (u user) GetTimeTracked() int {
	return u.trackedTime
}

func (u user) GetAvgTime() int {
	return u.averageTime
}
