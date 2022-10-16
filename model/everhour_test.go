package model

import (
	"rosswilson/usercapacity/utility"
	"testing"
)

func TestCreateEverhourUserModel(t *testing.T) {
	user := CreateEverhourUserModel(nil, nil)
	if user == nil {
		t.Error("Expected non-nil user")
	}
}

// Tests adding users from data source
func TestEverhourUserModel_buildModel(t *testing.T) {
	logger := &mockLogger{}
	utility.GetLogger().SetLogger(logger)

	data := []byte(`[{"id":1,"name":"test"}]`)
	userModel := CreateEverhourUserModel(nil, data)
	userModel.buildModel()

	if len(userModel.GetUsers()) != 1 {
		t.Error("Expected 1 user")
	}

	if userModel.GetUsers()[1].GetName() != "test" {
		t.Error("Expected user")
	}
}

func TestEverhourUserModel_GetPrevious(t *testing.T) {
	previous := &mockModel{}
	user := CreateEverhourUserModel(previous, nil)
	if user.GetPrevious() != previous {
		t.Error("Expected nil previous")
	}
}

func TestEverhourUserModel_GetUsers(t *testing.T) {
	user := CreateUser("test", 1, 1, "test", 1, 1)
	users := map[int]Userable{1: user}
	userModel := CreateEverhourUserModel(nil, nil)
	userModel.users = users

	if userModel.GetUsers()[1] != user {
		t.Error("Expected users")
	}
}

func TestCreateEverhourTimeModel(t *testing.T) {
	time := CreateEverhourTimeModel(nil, nil)
	if time == nil {
		t.Error("Expected non-nil time")
	}
}

// Tests adding tracked time
func TestEverhourTimeModel_buildModel(t *testing.T) {
	logger := &mockLogger{}
	utility.GetLogger().SetLogger(logger)

	mockModel := &mockModel{}
	mockModel.users = map[int]Userable{1: CreateUser("test", 1, 1, "test", 1, 1)}

	data := []byte(`[{"memberId":1,"time":1}]`)
	timeModel := CreateEverhourTimeModel(mockModel, data)
	timeModel.buildModel()

	if len(timeModel.GetUsers()) != 1 {
		t.Error("Expected 1 user")
	}

	if timeModel.GetUsers()[1].GetTimeTracked() != 1 {
		t.Error("Expected tracked time")
	}
}

func TestEverhourTimeModel_GetPrevious(t *testing.T) {
	previous := &mockModel{}
	time := CreateEverhourTimeModel(previous, nil)
	if time.GetPrevious() != previous {
		t.Error("Expected previous")
	}
}

func TestEverhourTimeModel_GetUsers(t *testing.T) {
	user := CreateUser("test", 1, 1, "test", 1, 1)
	users := map[int]Userable{1: user}
	time := CreateEverhourTimeModel(nil, nil)
	time.users = users

	if time.GetUsers()[1] != user {
		t.Error("Expected users")
	}
}
