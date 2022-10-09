package model

import (
	"bytes"
	"fmt"
	"github.com/apognu/gocal"
	"os"
	"rosswilson/usercapacity/utility"
	"strings"
	"time"
)

type VacationModel struct {
	users    map[int]User
	previous Modeler
	ics      []byte
	clock    utility.Clocker
}

func (v *VacationModel) buildModel() {
	v.users = map[int]User{}

	offsetDaysBefore := -1 * v.clock.GetCalendarDays() * v.clock.GetAverageOver()
	offsetDaysAfter := v.clock.GetCalendarDays()

	for i, t := range v.GetPrevious().GetUsers() {
		daysHadOff := v.getUserDaysOff(t, time.Now().AddDate(0, 0, offsetDaysBefore), time.Now())
		daysHaveOff := v.getUserDaysOff(t, time.Now(), time.Now().AddDate(0, 0, offsetDaysAfter))

		v.users[i] = user{
			name:        t.GetName(),
			trackedTime: t.GetTimeTracked(),
			daysHadOff:  daysHadOff,
			daysHaveOff: daysHaveOff,
		}
	}

	utility.GetLogger().Write("model: built VacationModel")
}

func (v *VacationModel) GetPrevious() Modeler {
	return v.previous
}

func (v *VacationModel) GetUsers() map[int]User {
	return v.users
}

func CreateVacationModel(previous Modeler, ics []byte, clock utility.Clocker) *VacationModel {
	return &VacationModel{
		previous: previous,
		ics:      ics,
		clock:    clock,
	}
}

func (v *VacationModel) getUserDaysOff(user User, start time.Time, end time.Time) int {
	reader := bytes.NewReader(v.ics)
	t0, t1 := sortTimes(start, end)

	cal := gocal.NewParser(reader)
	cal.Start, cal.End = &t0, &t1

	err := cal.Parse()
	if err != nil {
		utility.GetLogger().Write(fmt.Errorf("model: error parsing ics file: %s", err))
		os.Exit(1)
	}

	occurances := 0
	for _, e := range cal.Events {
		if strings.Contains(e.Summary, user.GetName()) {
			occurances++
		}
	}

	return occurances
}

func sortTimes(t0 time.Time, t1 time.Time) (time.Time, time.Time) {
	if t0.Before(t1) {
		return t0, t1
	}

	return t1, t0
}
