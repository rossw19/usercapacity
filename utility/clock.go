package utility

import (
	"fmt"
)

type Clocker interface {
	GetCalendarDays() int
	GetWorkingDays() int
	GetAverageOver() int
}

type clock struct {
	calendarDays int
	workingDays  int
	averageOver  int
}

func (c clock) GetCalendarDays() int {
	return c.calendarDays
}

func (c clock) GetWorkingDays() int {
	return c.workingDays
}

func (c clock) GetAverageOver() int {
	return c.averageOver
}

func CreateClock() *clock {
	return &clock{
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
