package utility

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"sync"
)

var lock = &sync.Mutex{}

type Logger struct {
	active    bool
	fileparts struct {
		file string
		path string
	}
}

func (l *Logger) SetFile(filename string) *Logger {
	l.fileparts.file = filename
	l.splitFileparts()

	file, err := l.createOrOpenFile(0711)
	if err != nil {
		fmt.Printf("utility: could not create or open file %s, disabling logging", filename)
		l.active = false
		return l
	}

	log.SetOutput(file)

	return l
}

func (l Logger) createOrOpenFile(chmod fs.FileMode) (*os.File, error) {
	if err := os.MkdirAll(l.fileparts.path, chmod); err != nil {
		return nil, err
	}

	if fileInfo, err := os.Stat(l.fileparts.file); err == nil {
		return os.Open(fmt.Sprintf("%s/%s", l.fileparts.path, fileInfo.Name()))
	}

	return os.Create(l.fileparts.file)
}

func (l *Logger) splitFileparts() {
	if !strings.Contains(l.fileparts.file, "/") {
		l.fileparts.path = ""
		return
	}

	parts := strings.Split(l.fileparts.file, "/")
	path := strings.Join(parts[:len(parts)-1], "/")

	l.fileparts.path = path
}

func (l *Logger) SetActive(active bool) *Logger {
	l.active = active
	return l
}

func (l Logger) Write(line any) {
	if l.active {
		log.Println(line)
		return
	}

	fmt.Println(line)
}

var loggerInstance *Logger

func GetLogger() *Logger {
	if loggerInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if loggerInstance == nil {
			loggerInstance = &Logger{}
		}
	}

	return loggerInstance
}
