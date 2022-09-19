package model

type Model interface {
	buildModel()
	GetPrototype() Model
}
