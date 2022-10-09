package model

import (
	"testing"
	"time"
)

func TestSortTimes(t *testing.T) {
	timeBefore := time.Now().AddDate(-1, 0, 0)
	timeAfter := time.Now()

	timeBefore, timeAfter = sortTimes(timeBefore, timeAfter)

	if timeBefore.After(timeAfter) {
		t.Fail()
	}
}
