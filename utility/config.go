package utility

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

var configLock = &sync.Mutex{}

type Config interface {
	ReadConfig() Config
}

type Configuration struct {
	Env struct {
		Everhour struct {
			Url     string `yaml:"url"`
			Auth    string `yaml:"auth"`
			Version string `yaml:"version"`
		}
		Jira struct {
			Url  string `yaml:"url"`
			User string `yaml:"user"`
			Pass string `yaml:"pass"`
		}
		Logging bool `yaml:"loggin"`
	}
}

var configInstance *Configuration

func (c *Configuration) ReadConfig() *Configuration {
	data, err := os.ReadFile("config.yml")
	if err != nil {
		fmt.Println("utility: could not read file config.yml")
		os.Exit(1)
	}

	err = yaml.Unmarshal(data, c)
	if err != nil {
		fmt.Println("utility: could not parse yml, check config validity")
	}

	return c
}

func GetConfig() *Configuration {
	if configInstance == nil {
		configLock.Lock()
		defer configLock.Unlock()

		if configInstance == nil {
			configInstance = &Configuration{}
		}
	}

	return configInstance
}
