package core

import (
	"os"
	"rosswilson/usercapacity/model"
	"rosswilson/usercapacity/utility"
)

func createModels(userResp []byte, timeResp []byte) []model.Model {
	userModel := model.CreateEverhourUserModel(userResp)
	timeModel := model.CreateEverhourTimeModel(timeResp, userModel)
	mathModel := model.CreateMathModel(timeModel)

	return []model.Model{userModel, timeModel, mathModel}
}

func bubbleModel(models []model.Model) model.Model {
	handler := model.CreateHandler(models)
	model, err := handler.Handle().GetLastModel()

	if err != nil {
		utility.GetLogger().Write(err)
		os.Exit(1)
	}

	return model
}
