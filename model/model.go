package model

type Modeler interface {
	buildModel()
	GetPrevious() Modeler
	GetUsers() map[int]User
}

type User interface {
	GetName() string
	GetTimeTracked() int
	GetAvgTime() int
	GetJiraId() string
}

type user struct {
	name        string
	trackedTime int
	averageTime int
	jiraId      string
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

func (u user) GetJiraId() string {
	return u.jiraId
}
