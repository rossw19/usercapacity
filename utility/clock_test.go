package utility

import "testing"

func TestGetFormattedTime(t *testing.T) {
	formattedTime := GetFormattedTime(8193)

	if formattedTime != "02h 16m 33s" {
		t.Errorf("Expected 02h 16m 33s, got %s", formattedTime)
	}
}

func createMockConfig() *Config {
	calendarDaysScope := CreateScope("application_context_calendar_days", "7")
	workingDaysScope := CreateScope("application_context_working_days", "5")
	averageOverScope := CreateScope("application_context_average_over", "3")

	config := CreateConfig()
	config.AddScope(*calendarDaysScope)
	config.AddScope(*workingDaysScope)
	config.AddScope(*averageOverScope)

	return config
}

func TestGetCalendarDays(t *testing.T) {
	proxy := GetConfigProxy()
	proxy.SetConfig(createMockConfig())
	clock := CreateClock()

	if clock.GetCalendarDays() != 7 {
		t.Errorf("Expected 7, got %d", clock.GetCalendarDays())
	}
}

func TestGetWorkingDays(t *testing.T) {
	proxy := GetConfigProxy()
	proxy.SetConfig(createMockConfig())
	clock := CreateClock()

	if clock.GetWorkingDays() != 5 {
		t.Errorf("Expected 5, got %d", clock.GetWorkingDays())
	}
}

func TestGetAverageOver(t *testing.T) {
	proxy := GetConfigProxy()
	proxy.SetConfig(createMockConfig())
	clock := CreateClock()

	if clock.GetAverageOver() != 3 {
		t.Errorf("Expected 3, got %d", clock.GetAverageOver())
	}
}

func TestCreateClock(t *testing.T) {
	proxy := GetConfigProxy()
	proxy.SetConfig(createMockConfig())
	clock := CreateClock()

	if clock == nil {
		t.Errorf("Expected clock to not be nil")
	}
}
