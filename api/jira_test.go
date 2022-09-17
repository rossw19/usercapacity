package api

import "testing"

func TestEncodeUsernameAndPassword(t *testing.T) {
	jira := jiraStrategy{
		username: "username",
		password: "password",
	}

	jira.encodeUsernamePassword()
	if jira.encoded != "dXNlcm5hbWU6cGFzc3dvcmQ=" {
		t.Fail()
	}
}
