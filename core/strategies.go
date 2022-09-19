package core

import (
	"fmt"
	"rosswilson/usercapacity/api"
	"rosswilson/usercapacity/utility"
	"time"
)

func createStrategies(context api.Context) (*api.EverhourStrategy, *api.JiraStrategy) {
	return api.CreateEverhourStrategy(), api.CreateJiraStrategy()
}

func userCall(context api.Context, everhourStrategy *api.EverhourStrategy) []byte {
	everhourStrategy.SetRequestUri("/team/users")
	context.SetApiStrategy(everhourStrategy)
	context.ExecuteApi()
	return everhourStrategy.GetResponseBody()
}

func timeCall(context api.Context, everhourStrategy *api.EverhourStrategy) []byte {
	everhourStrategy.SetRequestUri(timeUri())
	context.SetApiStrategy(everhourStrategy)
	context.ExecuteApi()
	return everhourStrategy.GetResponseBody()
}

func timeUri() string {
	dates := utility.CreateDates(-21, time.Now())
	return fmt.Sprintf("/dashboards/users?date_gte=%s&date_lte=%s", dates.GetFrom(), dates.GetTo())
}

func jiraCall(context api.Context, jiraStrategy *api.JiraStrategy) {
	jiraStrategy.SetRequestUri("")
	context.SetApiStrategy(jiraStrategy)
	context.ExecuteApi()
}
