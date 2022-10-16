package model

import (
	"rosswilson/usercapacity/utility"
	"testing"
)

// Tests filtering from two to one user
func TestFilterModel_buildModel(t *testing.T) {
	logger := &mockLogger{}
	utility.GetLogger().SetLogger(logger)

	configUser := utility.CreateUser(1, "test", "test")
	config := utility.CreateConfig()
	config.AddUser(configUser)
	proxy := utility.GetConfigProxy()
	proxy.SetConfig(config)

	previous := &mockModel{}
	previous.users = map[int]Userable{
		1: CreateUser("test", 1, 1, "test", 1, 1),
		2: CreateUser("test", 1, 1, "test", 1, 1),
	}

	filter := CreateFilterModel(previous)
	filter.buildModel()

	if len(filter.GetUsers()) != 1 {
		t.Error("Expected 1 user")
	}
}

func TestFilterModel_GetPrevious(t *testing.T) {
	previous := &mockModel{}
	filter := CreateFilterModel(previous)
	if filter.GetPrevious() != previous {
		t.Error("Expected previous")
	}
}

func TestFilterModel_GetUsers(t *testing.T) {
	user := CreateUser("test", 1, 1, "test", 1, 1)
	users := map[int]Userable{1: user}
	filter := CreateFilterModel(nil)
	filter.users = users

	if filter.GetUsers()[1] != user {
		t.Error("Expected users")
	}
}

func TestCreateFilterModel(t *testing.T) {
	filter := CreateFilterModel(nil)
	if filter == nil {
		t.Error("Expected non-nil filter")
	}
}
