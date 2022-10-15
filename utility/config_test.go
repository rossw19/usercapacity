package utility

import "testing"

func TestResolveInt(t *testing.T) {
	scope := CreateScope("test", "1")
	num, ok := scope.ResolveInt()

	if !ok {
		t.Error("Failed to resolve int")
	}

	if num != 1 {
		t.Error("Failed to resolve int")
	}
}

func TestResolveBoolean(t *testing.T) {
	scope := CreateScope("test", "true")
	boolean, ok := scope.ResolveBoolean()

	if !ok {
		t.Error("Failed to resolve boolean")
	}

	if !boolean {
		t.Error("Failed to resolve boolean")
	}
}

func TestResolveString(t *testing.T) {
	scope := CreateScope("test", "test")
	str, ok := scope.ResolveString()

	if !ok {
		t.Error("Failed to resolve string")
	}

	if str != "test" {
		t.Error("Failed to resolve string")
	}
}

func TestGetPath(t *testing.T) {
	scope := CreateScope("test", "test")
	path := scope.GetPath()

	if path != "test" {
		t.Error("Failed to get path")
	}
}

func TestCreateScope(t *testing.T) {
	scope := CreateScope("test", "test")

	if scope == nil {
		t.Error("Failed to create scope")
	}
}

func TestGetScope(t *testing.T) {
	scope := CreateScope("test", "test")
	config := CreateConfig()
	config.AddScope(*scope)

	if config.GetScope("test") == (Scope{}) {
		t.Error("Failed to get scope")
	}

	if config.GetScope("test").GetPath() != "test" {
		t.Error("Failed to get scope")
	}
}

func TestAddScope(t *testing.T) {
	scope := CreateScope("test", "test")
	config := CreateConfig()
	config.AddScope(*scope)

	if config.GetScope("test") == (Scope{}) {
		t.Error("Failed to add scope")
	}

	if config.GetScope("test").GetPath() != "test" {
		t.Error("Failed to add scope")
	}
}

func TestCreateUser(t *testing.T) {
	user := CreateUser(0, "test", "test")

	if user == nil {
		t.Error("Failed to create user")
	}
}

func TestAddUser(t *testing.T) {
	user := CreateUser(0, "test", "test")
	config := CreateConfig()
	config.AddUser(*user)
	users := config.GetUsers()

	if len(users) != 1 {
		t.Error("Failed to add user")
	}

	if users[0].GetName() != "test" {
		t.Error("Failed to add user")
	}

	if users[0].GetJiraId() != "test" {
		t.Error("Failed to add user")
	}

	if users[0].GetEverhourId() != 0 {
		t.Error("Failed to add user")
	}
}

func TestSetConfig(t *testing.T) {
	proxy := GetConfigProxy()
	config := CreateConfig()
	user := CreateUser(0, "test", "test")
	config.AddUser(*user)
	proxy.SetConfig(config)

	if len(proxy.GetUsers()) != 1 {
		t.Error("Failed to set config")
	}
}

func TestGetConfigProxy(t *testing.T) {
	proxy0 := GetConfigProxy()
	proxy1 := GetConfigProxy()

	if proxy0 == nil {
		t.Error("Failed to get config proxy")
	}

	if proxy1 == nil {
		t.Error("Failed to get config proxy")
	}

	if proxy0 != proxy1 {
		t.Error("Failed to get config proxy")
	}
}

func TestGetEverhourId(t *testing.T) {
	user := CreateUser(0, "test", "test")
	id := user.GetEverhourId()

	if id != 0 {
		t.Error("Failed to get everhour id")
	}
}

func TestGetName(t *testing.T) {
	user := CreateUser(0, "test", "test")
	name := user.GetName()

	if name != "test" {
		t.Error("Failed to get name")
	}
}

func TestGetJiraId(t *testing.T) {
	user := CreateUser(0, "test", "test")
	id := user.GetJiraId()

	if id != "test" {
		t.Error("Failed to get jira id")
	}
}

// Uses library which is quite volatile to changes
// in structs, this is why this test exists
func TestUnmarshalConfig(t *testing.T) {
	config := CreateConfig()
	yaml := `
scopes:
  - path: "test"
    value: "test"

users:
  - everhourId: 0
    jiraId: "test"
    name: "test"
`

	err := config.UnmarshalConfig([]byte(yaml))
	if err != nil {
		t.Error(err)
	}

	if config.GetScope("test").GetPath() != "test" {
		t.Error("Failed to unmarshal config")
	}

	if len(config.GetUsers()) != 1 {
		t.Error("Failed to unmarshal config")
	}
}
