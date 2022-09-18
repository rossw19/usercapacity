package model

type mathModel struct {
	prototype *Model
}

func (m *mathModel) buildModel() {

}

func (m *mathModel) GetPrototype() *Model {
	return m.prototype
}
