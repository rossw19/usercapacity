package model

import (
	"encoding/json"
	"rosswilson/usercapacity/utility"
)

type EverhourUserModel struct {
	stream   []byte
	users    map[int]Userable
	previous Modeler
}

func (e *EverhourUserModel) buildModel() {
	type jsonUser struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	var jsonUsers []jsonUser
	json.Unmarshal(e.stream, &jsonUsers)

	e.users = map[int]Userable{}
	for _, j := range jsonUsers {
		e.users[j.Id] = User{
			name: j.Name,
		}
	}

	utility.GetLogger().Write("model: built everhourModel")
}

func (e *EverhourUserModel) GetPrevious() Modeler {
	return e.previous
}

func (e *EverhourUserModel) GetUsers() map[int]Userable {
	return e.users
}

func CreateEverhourUserModel(previous Modeler, data []byte) *EverhourUserModel {
	return &EverhourUserModel{
		previous: previous,
		stream:   data,
	}
}

type EverhourTimeModel struct {
	stream   []byte
	users    map[int]Userable
	previous Modeler
}

func (e *EverhourTimeModel) buildModel() {
	type jsonTime struct {
		Id   int `json:"memberId"`
		Time int `json:"time"`
	}

	var jsonTimes []jsonTime
	json.Unmarshal(e.stream, &jsonTimes)

	previousUsers := e.GetPrevious().GetUsers()

	e.users = map[int]Userable{}
	for _, j := range jsonTimes {
		e.users[j.Id] = User{
			name:        previousUsers[j.Id].GetName(),
			trackedTime: j.Time,
		}
	}

	utility.GetLogger().Write("model: built EverhourTimeModel")
}

func (e *EverhourTimeModel) GetPrevious() Modeler {
	return e.previous
}

func (e *EverhourTimeModel) GetUsers() map[int]Userable {
	return e.users
}

func CreateEverhourTimeModel(previous Modeler, data []byte) *EverhourTimeModel {
	return &EverhourTimeModel{
		previous: previous,
		stream:   data,
	}
}
