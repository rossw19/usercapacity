package model

import (
	"fmt"
	"rosswilson/usercapacity/utility"
	"testing"
	"time"
)

// Tests parsing of ics and days off
func TestVacationModel_buildModel(t *testing.T) {
	logger := &mockLogger{}
	utility.GetLogger().SetLogger(logger)

	clock := &mockClock{
		calendarDays: 7,
		workingDays:  5,
		averageOver:  3,
	}

	ics := `
BEGIN:VEVENT
UID:test@scheduleleave.com
DTSTAMP:%s
SUMMARY:Test User - Holiday (Approved)
DTSTART;TZID=UTC:%s
DTEND;TZID=UTC:%s
EMPLOYEEID:
END:VEVENT
	`

	data := fmt.Sprintf(ics, time.Now().Format("20060102T150405Z"), time.Now().Format("20060102T150405Z"), time.Now().AddDate(0, 0, 1).Format("20060102T150405Z"))

	previous := &mockModel{}
	previous.users = map[int]Userable{
		1: CreateUser("Test User", 0, 0, "test", 0, 0),
	}

	vacation := CreateVacationModel(previous, []byte(data), clock)
	vacation.buildModel()

	if len(vacation.GetUsers()) != 1 {
		t.Error("Expected 1 users")
	}

	if vacation.GetUsers()[1].GetDaysHaveOff() != 1 {
		t.Error("Expected 1 day off")
	}
}

func TestVacationModel_GetPrevious(t *testing.T) {
	previous := &mockModel{}
	vacation := CreateVacationModel(previous, nil, nil)
	if vacation.GetPrevious() != previous {
		t.Error("Expected previous")
	}
}

func TestVacationModel_GetUsers(t *testing.T) {
	user := CreateUser("test", 1, 1, "test", 1, 1)
	users := map[int]Userable{1: user}
	vacation := CreateVacationModel(nil, nil, nil)
	vacation.users = users

	if vacation.GetUsers()[1] != user {
		t.Error("Expected users")
	}
}

// Similar to building model, but testing reciever directly
func TestGetUserDaysOff(t *testing.T) {
	clock := &mockClock{
		calendarDays: 7,
		workingDays:  5,
		averageOver:  3,
	}

	ics := `
BEGIN:VEVENT
UID:test@scheduleleave.com
DTSTAMP:%s
SUMMARY:Test User - Holiday (Approved)
DTSTART;TZID=UTC:%s
DTEND;TZID=UTC:%s
EMPLOYEEID:
END:VEVENT
	`

	data := fmt.Sprintf(ics, time.Now().Format("20060102T150405Z"), time.Now().Format("20060102T150405Z"), time.Now().AddDate(0, 0, 1).Format("20060102T150405Z"))

	previous := &mockModel{}
	previous.users = map[int]Userable{
		1: CreateUser("Test User", 0, 0, "test", 0, 0),
	}

	vacation := CreateVacationModel(previous, []byte(data), clock)
	user := vacation.GetPrevious().GetUsers()[1]

	if vacation.getUserDaysOff(user, time.Now(), time.Now().AddDate(0, 0, 1)) != 1 {
		t.Error("Expected 1 day off")
	}
}

func TestSortTimes(t *testing.T) {
	timeBefore := time.Now().AddDate(-1, 0, 0)
	timeAfter := time.Now()

	timeBefore, timeAfter = sortTimes(timeBefore, timeAfter)

	if timeBefore.After(timeAfter) {
		t.Fail()
	}
}
