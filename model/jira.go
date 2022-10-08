package model

import "rosswilson/usercapacity/utility"

type JiraModel struct {
	users    map[int]User
	previous Modeler
}

func (m *JiraModel) buildModel() {
	targetUsers := utility.GetConfig().Mapping.Users
	previousUsers := m.GetPrevious().GetUsers()

	m.users = map[int]User{}
	for _, t := range targetUsers {
		m.users[t.EverhourId] = user{
			name:        previousUsers[t.EverhourId].GetName(),
			trackedTime: previousUsers[t.EverhourId].GetTimeTracked(),
			averageTime: previousUsers[t.EverhourId].GetAvgTime(),
			jiraId:      t.JiraId,
		}
	}

	utility.GetLogger().Write("model: built everhourModel")
}

func (m *JiraModel) GetPrevious() Modeler {
	return m.previous
}

func (m *JiraModel) GetUsers() map[int]User {
	return m.users
}

func CreateJiraModel(previous Modeler) *JiraModel {
	return &JiraModel{
		previous: previous,
	}
}
