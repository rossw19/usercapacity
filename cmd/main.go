package main

import (
	"fmt"
	"rosswilson/usercapacity/api"
	"rosswilson/usercapacity/utility"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	initLogger(loggingStatus())

	dates := utility.CreateDates(-21, time.Now())

	everhourStrategy := api.CreateEverhourStrategy()
	everhourStrategy.SetRequestUri(fmt.Sprintf("/team/time?from=%s&to=%s", dates.GetFrom(), dates.GetTo()))

	apiContext := api.CreateApiContext()
	apiContext.SetApiStrategy(everhourStrategy)
	apiContext.ExecuteApi()
}

func loggingStatus() bool {
	loggingEnv := utility.GetEnvOrPanic("LOGGING")
	return utility.StringToBool(loggingEnv)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("cmd: error loading .env file")
	}
}

func initLogger(loggingStatus bool) {
	logger := utility.GetLogger()
	logger.SetFile("var/log/debug.log")
	logger.SetActive(loggingStatus)
}
