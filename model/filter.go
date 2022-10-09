package model

import (
	"rosswilson/usercapacity/utility"
)

type FilterModel struct {
	users    map[int]User
	previous Modeler
}

func (f *FilterModel) buildModel() {
	targetUsers := utility.GetConfig().Mapping.Users
	previousUsers := f.GetPrevious().GetUsers()

	f.users = map[int]User{}
	for _, t := range targetUsers {
		f.users[t.EverhourId] = user{
			name:        previousUsers[t.EverhourId].GetName(),
			trackedTime: previousUsers[t.EverhourId].GetTimeTracked(),
			averageTime: previousUsers[t.EverhourId].GetAvgTime(),
			daysHadOff:  previousUsers[t.EverhourId].GetDaysHadOff(),
			daysHaveOff: previousUsers[t.EverhourId].GetDaysHaveOff(),
		}
	}

	utility.GetLogger().Write("model: built FilterModel")
}

func (f *FilterModel) GetPrevious() Modeler {
	return f.previous
}

func (f *FilterModel) GetUsers() map[int]User {
	return f.users
}

func CreateFilterModel(previous Modeler) *FilterModel {
	return &FilterModel{
		previous: previous,
	}
}
