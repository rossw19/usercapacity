package internal

import (
	"fmt"
	"os"
	"rosswilson/usercapacity/utility"
)

func loadUtilities() {
	config := utility.CreateConfig()
	proxy := utility.GetConfigProxy()
	proxy.SetConfig(config)

	data, err := utility.ReadConfig("config.yml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = proxy.UnmarshalConfig(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	initLogger(loggingStatus())
}

func loggingStatus() bool {
	scope, ok := utility.GetConfigProxy().GetScope("application_utility_logging").ResolveBoolean()
	if ok {
		return scope
	}

	return false
}

func initLogger(loggingStatus bool) {
	logger := utility.GetLogger()
	logger.SetFile("var/log/debug.log")
	logger.SetActive(loggingStatus)
}
