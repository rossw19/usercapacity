package utility

import "testing"

func TestGetFormattedTime(t *testing.T) {
	formattedTime := GetFormattedTime(8193)

	if formattedTime != "02h 16m 33s" {
		t.Fail()
	}
}
