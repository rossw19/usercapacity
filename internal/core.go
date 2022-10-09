package internal

import (
	"fmt"
	"rosswilson/usercapacity/api"
	"rosswilson/usercapacity/utility"
)

func Run() {
	loadUtilities()

	clock := utility.CreateClock()
	apiContext := api.CreateApiContext()
	everhourStrategy := api.CreateEverhourStrategy()
	scheduleStrategy := api.CreateFileStrategy()

	userData := make(chan []byte)
	go func() {
		data := userCall(apiContext, everhourStrategy)
		userData <- data
	}()

	scheduleData := make(chan []byte)
	go func() {
		data := scheduleLeaveCall(apiContext, scheduleStrategy)
		scheduleData <- data
	}()

	timeData := timeCall(apiContext, everhourStrategy, clock)
	models := createModels(<-userData, timeData, <-scheduleData, clock)
	model := bubbleModel(models)

	jiraStrategies := createJiraStrategies(model.GetUsers())
	jiraCalls(apiContext, jiraStrategies)

	fmt.Printf("%+v", model)
}
