package model

type Modeler interface {
	buildModel()
	GetPrevious() Modeler
	GetUsers() map[int]Userable
}

type Userable interface {
	GetName() string
	GetTimeTracked() int
	GetAvgTime() int
	GetJiraId() string
	GetDaysHadOff() int
	GetDaysHaveOff() int
}

type User struct {
	name        string
	trackedTime int
	averageTime int
	jiraId      string
	daysHadOff  int
	daysHaveOff int
}

func (u User) GetName() string {
	return u.name
}

func (u User) GetTimeTracked() int {
	return u.trackedTime
}

func (u User) GetAvgTime() int {
	return u.averageTime
}

func (u User) GetJiraId() string {
	return u.jiraId
}

func (u User) GetDaysHadOff() int {
	return u.daysHadOff
}

func (u User) GetDaysHaveOff() int {
	return u.daysHaveOff
}

func CreateUser(name string, trackedTime int, averageTime int, jiraId string, daysHadOff int, daysHaveOff int) *User {
	return &User{
		name:        name,
		trackedTime: trackedTime,
		averageTime: averageTime,
		jiraId:      jiraId,
		daysHadOff:  daysHadOff,
		daysHaveOff: daysHaveOff,
	}
}
