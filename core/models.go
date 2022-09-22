package core

import (
	"os"
	"rosswilson/usercapacity/model"
	"rosswilson/usercapacity/utility"
)

func createModels(userResp []byte, timeResp []byte) []model.Modeler {
	userModel := model.CreateEverhourUserModel(userResp)
	timeModel := model.CreateEverhourTimeModel(timeResp, userModel)
	mathModel := model.CreateMathModel(timeModel)

	return []model.Modeler{userModel, timeModel, mathModel}
}

func bubbleModel(models []model.Modeler) model.Modeler {
	handler := model.CreateHandler(models)
	model, err := handler.Handle().GetLastModel()

	if err != nil {
		utility.GetLogger().Write(err)
		os.Exit(1)
	}

	return model
}
