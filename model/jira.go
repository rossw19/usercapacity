package model

import "rosswilson/usercapacity/utility"

type jiraModel struct {
	prototype Model
}

func (j *jiraModel) buildModel() {
	utility.GetLogger().Write("model: built jiraModel")
}

func (j *jiraModel) GetPrototype() Model {
	return j.prototype
}
