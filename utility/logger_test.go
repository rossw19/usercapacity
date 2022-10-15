package utility

import (
	"testing"
)

func TestSplitFileParts(t *testing.T) {
	logger := CreateLogger()
	logger.fileparts.file = "var/log/debug.log"
	logger.splitFileparts()

	if logger.fileparts.path != "var/log" {
		t.Fail()
	}
}

func TestSetLogger(t *testing.T) {
	logger := CreateLogger()
	loggerProxy := GetLogger()
	loggerProxy.SetLogger(logger)

	if loggerProxy.logger != logger {
		t.Error("could not set logger")
	}
}

func TestGetLogger(t *testing.T) {
	logger0 := GetLogger()
	logger1 := GetLogger()

	if logger0 == nil {
		t.Error("could not instantiate logger")
	}

	if logger1 == nil {
		t.Error("could not instantiate logger")
	}

	if logger0 != logger1 {
		t.Error("logger is not singleton")
	}
}

func TestSetActive(t *testing.T) {
	logger := CreateLogger()
	logger.SetActive(true)

	if !logger.active {
		t.Error("could not activate logger")
	}

	logger.SetActive(false)

	if logger.active {
		t.Error("could not deactivate logger")
	}
}

func TestCreateLogger(t *testing.T) {
	logger := CreateLogger()

	if logger == nil {
		t.Error("could not instantiate logger")
	}
}
