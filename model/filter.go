package model

import "rosswilson/usercapacity/utility"

type filterModel struct {
	users    map[int]user
	previous Modeler
}

func (f *filterModel) buildModel() {
	targetUsers := utility.GetConfig().Mapping.Users
	previousUsers := f.GetPrevious().GetUsers()

	f.users = map[int]user{}
	for _, t := range targetUsers {
		f.users[t.Id] = user{
			name:        previousUsers[t.Id].GetName(),
			trackedTime: previousUsers[t.Id].GetTimeTracked(),
			averageTime: previousUsers[t.Id].GetAvgTime(),
		}
	}

	utility.GetLogger().Write("model: built filterModel")
}

func (f filterModel) GetPrevious() Modeler {
	return f.previous
}

func (f filterModel) GetUsers() map[int]user {
	return f.users
}

func CreateFilterModel(previous Modeler) *filterModel {
	return &filterModel{
		previous: previous,
	}
}
