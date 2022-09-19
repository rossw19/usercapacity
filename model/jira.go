package model

import "rosswilson/usercapacity/utility"

type jiraModel struct {
	users     map[int]user
	prototype Model
}

func (j *jiraModel) buildModel() {
	utility.GetLogger().Write("model: built jiraModel")
}

func (j jiraModel) GetPrototype() Model {
	return j.prototype
}

func (e jiraModel) GetUsers() map[int]user {
	return e.users
}

func CreateJiraModel(prototype Model) *jiraModel {
	return &jiraModel{
		prototype: prototype,
	}
}
