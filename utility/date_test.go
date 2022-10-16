package utility

import (
	"testing"
	"time"
)

func TestGetFrom(t *testing.T) {
	date := Dates{
		from: "2022-09-17",
	}

	if date.GetFrom() != "2022-09-17" {
		t.Errorf("Expected 2022-09-17, got %s", date.GetFrom())
	}
}

func TestGetTo(t *testing.T) {
	date := Dates{
		to: "2022-09-17",
	}

	if date.GetTo() != "2022-09-17" {
		t.Errorf("Expected 2022-09-17, got %s", date.GetTo())
	}
}

func TestCreateDates(t *testing.T) {
	timestamp, _ := time.Parse("2006-01-02", "2022-09-17")
	dates := CreateDates(-21, timestamp)

	if dates.GetTo() != "2022-08-27" {
		t.Errorf("Expected 2022-08-27, got %s", dates.GetTo())
	}

	if dates.GetFrom() != "2022-09-17" {
		t.Errorf("Expected 2022-09-17, got %s", dates.GetFrom())
	}
}
