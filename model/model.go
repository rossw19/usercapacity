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
	GetDaysHadOff() int
	GetDaysHaveOff() int
}

type user struct {
	name        string
	trackedTime int
	averageTime int
	jiraId      string
	daysHadOff  int
	daysHaveOff int
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

func (u user) GetDaysHadOff() int {
	return u.daysHadOff
}

func (u user) GetDaysHaveOff() int {
	return u.daysHaveOff
}
