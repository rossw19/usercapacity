package model

import "rosswilson/usercapacity/utility"

type jiraModel struct {
	users     map[int]user
	prototype Modeler
}

func (j *jiraModel) buildModel() {
	utility.GetLogger().Write("model: built jiraModel")
}

func (j jiraModel) GetPrototype() Modeler {
	return j.prototype
}

func (e jiraModel) GetUsers() map[int]user {
	return e.users
}

func CreateJiraModel(prototype Modeler) *jiraModel {
	return &jiraModel{
		prototype: prototype,
	}
}
