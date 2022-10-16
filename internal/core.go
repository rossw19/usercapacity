package internal

import (
	"fmt"
	"rosswilson/usercapacity/api"
	"rosswilson/usercapacity/model"
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

	logResults(model)
}

func logResults(model model.Modeler) {
	logger := utility.GetLogger()

	for _, u := range model.GetUsers() {
		formatted := utility.GetFormattedTime(u.GetAvgTime())
		message := fmt.Sprintf("%s has %s capacity remaining", u.GetName(), formatted)
		logger.Write(message)
	}
}
