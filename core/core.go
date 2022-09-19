package core

import (
	"fmt"
	"rosswilson/usercapacity/api"
)

func Run() {
	apiContext := api.CreateApiContext()
	everhourStrategy, _ := createStrategies(apiContext)

	userData := make(chan []byte)
	go func() {
		data := userCall(apiContext, everhourStrategy)
		userData <- data
	}()

	timeData := timeCall(apiContext, everhourStrategy)
	models := createModels(<-userData, timeData)
	model := bubbleModel(models)

	fmt.Printf("%+v", model)
}
