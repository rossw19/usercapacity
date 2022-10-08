package utility

import (
	"fmt"
)

type Clocker interface {
	GetCalendarDays() int
	GetWorkingDays() int
	GetAverageOver() int
}

type Clock struct {
	calendarDays int
	workingDays  int
	averageOver  int
}

func (c Clock) GetCalendarDays() int {
	return c.calendarDays
}

func (c Clock) GetWorkingDays() int {
	return c.workingDays
}

func (c Clock) GetAverageOver() int {
	return c.averageOver
}

func CreateClock() *Clock {
	return &Clock{
		calendarDays: GetConfig().Period.CalendarDays,
		workingDays:  GetConfig().Period.WorkingDays,
		averageOver:  GetConfig().Period.AverageOver,
	}
}

func GetFormattedTime(totalSeconds int) string {
	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60

	return fmt.Sprintf("%02dh %02dm %02ds", hours, minutes, seconds)
}
