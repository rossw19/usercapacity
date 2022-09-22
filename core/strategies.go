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

func timeCall(context api.Contexter, everhourStrategy *api.EverhourStrategy) []byte {
	everhourStrategy.SetRequestUri(timeUri())
	context.SetApiStrategy(everhourStrategy)
	context.ExecuteApi()
	return everhourStrategy.GetResponseBody()
}

func timeUri() string {
	dates := utility.CreateDates(-21, time.Now())
	return fmt.Sprintf("/dashboards/users?date_gte=%s&date_lte=%s", dates.GetTo(), dates.GetFrom())
}

func jiraCall(context api.Contexter, jiraStrategy *api.JiraStrategy) {
	jiraStrategy.SetRequestUri("")
	context.SetApiStrategy(jiraStrategy)
	context.ExecuteApi()
}
