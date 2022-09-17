package main

import (
	"rosswilson/usercapacity/api"
	"rosswilson/usercapacity/utility"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("cmd: error loading .env file")
	}

	loggingEnv := utility.GetEnvOrPanic("LOGGING")
	loggingActive := utility.StringToBool(loggingEnv)
	utility.GetLogger().SetFile("var/log/debug.log").SetActive(loggingActive)

	everhourStrategy := api.CreateEverhourStrategy()
	apiContext := api.CreateApiContext()
	apiContext.SetApiStrategy(everhourStrategy)
	apiContext.ExecuteApi()
}
