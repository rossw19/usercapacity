package internal

import (
	"rosswilson/usercapacity/utility"
)

func loadUtilities() {
	utility.GetConfig().ReadConfig()
	initLogger(loggingStatus())
}

func loggingStatus() bool {
	return utility.GetConfig().Env.Logging
}

func initLogger(loggingStatus bool) {
	logger := utility.GetLogger()
	logger.SetFile("var/log/debug.log")
	logger.SetActive(loggingStatus)
}
