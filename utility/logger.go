package utility

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"sync"
)

var loggerLock = &sync.Mutex{}

type Loggable interface {
	SetFile(filename string) Loggable
	SetActive(active bool) Loggable
	Write(line any)
}

type LoggerProxy struct {
	logger Loggable
}

func (l *LoggerProxy) SetLogger(logger Loggable) *LoggerProxy {
	l.logger = logger
	return l
}

func (l *LoggerProxy) SetFile(filename string) Loggable {
	l.logger.SetFile(filename)
	return l
}

func (l *LoggerProxy) SetActive(active bool) Loggable {
	l.logger.SetActive(active)
	return l
}

func (l *LoggerProxy) Write(line any) {
	l.logger.Write(line)
}

var loggerInstance *LoggerProxy

func GetLogger() *LoggerProxy {
	if loggerInstance == nil {
		loggerLock.Lock()
		defer loggerLock.Unlock()

		if loggerInstance == nil {
			loggerInstance = &LoggerProxy{}
		}
	}

	return loggerInstance
}

type Logger struct {
	active    bool
	fileparts struct {
		file string
		path string
	}
}

func (l *Logger) SetFile(filename string) Loggable {
	l.fileparts.file = filename
	l.splitFileparts()

	file, err := l.createOrOpenFile(0666)

	if err != nil {
		fmt.Printf("utility: could not create or open file %s, disabling logging", filename)
		l.active = false
		return l
	}

	log.SetOutput(file)

	return l
}

func (l *Logger) createOrOpenFile(chmod fs.FileMode) (*os.File, error) {
	if err := os.MkdirAll(l.fileparts.path, chmod); err != nil {
		return nil, err
	}

	if fileInfo, err := os.Stat(l.fileparts.file); err == nil {
		return os.OpenFile(fmt.Sprintf("%s/%s", l.fileparts.path, fileInfo.Name()), os.O_RDWR|os.O_APPEND, chmod)
	}

	file, err := os.Create(l.fileparts.file)
	if err != nil {
		return nil, err
	}

	err = file.Chmod(chmod)
	if err != nil {
		return nil, err
	}

	return file, nil
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

func (l *Logger) SetActive(active bool) Loggable {
	l.active = active
	return l
}

func (l *Logger) Write(line any) {
	if l.active {
		log.Println(line)
		return
	}

	fmt.Println(line)
}

func CreateLogger() *Logger {
	return &Logger{}
}
