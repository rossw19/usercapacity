package internal

import (
	"rosswilson/usercapacity/utility"
	"testing"
)

func TestLoggingStatus(t *testing.T) {
	scope := utility.CreateScope("application_utility_logging", "true")
	config := utility.CreateConfig()
	config.AddScope(*scope)

	proxy := utility.GetConfigProxy()
	proxy.SetConfig(config)

	if !loggingStatus() {
		t.Error("Expected loggingStatus() to return true")
	}
}
