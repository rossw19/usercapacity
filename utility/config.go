package utility

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strconv"
	"sync"
)

type Scoper interface {
	ResolveInt() (int, bool)
	ResolveBoolean() (bool, bool)
	ResolveString() (string, bool)
	GetPath() string
}

type Scope struct {
	path  string
	value string
}

func (s Scope) ResolveInt() (int, bool) {
	num, err := strconv.Atoi(s.value)
	if err != nil {
		return 0, false
	}

	return num, true
}

func (s Scope) ResolveBoolean() (bool, bool) {
	if s.value == "true" {
		return true, true
	}

	if s.value == "false" {
		return false, true
	}

	return false, false
}

func (s Scope) ResolveString() (string, bool) {
	if s.value == "" {
		return "", false
	}

	return s.value, true
}

func (s Scope) GetPath() string {
	return s.path
}

func CreateScope(path string, value string) *Scope {
	return &Scope{path: path, value: value}
}

type Configurable interface {
	GetScope(string) Scoper
	AddScope(Scoper)
	GetUsers() []Userable
	AddUser(Userable)
	UnmarshalConfig([]byte) error
}

type Config struct {
	scopes []Scoper
	users  []Userable
}

func (c *Config) GetScope(path string) Scoper {
	for _, s := range c.scopes {
		if s.GetPath() == path {
			return s.(Scope)
		}
	}

	return Scope{}
}

func (c *Config) AddScope(scope Scoper) {
	c.scopes = append(c.scopes, scope)
}

func (c *Config) GetUsers() []Userable {
	return c.users
}

func (c *Config) AddUser(user Userable) {
	c.users = append(c.users, user)
}

func (c *Config) UnmarshalConfig(data []byte) error {
	type InternalScope struct {
		Path  string `yaml:"path"`
		Value string `yaml:"value"`
	}

	type InternalUser struct {
		EverhourId int    `yaml:"everhourId"`
		JiraId     string `yaml:"jiraId"`
		Name       string `yaml:"name"`
	}

	type InternalConfig struct {
		InternalScopes []InternalScope `yaml:"scopes"`
		InternalUsers  []InternalUser  `yaml:"users"`
	}

	config := InternalConfig{}

	err := yaml.Unmarshal(data, &config)
	if err != nil {
		return fmt.Errorf("utility: could not parse yml, check config validity: %s", err)
	}

	for _, s := range config.InternalScopes {
		c.scopes = append(c.scopes, Scope{
			path:  s.Path,
			value: s.Value,
		})
	}

	for _, u := range config.InternalUsers {
		c.users = append(c.users, User{
			everhourId: u.EverhourId,
			jiraId:     u.JiraId,
			name:       u.Name,
		})
	}

	return nil
}

func ReadConfig(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []byte{}, errors.New("utility: could not read file config.yml")
	}

	return data, nil
}

func CreateConfig() *Config {
	return &Config{}
}

type ConfigProxy struct {
	config Configurable
}

func (c *ConfigProxy) GetScope(path string) Scoper {
	return c.config.GetScope(path)
}

func (c *ConfigProxy) AddScope(scope Scoper) {
	c.config.AddScope(scope)
}

func (c *ConfigProxy) GetUsers() []Userable {
	return c.config.GetUsers()
}

func (c *ConfigProxy) AddUser(user Userable) {
	c.config.AddUser(user)
}

func (c *ConfigProxy) UnmarshalConfig(data []byte) error {
	return c.config.UnmarshalConfig(data)
}

func (c *ConfigProxy) SetConfig(config Configurable) {
	c.config = config
}

var configProxyLock = &sync.Mutex{}
var configProxyInstance *ConfigProxy

func GetConfigProxy() *ConfigProxy {
	if configProxyInstance == nil {
		configProxyLock.Lock()
		defer configProxyLock.Unlock()

		if configProxyInstance == nil {
			configProxyInstance = &ConfigProxy{}
		}
	}

	return configProxyInstance
}

type Userable interface {
	GetEverhourId() int
	GetJiraId() string
	GetName() string
}

type User struct {
	everhourId int
	jiraId     string
	name       string
}

func (u User) GetEverhourId() int {
	return u.everhourId
}

func (u User) GetJiraId() string {
	return u.jiraId
}

func (u User) GetName() string {
	return u.name
}

func CreateUser(everhourId int, jiraId string, name string) *User {
	return &User{
		everhourId: everhourId,
		jiraId:     jiraId,
		name:       name,
	}
}
