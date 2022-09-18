package model

import (
	"encoding/json"
	"rosswilson/usercapacity/utility"
)

type everhourUserModel struct {
	stream    []byte
	users     []user
	prototype *Model
}

type user struct {
	id   int
	name string
}

func (e *everhourUserModel) buildModel() {
	type jsonUser struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	var jsonUsers []jsonUser
	json.Unmarshal(e.stream, &jsonUsers)

	// Encapsulate into our model
	for _, j := range jsonUsers {
		e.users = append(e.users, user{
			id:   j.Id,
			name: j.Name,
		})
	}

	utility.GetLogger().Write("model: built everhourUserModel")
}

func (e everhourUserModel) GetPrototype() *Model {
	return e.prototype
}

func CreateEverhourUserModel(data []byte) *everhourUserModel {
	return &everhourUserModel{
		stream: data,
	}
}

type everhourTimeModel struct {
	stream    []byte
	times     []time
	prototype *Model
}

type time struct {
	id          int
	trackedTime int
}

func (e *everhourTimeModel) buildModel() {
	type jsonTime struct {
		Id   int `json:"memberId"`
		Time int `json:"time"`
	}

	var jsonTimes []jsonTime
	json.Unmarshal(e.stream, &jsonTimes)

	// Encapsulate into our model
	for _, j := range jsonTimes {
		e.times = append(e.times, time{
			id:          j.Id,
			trackedTime: j.Time,
		})
	}

	utility.GetLogger().Write("model: built everhourTimeModel")
}

func (e everhourTimeModel) GetPrototype() *Model {
	return e.prototype
}

func CreateEverhourTimeModel(data []byte, model *Model) *everhourTimeModel {
	return &everhourTimeModel{
		stream:    data,
		prototype: model,
	}
}
