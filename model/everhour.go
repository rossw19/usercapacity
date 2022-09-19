package model

import (
	"encoding/json"
	"rosswilson/usercapacity/utility"
)

type everhourUserModel struct {
	stream    []byte
	users     map[int]user
	prototype Model
}

func (e *everhourUserModel) buildModel() {
	type jsonUser struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	var jsonUsers []jsonUser
	json.Unmarshal(e.stream, &jsonUsers)

	e.users = map[int]user{}
	for _, j := range jsonUsers {
		e.users[j.Id] = user{
			name: j.Name,
		}
	}

	utility.GetLogger().Write("model: built everhourModel")
}

func (e everhourUserModel) GetPrototype() Model {
	return e.prototype
}

func (e everhourUserModel) GetUsers() map[int]user {
	return e.users
}

func CreateEverhourUserModel(data []byte) *everhourUserModel {
	return &everhourUserModel{
		stream: data,
	}
}

type everhourTimeModel struct {
	stream    []byte
	users     map[int]user
	prototype Model
}

func (e *everhourTimeModel) buildModel() {
	type jsonTime struct {
		Id   int `json:"memberId"`
		Time int `json:"time"`
	}

	var jsonTimes []jsonTime
	json.Unmarshal(e.stream, &jsonTimes)

	e.users = map[int]user{}
	for _, j := range jsonTimes {
		e.users[j.Id] = user{
			name:        e.users[j.Id].name, // Need to use interface method
			trackedTime: j.Time,
		}
	}

	utility.GetLogger().Write("model: built everhourTimeModel")
}

func (e everhourTimeModel) GetPrototype() Model {
	return e.prototype
}

func (e everhourTimeModel) GetUsers() map[int]user {
	return e.users
}

func CreateEverhourTimeModel(data []byte, prototype Model) *everhourTimeModel {
	return &everhourTimeModel{
		stream:    data,
		prototype: prototype,
	}
}
