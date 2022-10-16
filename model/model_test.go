package model

import (
	"rosswilson/usercapacity/utility"
	"testing"
)

type mockModel struct {
	built    bool
	users    map[int]Userable
	previous Modeler
}

func (m *mockModel) buildModel() {
	m.built = true
}

func (m *mockModel) GetPrevious() Modeler {
	return m.previous
}

func (m *mockModel) GetUsers() map[int]Userable {
	return m.users
}

type mockLogger struct {
	active bool
}

func (m *mockLogger) SetFile(filename string) utility.Loggable {
	return m
}

func (m *mockLogger) Write(message any) {
}

func (m *mockLogger) SetActive(active bool) utility.Loggable {
	return m
}

type mockClock struct {
	calendarDays int
	workingDays  int
	averageOver  int
}

func (m mockClock) GetCalendarDays() int {
	return m.calendarDays
}

func (m mockClock) GetWorkingDays() int {
	return m.workingDays
}

func (m mockClock) GetAverageOver() int {
	return m.averageOver
}

func TestCreateUser(t *testing.T) {
	user := CreateUser("test", 0, 0, "test", 0, 0)

	if user == nil {
		t.Error("User is nil")
	}
}

func TestGetName(t *testing.T) {
	user := CreateUser("test", 0, 0, "test", 0, 0)

	if user.GetName() != "test" {
		t.Error("Could not get name")
	}
}

func TestGetTimeTracked(t *testing.T) {
	user := CreateUser("test", 1, 0, "test", 0, 0)

	if user.GetTimeTracked() != 1 {
		t.Error("Could not get time tracked")
	}
}

func TestGetAvgTime(t *testing.T) {
	user := CreateUser("test", 0, 1, "test", 0, 0)

	if user.GetAvgTime() != 1 {
		t.Error("Could not get average time")
	}
}

func TestGetJiraId(t *testing.T) {
	user := CreateUser("test", 0, 0, "test", 0, 0)

	if user.GetJiraId() != "test" {
		t.Error("Could not get jira id")
	}
}

func TestGetDaysHadOff(t *testing.T) {
	user := CreateUser("test", 0, 0, "test", 1, 0)

	if user.GetDaysHadOff() != 1 {
		t.Error("Could not get days had off")
	}
}

func TestGetDaysHaveOff(t *testing.T) {
	user := CreateUser("test", 0, 0, "test", 0, 1)

	if user.GetDaysHaveOff() != 1 {
		t.Error("Could not get days have off")
	}
}
