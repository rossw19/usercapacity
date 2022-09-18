package model

type Model interface {
	buildModel()
	GetPrototype() *Model
	getType() modelType
}

type modelType byte

const (
	everhourUserType modelType = iota
	everhourTimeType
	mathType
	jiraType
)
