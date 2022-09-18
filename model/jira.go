package model

type jiraModel struct {
	prototype *Model
}

func (j *jiraModel) buildModel() {

}

func (j *jiraModel) GetPrototype() *Model {
	return j.prototype
}
