package internal

import (
	"rosswilson/usercapacity/api"
	"rosswilson/usercapacity/utility"
)

func Run() {
	loadUtilities()

	clock := utility.CreateClock()
	apiContext := api.CreateApiContext()
	everhourStrategy := api.CreateEverhourStrategy()

	userData := make(chan []byte)
	go func() {
		data := userCall(apiContext, everhourStrategy)
		userData <- data
	}()

	timeData := timeCall(apiContext, everhourStrategy, clock)
	models := createModels(<-userData, timeData, clock)
	model := bubbleModel(models)

	jiraStrategies := createJiraStrategies(model.GetUsers())
	jiraCalls(apiContext, jiraStrategies)
}
