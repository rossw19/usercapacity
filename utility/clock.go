package utility

import (
	"fmt"
)

const (
	defaultCalendarDays = 7
	defaultWorkingDays  = 5
	defaultAverageOver  = 3
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
	config := GetConfigProxy()

	calendarDays, ok := config.GetScope("application_context_calendar_days").ResolveInt()
	if !ok {
		calendarDays = defaultCalendarDays
		GetLogger().Write(fmt.Errorf("api: could not resolve application_context_calendar_days, using %d", calendarDays))
	}

	workingDays, ok := config.GetScope("application_context_working_days").ResolveInt()
	if !ok {
		workingDays = defaultWorkingDays
		GetLogger().Write(fmt.Errorf("api: could not resolve application_context_working_days, using %d", workingDays))
	}

	averageOver, ok := config.GetScope("application_context_average_over").ResolveInt()
	if !ok {
		averageOver = defaultAverageOver
		GetLogger().Write(fmt.Errorf("api: could not resolve application_context_average_over, using %d", averageOver))
	}

	return &Clock{
		calendarDays: calendarDays,
		workingDays:  workingDays,
		averageOver:  averageOver,
	}
}

func GetFormattedTime(totalSeconds int) string {
	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60

	return fmt.Sprintf("%02dh %02dm %02ds", hours, minutes, seconds)
}
