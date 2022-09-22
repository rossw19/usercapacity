package model

import (
	"encoding/json"
	"rosswilson/usercapacity/utility"
)

type everhourUserModel struct {
	stream   []byte
	users    map[int]user
	previous Modeler
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

func (e everhourUserModel) GetPrevious() Modeler {
	return e.previous
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
	stream   []byte
	users    map[int]user
	previous Modeler
}

func (e *everhourTimeModel) buildModel() {
	type jsonTime struct {
		Id   int `json:"memberId"`
		Time int `json:"time"`
	}

	var jsonTimes []jsonTime
	json.Unmarshal(e.stream, &jsonTimes)

	previousUsers := e.GetPrevious().GetUsers()

	e.users = map[int]user{}
	for _, j := range jsonTimes {
		e.users[j.Id] = user{
			name:        previousUsers[j.Id].GetName(),
			trackedTime: j.Time,
		}
	}

	utility.GetLogger().Write("model: built everhourTimeModel")
}

func (e everhourTimeModel) GetPrevious() Modeler {
	return e.previous
}

func (e everhourTimeModel) GetUsers() map[int]user {
	return e.users
}

func CreateEverhourTimeModel(data []byte, previous Modeler) *everhourTimeModel {
	return &everhourTimeModel{
		stream:   data,
		previous: previous,
	}
}
