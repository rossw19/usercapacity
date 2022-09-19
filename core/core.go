package core

import (
	"fmt"
	"rosswilson/usercapacity/api"
)

func Run() {
	apiContext := api.CreateApiContext()
	everhourStrategy, _ := createStrategies(apiContext)
	userData := userCall(apiContext, everhourStrategy)
	timeData := timeCall(apiContext, everhourStrategy)
	models := createModels(userData, timeData)
	model := bubbleModel(models)

	fmt.Printf("%+v", model)
}
