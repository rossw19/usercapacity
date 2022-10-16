package model

import (
	"rosswilson/usercapacity/utility"
	"testing"
)

func TestMathModel_buildModel(t *testing.T) {
	logger := &mockLogger{}
	utility.GetLogger().SetLogger(logger)

	clock := &mockClock{
		calendarDays: 7,
		workingDays:  5,
		averageOver:  3,
	}

	previous := &mockModel{}
	previous.users = map[int]Userable{
		0: CreateUser("test", 259200, 0, "test", 0, 0),
		1: CreateUser("test", 259200, 0, "test", 0, 1),
	}

	math := CreateMathModel(previous, clock)
	math.buildModel()

	if len(math.GetUsers()) != 2 {
		t.Error("Expected 2 users")
	}

	if math.GetUsers()[0].GetAvgTime() != 86400 {
		t.Error("Expected 86400 average time")
	}

	if math.GetUsers()[1].GetAvgTime() != 69120 {
		t.Error("Expected 86400 average time")
	}
}

func TestMathModel_GetPrevious(t *testing.T) {
	previous := &mockModel{}
	math := CreateMathModel(previous, nil)
	if math.GetPrevious() != previous {
		t.Error("Expected previous")
	}
}

func TestMathModel_GetUsers(t *testing.T) {
	user := CreateUser("test", 1, 1, "test", 1, 1)
	users := map[int]Userable{1: user}
	math := CreateMathModel(nil, nil)
	math.users = users

	if math.GetUsers()[1] != user {
		t.Error("Expected users")
	}
}

func TestMathModel_calculateAverageTime(t *testing.T) {
	logger := &mockLogger{}
	utility.GetLogger().SetLogger(logger)

	clock := &mockClock{
		calendarDays: 7,
		workingDays:  5,
		averageOver:  3,
	}

	previous := &mockModel{}
	previous.users = map[int]Userable{
		0: CreateUser("test", 259200, 0, "test", 0, 0),
	}

	math := CreateMathModel(previous, clock)
	avg := math.calculateAverageTime(259200, 0, 1)

	if avg != 69120 {
		t.Error("Expected 69120 average time")
	}
}

func TestCreateMathModel(t *testing.T) {
	math := CreateMathModel(nil, nil)
	if math == nil {
		t.Error("Expected non-nil math")
	}
}
