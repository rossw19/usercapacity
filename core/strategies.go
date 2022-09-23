package core

import (
	"fmt"
	"rosswilson/usercapacity/api"
	"rosswilson/usercapacity/utility"
	"time"
)

func createStrategies(context api.Contexter) (*api.EverhourStrategy, *api.JiraStrategy) {
	return api.CreateEverhourStrategy(), api.CreateJiraStrategy()
}

func userCall(context api.Contexter, everhourStrategy *api.EverhourStrategy) []byte {
	everhourStrategy.SetRequestUri("/team/users")
	context.SetApiStrategy(everhourStrategy)
	context.ExecuteApi()
	return everhourStrategy.GetResponseBody()
}

func timeCall(context api.Contexter, everhourStrategy *api.EverhourStrategy, clock utility.Clocker) []byte {
	everhourStrategy.SetRequestUri(timeUri(clock))
	context.SetApiStrategy(everhourStrategy)
	context.ExecuteApi()
	return everhourStrategy.GetResponseBody()
}

func timeUri(clock utility.Clocker) string {
	days := -1 * clock.GetCalendarDays() * clock.GetAverageOver()
	dates := utility.CreateDates(int64(days), time.Now())
	return fmt.Sprintf("/dashboards/users?date_gte=%s&date_lte=%s", dates.GetTo(), dates.GetFrom())
}

func jiraCall(context api.Contexter, jiraStrategy *api.JiraStrategy) {
	jiraStrategy.SetRequestUri("")
	context.SetApiStrategy(jiraStrategy)
	context.ExecuteApi()
}
