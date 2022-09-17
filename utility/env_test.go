package utility

import "testing"

func TestEnvToBool(t *testing.T) {
	truthy := StringToBool("true")
	falsey := StringToBool("qwerty")

	if truthy != true {
		t.Fail()
	}

	if falsey == true {
		t.Fail()
	}
}
