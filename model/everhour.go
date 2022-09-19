package model

import (
	"encoding/json"
	"rosswilson/usercapacity/utility"
)

type everhourUserModel struct {
	stream    []byte
	users     map[int]string
	prototype Model
}

func (e *everhourUserModel) buildModel() {
	type jsonUser struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	var jsonUsers []jsonUser
	json.Unmarshal(e.stream, &jsonUsers)

	e.users = map[int]string{}
	for _, j := range jsonUsers {
		e.users[j.Id] = j.Name
	}

	utility.GetLogger().Write("model: built everhourModel")
}

func (e everhourUserModel) GetPrototype() Model {
	return e.prototype
}

func CreateEverhourUserModel(data []byte) *everhourUserModel {
	return &everhourUserModel{
		stream: data,
	}
}

type everhourTimeModel struct {
	stream    []byte
	times     map[int]time
	prototype *everhourUserModel
}

type time struct {
	name        string
	trackedTime int
}

func (e *everhourTimeModel) buildModel() {
	type jsonTime struct {
		Id   int `json:"memberId"`
		Time int `json:"time"`
	}

	var jsonTimes []jsonTime
	json.Unmarshal(e.stream, &jsonTimes)

	e.times = map[int]time{}
	for _, j := range jsonTimes {
		e.times[j.Id] = time{
			name:        e.prototype.users[j.Id],
			trackedTime: j.Time,
		}
	}

	utility.GetLogger().Write("model: built everhourTimeModel")
}

func (e everhourTimeModel) GetPrototype() Model {
	return e.prototype
}

func CreateEverhourTimeModel(data []byte, prototype *everhourUserModel) *everhourTimeModel {
	return &everhourTimeModel{
		stream:    data,
		prototype: prototype,
	}
}
