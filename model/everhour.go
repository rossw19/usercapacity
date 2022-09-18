package model

import "encoding/json"

type everhourUserModel struct {
	stream    []byte
	users     []user
	modelType modelType
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
}

func (e everhourUserModel) GetPrototype() *Model {
	return e.prototype
}

func (e everhourUserModel) getType() modelType {
	return e.modelType
}

func CreateEverhourUserModel(data []byte) *everhourUserModel {
	return &everhourUserModel{
		stream:    data,
		modelType: everhourUserType,
	}
}

type everhourTimeModel struct {
	stream    []byte
	times     []time
	modelType modelType
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
}

func (e everhourTimeModel) GetPrototype() *Model {
	return e.prototype
}

func (e everhourTimeModel) getType() modelType {
	return e.modelType
}

func CreateEverhourTimeModel(data []byte, model *Model) *everhourTimeModel {
	return &everhourTimeModel{
		stream:    data,
		modelType: everhourTimeType,
		prototype: model,
	}
}
