package model

import "rosswilson/usercapacity/utility"

type mathModel struct {
	prototype *Model
}

func (m *mathModel) buildModel() {
	utility.GetLogger().Write("model: built mathModel")
}

func (m *mathModel) GetPrototype() *Model {
	return m.prototype
}
