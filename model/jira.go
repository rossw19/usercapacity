package model

import "rosswilson/usercapacity/utility"

type JiraModel struct {
	users    map[int]Userable
	previous Modeler
}

func (m *JiraModel) buildModel() {
	targetUsers := utility.GetConfigProxy().GetUsers()
	previousUsers := m.GetPrevious().GetUsers()

	m.users = map[int]Userable{}
	for _, t := range targetUsers {
		everhourId := t.GetEverhourId()

		m.users[everhourId] = User{
			name:        previousUsers[everhourId].GetName(),
			trackedTime: previousUsers[everhourId].GetTimeTracked(),
			averageTime: previousUsers[everhourId].GetAvgTime(),
			daysHadOff:  previousUsers[everhourId].GetDaysHadOff(),
			daysHaveOff: previousUsers[everhourId].GetDaysHaveOff(),
			jiraId:      t.GetJiraId(),
		}
	}

	utility.GetLogger().Write("model: built everhourModel")
}

func (m *JiraModel) GetPrevious() Modeler {
	return m.previous
}

func (m *JiraModel) GetUsers() map[int]Userable {
	return m.users
}

func CreateJiraModel(previous Modeler) *JiraModel {
	return &JiraModel{
		previous: previous,
	}
}
