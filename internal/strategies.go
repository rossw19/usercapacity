package internal

import (
	"fmt"
	"rosswilson/usercapacity/api"
	"rosswilson/usercapacity/model"
	"rosswilson/usercapacity/utility"
	"time"
)

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

func jiraCall(context api.Contexter, jiraStrategy *api.JiraStrategy, requestUri string) {
	jiraStrategy.SetRequestUri(requestUri)
	context.SetApiStrategy(jiraStrategy)
	context.ExecuteApi()
}

func createJiraStrategies(users map[int]model.Userable) []api.JiraStrategy {
	var strategies []api.JiraStrategy
	for _, u := range users {
		strategy := api.CreateJiraStrategy(u)
		strategies = append(strategies, *strategy)
	}

	return strategies
}

func jiraCalls(context api.Contexter, jiraStrategies []api.JiraStrategy) {
	for _, u := range jiraStrategies {
		jiraCall(context, &u, "/rest/api/3/user/properties/capacity?accountId=")
	}
}

func scheduleLeaveCall(context api.Contexter, scheduleStrategy *api.FileStrategy) []byte {
	scheduleStrategy.SetRequestUri("***REMOVED***")
	context.SetApiStrategy(scheduleStrategy)
	context.ExecuteApi()
	return scheduleStrategy.GetFileContents()
}
