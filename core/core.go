package core

import (
	"fmt"
	"rosswilson/usercapacity/api"
	"rosswilson/usercapacity/utility"
)

func Run() {
	loadUtilities()

	clock := utility.CreateClock()
	apiContext := api.CreateApiContext()
	everhourStrategy, _ := createStrategies(apiContext)

	userData := make(chan []byte)
	go func() {
		data := userCall(apiContext, everhourStrategy)
		userData <- data
	}()

	timeData := timeCall(apiContext, everhourStrategy, clock)
	models := createModels(<-userData, timeData, clock)
	model := bubbleModel(models)

	fmt.Printf("%+v", model)
}
