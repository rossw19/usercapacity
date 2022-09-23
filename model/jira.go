package model

import "rosswilson/usercapacity/utility"

type jiraModel struct {
	users    map[int]user
	previous Modeler
}

func (j *jiraModel) buildModel() {
	utility.GetLogger().Write("model: built jiraModel")
}

func (j jiraModel) GetPrevious() Modeler {
	return j.previous
}

func (j jiraModel) GetUsers() map[int]user {
	return j.users
}

func CreateJiraModel(previous Modeler) *jiraModel {
	return &jiraModel{
		previous: previous,
	}
}
