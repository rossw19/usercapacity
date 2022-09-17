package utility

import (
	"testing"
)

func TestSplitFileParts(t *testing.T) {
	logger := GetLogger()
	logger.fileparts.file = "var/log/debug.log"
	logger.splitFileparts()

	if logger.fileparts.path != "var/log" {
		t.Fail()
	}
}
