package model

import (
	"encoding/json"
	"rosswilson/usercapacity/utility"
)

type everhourUserModel struct {
	stream    []byte
	users     []user
	prototype Model
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
	times     []time
	prototype *everhourUserModel
}

type time struct {
	id          int
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

	// Encapsulate into our model
	for _, j := range jsonTimes {
		userTime := time{
			id:          j.Id,
			trackedTime: j.Time,
		}

		for _, u := range e.prototype.users {
			if u.id == userTime.id {
				userTime.name = u.name
			}
		}

		e.times = append(e.times, userTime)
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
