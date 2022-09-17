package utility

import (
	"testing"
	"time"
)

func TestCreateDates(t *testing.T) {
	timestamp, _ := time.Parse("2006-01-02", "2022-09-17")
	dates := CreateDates(-21, timestamp)

	if dates.GetTo() != "2022-08-27" {
		t.Fail()
	}

	if dates.GetFrom() != "2022-09-17" {
		t.Fail()
	}
}
