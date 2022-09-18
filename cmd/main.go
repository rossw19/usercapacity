package main

import (
	"errors"
	"fmt"
	"os"
	"rosswilson/usercapacity/core"
	"rosswilson/usercapacity/utility"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	initLogger(loggingStatus())
	core.CreateCore().Run()
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
