package utility

import (
	"fmt"
	"os"
	"strconv"
)

func GetEnvOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("utility: env variable %s could not be found", key))
	}

	return value
}

func StringToBool(value string) bool {
	b, err := strconv.ParseBool(value)
	if err != nil {
		fmt.Printf("utility: could not convert string %s to bool, using falsey", value)
		return false
	}

	return b
}
