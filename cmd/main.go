package main

import (
	"errors"
	"fmt"
	"os"
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
	loggingEnv := utility.GetEnvOrExit("LOGGING")
	return utility.StringToBool(loggingEnv)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(errors.New("cmd: error loading .env file"))
		os.Exit(1)
	}
}

func initLogger(loggingStatus bool) {
	logger := utility.GetLogger()
	logger.SetFile("var/log/debug.log")
	logger.SetActive(loggingStatus)
}
