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
