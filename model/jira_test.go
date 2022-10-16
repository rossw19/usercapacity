package model

import (
	"rosswilson/usercapacity/utility"
	"testing"
)

// Tests application of jira id
func TestJiraModel_buildModel(t *testing.T) {
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
	}

	jira := CreateJiraModel(previous)
	jira.buildModel()

	if len(jira.GetUsers()) != 1 {
		t.Error("Expected 1 user")
	}

	if jira.GetUsers()[1].GetJiraId() != "test" {
		t.Error("Expected jira id")
	}
}

func TestJiraModel_GetPrevious(t *testing.T) {
	previous := &mockModel{}
	jira := CreateJiraModel(previous)
	if jira.GetPrevious() != previous {
		t.Error("Expected previous")
	}
}

func TestJiraModel_GetUsers(t *testing.T) {
	user := CreateUser("test", 1, 1, "test", 1, 1)
	users := map[int]Userable{1: user}
	jira := CreateJiraModel(nil)
	jira.users = users

	if jira.GetUsers()[1] != user {
		t.Error("Expected users")
	}
}

func TestCreateJiraModel(t *testing.T) {
	jira := CreateJiraModel(nil)
	if jira == nil {
		t.Error("Expected non-nil jira")
	}
}
